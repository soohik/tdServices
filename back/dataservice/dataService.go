// Package dataservice and it's sub-package represents data persistence service, mainly access database,
// but also including data persisted by other Micro-service.
// For Micro-server, only the interface is defined in this package, the data transformation code is in adapter package
// This is the top level package and it only defines interface, and all implementations are defined in sub-package
// Use case package depends on it.
package dataservice

import (
	"tdapi/config"
	"tdapi/model"
	"tdapi/tool/gdbc"
)

// UserDataInterface represents interface for user data access through database
type ListDataInterface interface {
	// Remove deletes a user by user name from database.
	Remove(username string) (rowsAffected int64, err error)
	// Find retrieves a user from database based on a user's id
	Find(id int) (*model.User, error)
	// FindByName retrieves a user from database by User.Name
	FindByName(name string) (user *model.User, err error)
	// FindAll retrieves all users from database as an array of user
	FindAll() ([]model.User, error)
	// Update changes user information on the User.Id passed in.
	Update(user *model.User) (rowsAffected int64, err error)
	// Insert adds a user to a database. The returned resultUser has a Id, which is auto generated by database
	Insert(user *model.User) (resultUser *model.User, err error)
	// Need to add this for transaction support
	EnableTxer
}

// UserDataInterface represents interface for user data access through database
type UserDataInterface interface {
	// Remove deletes a user by user name from database.
	Remove(username string) (rowsAffected int64, err error)
	// Find retrieves a user from database based on a user's id
	Find(id int) (*model.User, error)
	// FindByName retrieves a user from database by User.Name
	FindByName(name string) (user *model.User, err error)
	// FindAll retrieves all users from database as an array of user
	FindAll() ([]model.User, error)
	// Update changes user information on the User.Id passed in.
	Update(user *model.User) (rowsAffected int64, err error)
	// Insert adds a user to a database. The returned resultUser has a Id, which is auto generated by database
	Insert(user *model.User) (resultUser *model.User, err error)
	// get all user  from database.
	GetAllClient() ([]model.Phone, error)
	// Need to add this for transaction support
	EnableTxer

	InitUseCase
}

// CacheDataInterface represents interface for cache service, which is a micro-service
type CacheDataInterface interface {
	// Get handles call to Get function on Cache service
	Get(key string) ([]byte, error)
	// Store handles call to Store function on Cache service
	Store(key string, value []byte) error
}

// CourseDataInterface represents interface for persistence service for course data
// It is created for POC of courseDataServiceFactory, no real use.
type CourseDataInterface interface {
	FindAll() ([]model.Course, error)
	SetDB(gdbc gdbc.Gdbc)
}

// This interface needs to be included in every data service interface that needs transaction support
type InitUseCase interface {
	// EnableTx enables transaction, basically it replaces the underling database handle sql.DB with sql.Tx
	InitUseCase(appConfig *config.AppConfig)
}

// This interface needs to be included in every data service interface that needs transaction support
type EnableTxer interface {
	// EnableTx enables transaction, basically it replaces the underling database handle sql.DB with sql.Tx
	EnableTx(dataInterface TxDataInterface)
}

// TxDataInterface represents operations needed for transaction support.
// It only needs to be implemented once for each database
// For sqlGdbc, it is implemented for SqlDBTx in transaction.go
type TxDataInterface interface {
	// TxBegin starts a transaction. It gets a DB handler from the receiver and return a TxDataInterface, which has a
	// *sql.Tx inside. Any data access wrapped inside a transaction will go through the *sql.Tx
	TxBegin() (TxDataInterface, error)
	// TxEnd is called at the end of a transaction and based on whether there is an error, it commits or rollback the
	// transaction.
	// txFunc is the business function wrapped in a transaction
	TxEnd(txFunc func() error) error
	// Return the underline transaction handler, sql.Tx
	GetTx() gdbc.SqlGdbc
}