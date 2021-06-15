// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"errors"
	"fmt"
	"math"
	"tdapi/dataservice"
	"tdapi/log"
	"tdapi/model"
	"time"

	"github.com/Arman92/go-tdlib"
)

const (
	TDURL = "https://t.me/"
)

const (
	tddata = "../../tddata/"
	tdfile = "../../tdfile/"
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

func RegisterPhone(phonenumber, logincode string) (bool, model.Client) {

	_, find := ClientManager.GetClient(phonenumber)
	if !find {
		return false, model.Client{}
	}
	client, ret := ClientManager.AddInstance(phonenumber, logincode)
	if ret == model.SOK {
		return true, client
	}
	return false, model.Client{}

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
		ClientManager.AddInstance(phonenumber, "")

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
		c.AddInstance(value.Account, "")
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

func (c *ClientManagerUseCase) AddInstance(account, code string) (model.Client, int) {
	var ret int
	var repclient model.Client

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

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			if code != "" {
				_, err := client.SendAuthCode(code)
				if err != nil {
					fmt.Printf("Error sending auth code : %v", err)
				}
			} else {
				return repclient, model.AuthWaitCode
			}

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			log.Info("验证通过-- ", account)
			user, _ := client.GetMe()
			repclient.Id = user.ID
			repclient.Name = user.FirstName
			repclient.Username = user.Username
			repclient.PhoneNumber = user.PhoneNumber
			ret = model.SOK

			c.TdInstances[account] = client
			break

		}
	}

	// Main loop
	go func() {

		// Just fetch updates so the Updates channel won't cause the main routine to get blocked
		rawUpdates := client.GetRawUpdatesChannel(1000)

		for update := range rawUpdates {

			// fmt.Println(update)
			_ = update.Data
			// fmt.Println(m)
			// if m["@type"] == "updateNewChat" {
			// 	var update tdlib.UpdateNewChat
			// 	err = json.Unmarshal(result.Raw, &update)
			// 	return &update, err
			// }
			// if m["@type"] == "updateBasicGroup" ||
			// 	m["@type"] == "updateSupergroup" {
			// 	fmt.Println(m)

			// }

			// if m["@type"] == "updateBasicGroupFullInfo" ||
			// 	m["@type"] == "updateSupergroupFullInfo" {
			// 	fmt.Println(m["@type"])
			// 	fmt.Println(m)
			// }

		}
	}()

	return repclient, ret
}

func Getallgroups(account string, agent int) ([]model.Groups, error) {
	// return dataservice.GetAllGroups(agent)
	var groups []model.Groups

	client := ClientManager.TdInstances[account]
	if client == nil {
		return nil, errors.New("找不到账号！")
	}
	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	chats, err := getChatList(client, 2, false)
	if err != nil {
		return nil, err
	}
	for _, value := range chats {
		var m model.Groups
		m.Uid = value.ID
		m.Name = value.Title
		groups = append(groups, m)
	}

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
