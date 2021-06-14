package model

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type AddContacts struct {
	Phone    string   `json:"phone"`
	Contents []string `json:"contents"`
}

type Agent struct {
	Name  string `json:"name"`
	Agent int    `json:"agent"`
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type Text struct {
	Name  string `json:"name"`
	Agent int    `json:"agent"`
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type Me struct {
	Name  string `json:"name"`
	Agent int    `json:"agent"`
}

// ALTER TABLE `td`.`groups`
// ADD COLUMN `verified` INT NULL DEFAULT 0 COMMENT '1 需要验证' AFTER `linkurl`;

type Groups struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Linkurl  string `json:"linkurl"`
	Agent    int    `json:"agent"`
	Verified int    `json:"Verified"`
}

// User has a name, department and created date. Name and created are required, department is optional.
// Id is auto-generated by database after the user is persisted.
// json is for couchdb
type Groupinfos struct {
	Phone     string `json:"phone"`
	Uid       string `json:"uid"`
	Groupname string `json:"groupname"`
	Agent     int    `json:"agent"`
}

type Friends struct {
	Cids  []int32 `json:"cids"`
	Uname string  `json:"name"`
}

type Contacts struct {
	Account      string `json:"account"`
	Contactid    int    `json:"contactid"`
	Contactphone string `json:"contactphone"`
	Contactname  string `json:"contactname"`
	Status       string `json:"status"`
}

// SELECT `taskinfo`.`tid`,
//     `taskinfo`.`account`,
//     `taskinfo`.`groupid`,
//     `taskinfo`.`groupname`,
//     `taskinfo`.`counts`,
//     `taskinfo`.`cron`,
//     `taskinfo`.`cycle`,
//     `taskinfo`.`text`,
//     `taskinfo`.`createtime`
// FROM `td`.`taskinfo`;

type Taskinfo struct {
	Tid       int    `json:"tid"`
	Account   string `json:"account"`
	Groupid   string `json:"groupid"`
	Groupname string `json:"groupname"`
	Counts    int    `json:"counts"`
	Cron      int    `json:"cron"`
	Cycle     int    `json:"cycle"`
	Text      string `json:"text"`
}

//任务记录
type Tasklog struct {
	Tid      int    `json:"tid"`
	Counts   int    `json:"counts"`
	Countsed int    `json:"Countsed"`
	Status   int    `json:"status"`
	Operid   int    `json:"operid"`
	Logs     string `json:"Logs"`
}
