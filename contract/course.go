package contract

import (
	"encoding/json"
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func InsertCoursePlan(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id, year, semester, studentId, status, courseSemesterIdsRaw := c.ParamString("id"), c.ParamInt("year"), c.ParamString("semester"), c.ParamString("studentId"), c.ParamString("status"), c.ParamString("courseSemesterIds")

	courseSemesterIds := []string{}
	json.Unmarshal([]byte(courseSemesterIdsRaw), &courseSemesterIds)
	coursePlan := &state.CoursePlan{Id: id, Year: year, Semester: semester, StudentId: studentId, Status: status, CourseSemesterIds: courseSemesterIds}

	return cc.Service.InsertCoursePlan(coursePlan)
}

func UpdateCourseResult(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	courseSemesterId, courseResultId, scoreStr := c.ParamString("courseSemesterId"), c.ParamString("courseResultId"), c.ParamString("score")
	score, _ := strconv.ParseFloat(scoreStr, 64)

	courseResult, err := cc.Repository.GetCourseResult(courseResultId)
	if err != nil {
		return nil, err
	}

	return cc.Service.UpdateCourseResult(courseResult, courseSemesterId, score)
}

func GetCoursePlan(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Repository.GetCoursePlan(id)
}

func GetCourseResult(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	id := c.ParamString("id")

	return cc.Repository.GetCourseResult(id)
}
