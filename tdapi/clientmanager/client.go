package clientmanager

import (
	"fmt"

	"github.com/Arman92/go-tdlib"
)

const (
	TdlibDbDirectory    = "../../tddata/"
	TdlibFilesDirectory = "../../tdfile/"
)

// RegistrationUseCase implements RegistrationUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
// TxDataInterface is needed to support transaction
type TDClient struct {
	AcountName          string
	TdlibDbDirectory    string
	TdlibFilesDirectory string
	tdlibClient         *tdlib.Client
}

func (td TDClient) AddInstance() {
	//tdlib.SetLogVerbosityLevel(1)
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

	td.tdlibClient = client

	// Main loop
	go func() {
		// Just fetch updates so the Updates channel won't cause the main routine to get blocked
		rawUpdates := client.GetRawUpdatesChannel(100)
		for update := range rawUpdates {
			// Show all updates
			_ = update
			// fmt.Println(update.Data)
			// fmt.Print("\n\n")
		}
	}()

	for {
		currentState, err := client.Authorize()
		if err != nil {
			fmt.Printf("Error getting current state: %v", err)
			continue

		}
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready! Let's rock")
			break
		}
		fmt.Println(currentState)
	}

}
