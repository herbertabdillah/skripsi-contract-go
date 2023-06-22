package contract

import (
	"errors"
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func StartYear(c router.Context) (interface{}, error) {
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	var semesterNumber int
	if semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	_, err := c.State().Get("CourseYear."+strconv.Itoa(year)+strconv.Itoa(semesterNumber), &state.CourseYear{})
	if err == nil {
		return nil, errors.New("CousreYear exist")
	}

	courseYear := &state.CourseYear{Year: year, Semester: semester, Status: "start"}
	return courseYear, c.State().Insert("CourseYear."+strconv.Itoa(year)+strconv.Itoa(semesterNumber), courseYear)
}

func EndYear(c router.Context) (interface{}, error) {
	year, semester := c.ParamInt("year"), c.ParamString("semester")

	var semesterNumber int
	if semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	key := "CourseYear." + strconv.Itoa(year) + strconv.Itoa(semesterNumber)

	res, err := c.State().Get(key, &state.CourseYear{})
	courseYear := res.(state.CourseYear)
	if err != nil {
		return nil, err
	}

	courseYear.Status = "end"

	if semester == "odd" {
		studentYearRes, _ := c.State().Get("StudentYear."+strconv.Itoa(year-7), &state.StudentYear{})
		if studentYearRes != nil {
			studentYear := studentYearRes.(state.StudentYear)
			for _, studentId := range studentYear.StudentIds {
				studentRes, err := c.State().Get("Student."+studentId, &state.Student{})
				if err != nil {
					return nil, err
				}
				student := studentRes.(state.Student)
				if student.Status == "active" {
					student.Status = "drop_out"
					c.State().Put("Student."+studentId, student)
				}
			}
		}
	}

	return courseYear, c.State().Put(key, courseYear)
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
