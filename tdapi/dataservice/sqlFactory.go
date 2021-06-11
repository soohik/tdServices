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
	_, find := GetPhone(p.Phone)

	if find {
		return true
	}

	return InsertClient(p)

}

func GetPhone(phonenumber string) (model.Phone, bool) {
	var phone model.Phone
	result := sqlHelp.sqldb.Where("phone = ?", phonenumber).First(&phone)
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

func InsertGroup(groupname, linkurl string) bool {

	var group model.Groups
	result := sqlHelp.sqldb.Where("name = ? and linkurl = ?", groupname, linkurl).First(&group)

	if result.RowsAffected > 0 {
		return true
	}
	group.Name = groupname
	group.Linkurl = linkurl
	sqlHelp.sqldb.Create(&group)

	return true

}

func InsertGroupsInfo(phone, groupname string) bool {

	var group model.Groupinfos
	result := sqlHelp.sqldb.Where("phone = ? and groupname = ?", phone, groupname).First(&group)

	if result.RowsAffected > 0 {
		return true
	}
	group.Phone = phone
	group.Groupname = groupname

	sqlHelp.sqldb.Create(&group)
	return true

}

func GetAllGroups(agent int) ([]model.Groups, error) {
	var groups []model.Groups
	rows, err := sqlHelp.sqldb.Raw("select * from td.groups where agent = ?", agent).Rows()
	if err == nil {
		defer rows.Close()
	}

	if err != nil {

		return nil, nil
	}
	for rows.Next() {
		var group model.Groups
		// ScanRows 方法用于将一行记录扫描至结构体
		sqlHelp.sqldb.ScanRows(rows, &group)
		groups = append(groups, group)
	}
	return groups, nil
}

func GetMeGroups(phone string) ([]model.Groupinfos, error) {
	var groups []model.Groupinfos
	rows, err := sqlHelp.sqldb.Raw("select * from td.groupinfos  where phone = ?", phone).Rows()
	if err == nil {
		defer rows.Close()
	}

	if err != nil {

		return nil, nil
	}
	for rows.Next() {
		var group model.Groupinfos
		// ScanRows 方法用于将一行记录扫描至结构体
		sqlHelp.sqldb.ScanRows(rows, &group)
		groups = append(groups, group)
	}
	return groups, nil
}

func InsertContact([]model.AddContacts) (model.Phone, bool) {

	var phone model.Phone
	// result := sqlHelp.sqldb.Where("phone = ?", phonenumber).First(&phone)
	// if result.RowsAffected <= 0 {
	// 	return phone, false
	// }

	return phone, true
}
