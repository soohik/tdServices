package clientmanager

import (
	"fmt"

	"github.com/Arman92/go-tdlib"
)

const (
	TdlibDbDirectory    = "./tddata/"
	TdlibFilesDirectory = "./tdfile/"
)

// RegistrationUseCase implements RegistrationUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
// TxDataInterface is needed to support transaction
type TDClient struct {
	AcountName          string
	Code                chan string
	Registered          int //0 未验证， 1 验证通过
	Runed               bool
	Status              int
	TdlibDbDirectory    string
	TdlibFilesDirectory string

	tdlibClient *tdlib.Client
	Manager     *ClientManagerUseCase
}

func (td TDClient) AddInstance() {
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
		DatabaseDirectory:   td.TdlibDbDirectory,
		FileDirectory:       td.TdlibFilesDirectory,
		IgnoreFileNames:     false,
	})

	td.Runed = true

	td.tdlibClient = client
	td.Code = make(chan string)

	// Main loop
	go func() {
		// Just fetch updates so the Updates channel won't cause the main routine to get blocked
		rawUpdates := client.GetRawUpdatesChannel(100)
		for update := range rawUpdates {
			// Show all updates
			_ = update
			if !td.Runed {
				client.Close()
				break
			}
			// fmt.Println(update.Data)
			// fmt.Print("\n\n")
		}
	}()

	// currentState, _ := client.Authorize()
	// for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
	// 	time.Sleep(300 * time.Millisecond)
	// }

	for {
		currentState, err := client.Authorize()
		if err != nil {
			fmt.Printf("Error getting current state: %v", err)
			continue
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			_, err := client.SendPhoneNumber(td.AcountName)
			if err != nil {
				go td.Manager.RemoveClient(td.AcountName)
			}

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			data := <-td.Code //接收
			_, err := client.SendAuthCode(data)
			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready! Let's rock")
			break
		}
	}

}
