package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func StartYear(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	appConfig, err := cc.Repository.GetApplicationConfig()
	if err != nil {
		return nil, err
	}

	appConfig.Semester = semester
	appConfig.Year = year

	_, err = cc.Repository.UpdateApplicationConfig(appConfig)

	if err != nil {
		return nil, err
	}

	courseYear := &state.CourseYear{Year: year, Semester: semester, Status: "start"}

	return cc.Repository.InsertCourseYear(courseYear)
}

func EndYear(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	courseYear, err := cc.Repository.GetCourseYear(year, semester)
	if err != nil {
		return nil, err
	}

	courseYear.Status = "end"

	if semester == "odd" {
		err = cc.Service.DroupOut(year)

		if err != nil {
			return nil, err
		}
	}

	return cc.Repository.UpdateCourseYear(courseYear)
}

func GetCourseYear(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	return cc.Repository.GetCourseYear(year, semester)
}

func InsertCourseSemester(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id, year, semester, courseId, lecturerId := c.ParamString("id"), c.ParamInt("year"), c.ParamString("semester"), c.ParamString("courseId"), c.ParamString("lecturerId")
	courseSemester := &state.CourseSemester{Id: id, Year: year, Semester: semester, CourseId: courseId, LecturerId: lecturerId, StudentCount: 0}

	return cc.Repository.InsertCourseSemester(courseSemester)
}

func GetCourseSemester(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Repository.GetCourseSemester(id)
}
