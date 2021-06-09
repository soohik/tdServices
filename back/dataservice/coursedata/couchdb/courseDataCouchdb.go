package couchdb

import (
	"tdapi/model"
	"tdapi/tool/gdbc"
)

const (
	QUERY_USER = "SELECT * FROM course"
)

// CourseDataCouchdb is the NoSQL implementation of CourseDatainterface
type CourseDataCouchdb struct {
	DB gdbc.NoSqlGdbc
	//DB string
}

func (cdc *CourseDataCouchdb) FindAll() ([]model.Course, error) {
	return nil, nil
}

func (cdc *CourseDataCouchdb) SetDB(gdbc gdbc.Gdbc) {
	cdc.DB = gdbc
}
