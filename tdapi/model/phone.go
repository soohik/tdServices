// Package model represents domain model. Every domain model type should have it's own file.
// It shouldn't depends on any other package in the application.
// It should only has domain model type and limited domain logic, in this example, validation logic. Because all other
// package depends on this package, the import of this package should be as small as possible.

package model

type Invated struct {
	Phone  string  `json:"phone"`
	Cids   []int32 `json:"cids"`
	Chatid int64     `json:"Chatid"` //组ID
	Uname  string  `json:"name"`   //linkurl地址
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type Linkurl struct {
	Phone string `json:"phone"`
	Url   string `json:"url"`
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type RegPhone struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type Phone struct {
	Phone      string `json:"phone"`
	Account    string `json:"account"`
	Tddata     string `json:"tddata"`
	Tdfile     string `json:"tdfile"`
	Registered int    `json:"registered"`
	Createtime string `json:"create_time" gorm:"column:create_time;default:now()"`
	Updatetime string `json:"update_time" gorm:"column:update_time;default:now()"`
}

// Validate validates a newly created user, which has not persisted to database yet, so Id is empty
func (u Phone) Validate() error {
	return nil
	// return validation.ValidateStruct(&u,
	// 	validation.Field(&u.Name, validation.Required),
	// 	validation.Field(&u.Created, validation.Required))
}

//ValidatePersisted validate a user that has been persisted to database, basically Id is not empty
func (u Phone) ValidatePersisted() error {
	return nil
}
