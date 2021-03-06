package listcourse

import (
	"tdapi/dataservice"
	"tdapi/model"
)

// ListCourseUseCase implements ListCourseUseCaseInterface.
type ListCourseUseCase struct {
	// CourseDataInterface, which is a interface to underline database connection and can be used to access
	// persistence layer
	CourseDataInterface dataservice.CourseDataInterface
}

func (luc *ListCourseUseCase) ListCourse() ([]model.Course, error) {
	return luc.CourseDataInterface.FindAll()
}
