package dataservice

import (
	"tdapi/config"
	"tdapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// sqlFactory is receiver for Build method
type sqlFactory struct {
	sqldb *gorm.DB
}

const (
	FIND_PHONE = true
)

var sqlHelp sqlFactory

// implement Build method for SQL database
func SqlBuild() error {

	config := config.Config.SQLConfig

	db, err := gorm.Open(mysql.Open(config.UrlAddress), &gorm.Config{})
	if err != nil {
		return nil
	}
	// check the connection
	sqlDB, err := db.DB()
	sqlDB.Ping()

	sqlHelp.sqldb = db
	return err

}

func Preregister(p model.Phone) bool {
	// db := sqlHelp.sqldb
	phone, find := sqlHelp.GetPhone(p.Phone)

	if find == true && phone.Registered == 1 {

		return true
	}

	return InsertClient(p)

}

func (db *sqlFactory) GetPhone(phonenumber string) (model.Phone, bool) {
	var phone model.Phone
	result := db.sqldb.Where("phone = ?", phonenumber).First(&phone)
	if result.RowsAffected <= 0 {
		return phone, false
	}

	return phone, true

}

func GetAllPhone() ([]model.Phone, error) {
	var phones []model.Phone
	rows, err := sqlHelp.sqldb.Raw("select * from td.phones").Rows()

	defer rows.Close()
	if err != nil {

		return nil, nil
	}
	for rows.Next() {
		var phone model.Phone
		// ScanRows 方法用于将一行记录扫描至结构体
		sqlHelp.sqldb.ScanRows(rows, &phone)
		phones = append(phones, phone)
	}
	return phones, nil
}

func InsertClient(p model.Phone) bool {

	result := sqlHelp.sqldb.Create(&p)

	if result.Error != nil {

		return false
	}

	return true
}
