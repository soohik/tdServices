// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"tdapi/dataservice"
	"tdapi/log"
	"tdapi/model"
	"time"

	"github.com/Arman92/go-tdlib"
)

const limit int32 = 200

const (
	TDURL = "https://t.me/"
)

const (
	tddata = "../../tddata/"
	tdfile = "../../tdfile/"
)

const (
	REG_OK       = 1
	REG_NOTFOUND = 2
)

// ListUserUseCase implements ListUseCaseInterface.
type ClientManagerUseCase struct {
	TdInstances map[string]*tdlib.Client
}

var ClientManager ClientManagerUseCase

func (c *ClientManagerUseCase) LoadTdInstance() error {

	c.TdInstances = make(map[string]*tdlib.Client)

	clients, err := dataservice.GetAllPhone()
	if err != nil {
		return nil
	}
	c.AddTdlibClient(clients)

	return nil
}

//
func (c *ClientManagerUseCase) RemoveClient(phone string) bool {

	return false
}

//
func (c *ClientManagerUseCase) GetClient(phone string) (model.Phone, bool) {

	return dataservice.GetPhone(phone)

}

func RegisterPhone(account, code string) (model.Client, int) {

	_, find := ClientManager.GetClient(account)
	if !find {
		return model.Client{}, model.PhoneNOTFOUND
	}

	client, ret := ClientManager.GetTdClient(account)
	fmt.Println(client, ret)
	if ret == model.AuthWaitCode {
		return RegisterClient(account, code, client)
	} else {
		var repclient model.Client
		user, _ := client.GetMe()
		repclient.Id = user.ID
		repclient.Name = user.FirstName
		repclient.Username = user.Username
		repclient.PhoneNumber = user.PhoneNumber
		ClientManager.AddClient(account, client)
		return repclient, model.SOK

	}

}

func PreRegisterPhone(phonenumber string) (bool, int) {

	var phone model.Phone
	phone.Account = phonenumber
	phone.Phone = phonenumber
	phone.Registered = 0
	phone.Tddata = tddata + phonenumber + "-tdlib-db"

	phone.Tdfile = tdfile + phonenumber + "-tdlib-files"

	ok := dataservice.Preregister(phone)
	if ok {
		client, reg := ClientManager.GetTdClient(phone.Account)
		if reg == model.AuthWaitCode { //需要注册
			ret := SendPhoneNumber(phone.Account, client)

			return ret, model.SOK
		}
		// ClientManager.AddInstance(phonenumber, "")

	}

	return false, model.SOK

}

// //增加单个客户端
// func (c *ClientManagerUseCase) ReAddClient(phonenumber string) {

// 	find := c.RemoveClient(phonenumber)

// 	if find {

// 		var phone model.Phone
// 		phone.Account = phonenumber
// 		phone.Phone = phonenumber
// 		phone.Registered = 0
// 		phone.Tddata = tddata + phonenumber + "-tdlib-db"

// 		phone.Tdfile = tdfile + phonenumber + "-tdlib-files"

// 		if find { //插入成功
// 			ClientManager.AddClient(phone)
// 		}
// 	}

// }

func (c *ClientManagerUseCase) AddTdlibClient(m []model.Phone) {

	for _, value := range m {
		c.AddInstance(value.Account)
	}

}

func BuildClientManager() {

	ClientManager.LoadTdInstance()

}

//创建群，邀请别人加入群
func Joinlink(account, linkurl, groupname string) (int, error) {

	client := ClientManager.TdInstances[account]

	dataservice.InsertGroup(groupname, linkurl) //
	// dataservice.InsertGroupsInfo(account, groupname)

	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	char, err := client.SearchPublicChat(groupname)
	if err != nil {
		return model.BadRequest, errors.New("无效的组")
	}
	ok, _ := client.JoinChat(char.ID)
	fmt.Println(ok, err)

	if err != nil {
		return model.AuthSenCodeErr, errors.New("加入失败")
	}

	return model.SOK, nil
}

