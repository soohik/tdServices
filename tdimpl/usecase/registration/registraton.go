// Package registration represents the concrete implementation of RegistrationUseCaseInterface interface.
// Because the same business function can be created to support both transaction and non-transaction,
// a shared business function is created in a helper file, then we can wrap that function with transaction
// or non-transaction.
package registration

import (
	"tdimpl/dataservice"
	"tdimpl/model"
)

// RegistrationUseCase implements RegistrationUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
// TxDataInterface is needed to support transaction
type RegistrationUseCase struct {
	UserDataInterface dataservice.UserDataInterface
	TxDataInterface   dataservice.TxDataInterface
}

func (ruc *RegistrationUseCase) RegisterPhone(user *model.Phone) (*model.Phone, error) {
	// err := user.Validate()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "user validation failed")
	// }
	// isDup, err := ruc.isDuplicate(user.Name)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	// if isDup {
	// 	return nil, errors.New("duplicate user for " + user.Name)
	// }
	// resultUser, err := ruc.UserDataInterface.Insert(user)

	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	return nil, nil
}

// UnregisterUser unregister a user from an application by user name, basically removing it from a database.
func (ruc *RegistrationUseCase) UnregisterUser(username string) error {
	return nil
}

// The use case of ModifyAndUnregister without transaction
func (ruc *RegistrationUseCase) ModifyAndUnregister(user *model.Phone) error {
	return modifyAndUnregister(ruc, user)
}

func (ruc *RegistrationUseCase) EnableTx() {
	// Only UserDataInterface need transaction support here. If there are other data services need it,
	// then they also need to enable transaction here
	ruc.UserDataInterface.EnableTx(ruc.TxDataInterface)
}
