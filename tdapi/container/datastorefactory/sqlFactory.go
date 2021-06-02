package datastorefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/container/logger"
	"tdapi/tool/gdbc/databasehandler"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// sqlFactory is receiver for Build method
type sqlFactory struct{}

// implement Build method for SQL database
func (sf *sqlFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	logger.Log.Debug("sqlFactory")
	key := dsc.Code
	//if it is already in container, return
	if value, found := c.Get(key); found {
		//logger.Log.Debugf("found db value %+v\n:", value)
		sdb := value.(*gorm.DB)
		sdt := databasehandler.SqlDBTx{DB: sdb}
		logger.Log.Debug("found db in container for key:", key)
		return &sdt, nil
	}

	// db, err := sql.Open(dsc.DriverName, dsc.UrlAddress)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }

	db, err := gorm.Open(mysql.Open(dsc.UrlAddress), &gorm.Config{})

	// check the connection
	sqlDB, err := db.DB()
	sqlDB.Ping()

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt := databasehandler.SqlDBTx{DB: db}
	c.Put(key, db)
	return &dt, nil

}
