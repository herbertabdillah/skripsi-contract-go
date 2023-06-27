package contract

import (
	"errors"
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func StartYear(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	_, err := cc.Repository.GetCourseYear(year, semester)
	if err == nil {
		return nil, errors.New("CousreYear exist")
	}

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
		dropOut(cc, year)
	}

	return cc.Repository.UpdateCourseYear(courseYear)
}

func GetCourseYear(c router.Context) (interface{}, error) {
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	var semesterNumber int
	if semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	key := "CourseYear." + strconv.Itoa(year) + strconv.Itoa(semesterNumber)

	return c.State().Get(key, &state.CourseYear{})
}

func InsertCourseSemester(c router.Context) (interface{}, error) {
	id, year, semester, courseId, lecturerId := c.ParamString("id"), c.ParamInt("year"), c.ParamString("semester"), c.ParamString("courseId"), c.ParamString("lecturerId")
	courseSemester := &state.CourseSemester{Id: id, Year: year, Semester: semester, CourseId: courseId, LecturerId: lecturerId}

	return courseSemester, c.State().Insert("CourseSemester."+id, courseSemester)
}

func GetCourseSemester(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("CourseSemester."+id, &state.CourseSemester{})
}

func dropOut(cc Context, year int) error {
	studentYear, _ := cc.Repository.GetStudentYear(year - 7)

	if studentYear == nil {
		return nil
	}

	for _, studentId := range studentYear.StudentIds {
		student, err := cc.Repository.GetStudent(studentId)
		if err != nil {
			return err
		}

		if student.Status != "active" {
			continue
		}

		student.Status = "drop_out"
		_, err = cc.Repository.UpdateStudent(student)

		if err != nil {
			return err
		}
	}

	return nil
}
