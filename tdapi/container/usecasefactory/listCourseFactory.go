package usecasefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/usecase/listcourse"

	"github.com/pkg/errors"
)

type ListCourseFactory struct{}

func (lcf *ListCourseFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.ListCourse
	cdi, err := buildCourseData(c, &uc.CourseDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	lcuc := listcourse.ListCourseUseCase{CourseDataInterface: cdi}
	return &lcuc, nil
}
