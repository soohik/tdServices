package dataservicefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/container/datastorefactory"
	"tdapi/container/logger"
	"tdapi/dataservice"
	"tdapi/dataservice/coursedata/couchdb"
	"tdapi/dataservice/coursedata/sqldb"
	"tdapi/tool/gdbc"

	"github.com/pkg/errors"
)

var courseDataServiceMap = map[string]dataservice.CourseDataInterface{
	config.COUCHDB: &couchdb.CourseDataCouchdb{},
	config.SQLDB:   &sqldb.CourseDataSql{},
}

// courseDataServiceFactory is an empty receiver for Build method
type courseDataServiceFactory struct{}

// GetCourseDataServiceInterface is an accessor for factoryBuilderMap
func GetCourseDataServiceInterface(key string) dataservice.CourseDataInterface {
	return courseDataServiceMap[key]
}

func (tdsf *courseDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("courseDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	gdbc := dsi.(gdbc.Gdbc)
	gdi := GetCourseDataServiceInterface(dsc.Code)
	gdi.SetDB(gdbc)
	return gdi, nil
}