func (c *ClientManagerUseCase) AddInstance(account string) {

	tdlib.SetLogVerbosityLevel(1)
	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tddata + account + "-tdlib-db",
		FileDirectory:       tdfile + account + "-tdlib-files",
		IgnoreFileNames:     false,
	})

	for {
		currentState, err := client.Authorize()
		if err != nil {
			log.Infof("Error getting current state: %s %v", account, err)
			break
		}

		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			client.Close()
			break
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			client.Close()
			break

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			log.Info("验证通过-- ", account) //不用关闭客户端

			c.TdInstances[account] = client
			break

		}
	}
	go func() {
		rawUpdates := client.GetRawUpdatesChannel(100)
		me, _ := client.GetMe()

		for update := range rawUpdates {
			// Show all updates

			t, ok := update.Data["@type"]
			if !ok {
				continue
			}

			msgType, ok := t.(string)
			if !ok {
				continue
			}
			if msgType == "updateUser" {

				switch tdlib.UpdateEnum(update.Data["@type"].(string)) {

				case tdlib.UpdateUserType:

					var up tdlib.UpdateUser
					json.Unmarshal(update.Raw, &up)
					if up.User.PhoneNumber != me.PhoneNumber {
						insertUserIfNotExists(me.PhoneNumber, up.User)
					}

				}

				continue
			}
			if msgType == "updateUserStatus" {
				fmt.Println(update.Data["@type"])
				switch tdlib.UpdateEnum(update.Data["@type"].(string)) {
				case tdlib.UpdateUserStatusType:

					var up tdlib.UpdateUserStatus
					json.Unmarshal(update.Raw, &up)
					if up.UserID != me.ID {
						user, _ := client.GetUser(up.UserID)
						insertUserIfNotExists(me.PhoneNumber, user)
					}

				}
				continue
			}

		}

	}()

}

func (c *ClientManagerUseCase) GetInstance(account string) (*tdlib.Client, int) {
	var ret int

	tdlib.SetLogVerbosityLevel(1)
	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tddata + account + "-tdlib-db",
		FileDirectory:       tdfile + account + "-tdlib-files",
		IgnoreFileNames:     false,
	})

	for {
		currentState, err := client.Authorize()
		if err != nil {
			log.Infof("Error getting current state: %s %v", account, err)
			break
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			_, err := client.SendPhoneNumber(account)
			if err != nil {
				log.Infof("phone %s  err: %v", account, err)
				ret = model.AuthSendTimeout
				break
			}
			ret = model.AuthWaitCode

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			log.Info("验证通过-- ", account)

			break

		}
	}

	return client, ret
}

func Getallgroups(account string, agent int) ([]model.Groups, error) {
	// return dataservice.GetAllGroups(agent)
	var groups []model.Groups

	client, _ := ClientManager.GetTdClient(account)

	if client == nil {
		return nil, errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	chats, err := getChatList(client, 100, false)
	if err != nil {
		return nil, err
	}
	for _, value := range chats {
		switch value.Type.(type) {
		case *tdlib.ChatTypeSupergroup:
			chattype := value.Type.(*tdlib.ChatTypeSupergroup)
			if chattype == nil {
				continue
			}
			var m model.Groups
			// m.Uid = chattype.SupergroupID
			m.Name = value.Title
			groups = append(groups, m)
		case *tdlib.ChatTypeBasicGroup:
			continue
		}

	}
	defer client.Close()
	return groups, nil
}

func Getallchats(account string, agent int) ([]model.Groups, error) {
	// return dataservice.GetAllGroups(agent)
	var groups []model.Groups

	client, _ := ClientManager.GetTdClient(account)

	if client == nil {
		return nil, errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	chats, err := getChatList(client, 100, false)
	if err != nil {
		return nil, err
	}

	for _, chat := range chats {
		switch chat.Type.GetChatTypeEnum() {
		case tdlib.ChatTypeSupergroupType:
			spChat, ok := chat.Type.(*tdlib.ChatTypeSupergroup)
			if !ok {
				//log.Println("can't convert to super group")
				break
			}
			group, err := client.GetSupergroup(spChat.SupergroupID)
			if err != nil {
				//log.Println("can't get super group", err)
				break
			}
			fmt.Print("super group:", chat.Title, group.MemberCount, group.Username, chat.ID)
			fullInfo, err := client.GetSupergroupFullInfo(spChat.SupergroupID)
			if err != nil {
				//log.Println("can't get super group full info", err)
				break
			}
			if !fullInfo.CanGetMembers {
				//log.Println("can't get members from this group", chat.Title)
				break
			}
			{

				if fullInfo.MemberCount > 10000 {
					getSupergroupMemebers(client, spChat.SupergroupID)
				}

				if fullInfo.MemberCount > 10000 {
					getChatMembers(client, chat.ID)
				}
			}
		}

	}

	// for _, value := range chats {

	// 	var m model.Groups
	// 	m.Uid = strconv.FormatInt(value.ID, 10)
	// 	m.Name = value.Title
	// 	groups = append(groups, m)

	// }
	defer client.Close()
	return groups, nil
}

func GetMegroups(agent string) ([]model.Groupinfos, error) {
	return dataservice.GetMeGroups(agent)
}

func CreateBasicGroup(account string, f model.Friends) error {
	client := ClientManager.TdInstances[account]
	if client == nil {
		return errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	chat, err := client.CreateNewSupergroupChat(f.Uname, false, f.Uname, nil, false)
	chattype := chat.Type.(*tdlib.ChatTypeSupergroup)
	if chattype == nil {
		return errors.New("转换错误！")
	}

	if err != nil {
		return err
	}
	_, err = client.AddChatMembers(chat.ID, f.Cids)
	if err != nil {
		return err
	}

	var m model.Groupinfos
	m.Chatid = chat.ID
	m.Groupname = fmt.Sprintf("%s%s", TDURL, f.Uname)
	m.Phone = account
	m.Uid = chattype.SupergroupID

	_, err = client.SetSupergroupUsername(m.Uid, m.Groupname)

	if err != nil {
		dataservice.InsertGroupsInfo(m)
	}

	return nil
}

func AddContacts(c *model.AddContacts) error {

	client := ClientManager.TdInstances[c.Phone]
	if client == nil {
		return errors.New("找不到账号！")
	}

	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	var contacts []tdlib.Contact
	var m []model.Contacts

	for _, value := range c.Contents {

		contact := tdlib.Contact{}
		contact.PhoneNumber = value
		contacts = append(contacts, contact)

	}
	ok, _ := client.ImportContacts(contacts)
	for _, value := range ok.UserIDs {
		var k model.Contacts

		user, _ := client.GetUser(value)
		fmt.Println(user)

		k.Account = c.Phone
		k.Contactid = int(user.ID)
		k.Contactname = user.Username
		k.Contactphone = user.PhoneNumber
		k.Status = string(user.Status.GetUserStatusEnum())
		m = append(m, k)
	}

	return dataservice.InsertContact(m)

}

func GetmeContents(c *model.Me) error {
	client := ClientManager.TdInstances[c.Name]
	if client == nil {
		return errors.New("找不到账号！")
	}
	client.GetContacts()
	return nil
}

func (c *ClientManagerUseCase) InsertContact(m []model.Contacts) error {

	return dataservice.InsertContact(m)

}

func SendMessage(account, groupname, text string) error {

	client := ClientManager.TdInstances[account]
	if client == nil {
		return errors.New("找不到用户")
	}

	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	chat, err := client.SearchChatsOnServer(groupname, 1)

	if err != nil {
		return errors.New("找不到组！")
	}
	if chat.TotalCount == 0 {
		log.Info("找不到组, ", groupname)
		return errors.New("找不到组！")
	}

	inputMsgTxt := tdlib.NewInputMessageText(tdlib.NewFormattedText(text, nil), true, true)
	_, err = client.SendMessage(chat.ChatIDs[0], int64(0), int64(0), nil, nil, inputMsgTxt)
	if err != nil {
		return errors.New("发送失败")
	}

	return nil
}

// var allChats []*tdlib.Chat
// var haveFullChatList bool

// see https://stackoverflow.com/questions/37782348/how-to-use-getchats-in-tdlib
func getChatList(client *tdlib.Client, limit int, haveFullChatList bool) ([]*tdlib.Chat, error) {

	var allChats []*tdlib.Chat

	if !haveFullChatList && limit > len(allChats) {
		offsetOrder := int64(math.MaxInt64)
		offsetChatID := int64(0)
		var chatList = tdlib.NewChatListMain()
		var lastChat *tdlib.Chat

		if len(allChats) > 0 {
			lastChat = allChats[len(allChats)-1]
			for i := 0; i < len(lastChat.Positions); i++ {
				//Find the main chat list
				if lastChat.Positions[i].List.GetChatListEnum() == tdlib.ChatListMainType {
					offsetOrder = int64(lastChat.Positions[i].Order)
				}
			}
			offsetChatID = lastChat.ID
		}

		// get chats (ids) from tdlib
		chats, err := client.GetChats(chatList, tdlib.JSONInt64(offsetOrder),
			offsetChatID, int32(limit-len(allChats)))
		if err != nil {
			return nil, err
		}
		if len(chats.ChatIDs) == 0 {
			haveFullChatList = true
			return allChats, nil
		}

		for _, chatID := range chats.ChatIDs {
			// get chat info from tdlib
			chat, err := client.GetChat(chatID)
			if err == nil {
				allChats = append(allChats, chat)
			} else {
				return nil, err
			}
		}
		// return getChatList(client, limit, allChats, false)
	}
	return allChats, nil
}

//保存组联系人
func SaveGroupContents(phone string, superid int32) error {
	client, _ := ClientManager.GetTdClient(phone)

	if client == nil {
		return errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()

	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	var filter tdlib.SupergroupMembersFilter

	memer, err := client.GetSupergroupMembers(int32(superid), filter, 0, 100)
	fmt.Println(memer, err)
	defer client.Close()

	return nil
}

//保存组联系人
func Savechatcontacts(phone string, chatid int64) error {
	client, _ := ClientManager.GetTdClient(phone)

	if client == nil {
		return errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()

	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	// chats, _ := getChatList(client, 1, false)

	// fmt.Println(chats[0].ID)

	// getChatMembers(client, chatid)

	defer client.Close()

	return nil
}

func getChatMembers(client *tdlib.Client, chatID int64) {
	str := `0123456789abcdefghijklmnopqrstuvwxyz`
	var filter tdlib.ChatMembersFilter
	for _, ch := range str {
		searchStr := fmt.Sprintf("%c", ch)

		m, err := client.SearchChatMembers(chatID, searchStr, limit, filter)
		if err != nil {

			continue
		}

		if err = addMembers(client, m); err != nil {

		}
		time.Sleep(time.Duration(len(m.Members)/60+1) * time.Second)
	}
}

func addMembers(client *tdlib.Client, m *tdlib.ChatMembers) error {
	if m.TotalCount == 0 || len(m.Members) == 0 {

		return errors.New("no members")
	}
	fmt.Println("total count:", m.TotalCount, ", got member count:", len(m.Members))
	// for _, member := range m.Members {
	// 	user, err := client.GetUser(member.UserID)
	// 	if err != nil {

	// 		continue
	// 	}
	// 	if userExists(member.UserID) {
	// 		continue
	// 	}
	// 	// if err := insertUser(member.UserID, user.Username,); err != nil {
	// 	// 	// log.Println("insert user failed", err)
	// 	// }
	// }
	return nil
}

func userExists(userID int32) bool {
	return dataservice.Existcontacts(userID)
}

func insertUser(userID int32, userName, phone string) error {
	u := &model.Groupcontacts{Cid: userID, Cname: userName}
	return dataservice.SaveGroupcontacts(u)

}

func getSupergroupMemebers(client *tdlib.Client, supergroupID int32) {
	var offset int32 = 0
	var filter tdlib.SupergroupMembersFilter
	for ; ; offset += limit {
		m, err := client.GetSupergroupMembers(supergroupID, filter, offset, limit)
		if err != nil {
			// log.Println("getting supergroup member failed", err)
			break
		}

		if err = addMembers(client, m); err != nil {
			// log.Println("", err)
			break
		}
		time.Sleep(time.Duration(len(m.Members)/60+1) * time.Second)
	}
}

func (c *ClientManagerUseCase) GetTdClient(account string) (*tdlib.Client, int) {

	tdlib.SetLogVerbosityLevel(1)
	client := c.TdInstances[account]
	if client != nil {
		return client, REG_OK
	}

	// Create new instance of client
	client = tdlib.NewClient(tdlib.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tddata + account + "-tdlib-db",
		FileDirectory:       tdfile + account + "-tdlib-files",
		IgnoreFileNames:     false,
	})

	return client, model.AuthWaitCode
}

func SendPhoneNumber(account string, client *tdlib.Client) bool {
	defer client.Close()

	for {
		currentState, err := client.Authorize()
		if err != nil {
			log.Infof("Error getting current state: %s %v", account, err)
			break
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			_, err := client.SendPhoneNumber(account)
			if err != nil {
				log.Infof("phone %s  err: %v", account, err)
			}

			return true

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			return false

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			return true

		}
	}
	return false
}

func RegisterClient(account, code string, client *tdlib.Client) (model.Client, int) {
	for {
		currentState, err := client.Authorize()
		if err != nil {
			log.Infof("Error getting current state: %s %v", account, err)
			break
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			return model.Client{}, model.AuthSendTimeout

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			if code != "" {
				_, err := client.SendAuthCode(code)
				if err != nil {
					fmt.Printf("Error sending auth code : %v", err)
					return model.Client{}, model.AuthSenCodeErr
				}

			} else {
				return model.Client{}, model.AuthSenCodeErr
			}

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			var repclient model.Client
			user, _ := client.GetMe()
			repclient.Id = user.ID
			repclient.Name = user.FirstName
			repclient.Username = user.Username
			repclient.PhoneNumber = user.PhoneNumber
			ClientManager.AddClient(account, client)
			return repclient, model.SOK

		}
	}
	return model.Client{}, model.AuthSenCodeErr
}

func (c *ClientManagerUseCase) AddClient(account string, client *tdlib.Client) {
	dataservice.UpdateClient(account, REG_OK)
	c.TdInstances[account] = client
}

func insertUserIfNotExists(from string, user *tdlib.User) error {

	if userExists(user.ID) {
		return errors.New("User exists")
	}

	if err := insertUser(user.ID, user.FirstName, from); err != nil {
		// log.Println("insert user failed", err)
		return err
	}
	return nil
}

func insertUserIfNotExistsUserID(from string, user *tdlib.User) error {

	if userExists(user.ID) {
		return errors.New("User exists")
	}

	if err := insertUser(user.ID, user.FirstName, from); err != nil {
		// log.Println("insert user failed", err)
		return err
	}
	return nil
}
