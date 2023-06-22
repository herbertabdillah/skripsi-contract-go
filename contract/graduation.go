package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

// func InsertCoursePlan(c router.Context) (interface{}, error) {
// 	id, year, semester, studentId, status, courseSemesterIdsRaw := c.ParamString("id"), c.ParamInt("year"), c.ParamString("semester"), c.ParamString("studentId"), c.ParamString("status"), c.ParamString("courseSemesterIds")
// 	courseSemesterIds := []string{}
// 	json.Unmarshal([]byte(courseSemesterIdsRaw), &courseSemesterIds)
// 	coursePlan := &state.CoursePlan{Id: id, Year: year, Semester: semester, StudentId: studentId, Status: status, CourseSemesterIds: courseSemesterIds}

// 	var err = c.State().Insert("CoursePlan."+id, coursePlan)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := []state.CourseSemeterResult{}

// 	for _, courseSemesterId := range courseSemesterIds {
// 		csr := state.CourseSemeterResult{CourseSemesterId: courseSemesterId, Score: 0, Pass: false}
// 		result = append(result, csr)
// 	}

// 	courseResult := &state.CourseResult{Id: id, Year: year, Semester: semester, StudentId: studentId, CoursePlanId: id, Result: result}

// 	err = c.State().Insert("CourseResult."+id, courseResult)

// 	return coursePlan, err
// }

// func UpdateCourseResult(c router.Context) (interface{}, error) {
// 	courseSemesterId, courseResultId, scoreStr := c.ParamString("courseSemesterId"), c.ParamString("courseResultId"), c.ParamString("score")
// 	score, _ := strconv.ParseFloat(scoreStr, 64)
// 	result, err := c.State().Get("CourseResult."+courseResultId, &state.CourseResult{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	courseResult := result.(state.CourseResult)
// 	for i, courseSemesterResult := range courseResult.Result {
// 		if courseSemesterResult.CourseSemesterId == courseSemesterId {
// 			courseSemesterResult.Score = score
// 			if score >= 2 {
// 				courseSemesterResult.Pass = true
// 			} else {
// 				courseSemesterResult.Pass = false
// 			}
// 			courseResult.Result[i] = courseSemesterResult
// 		}
// 	}

// 	return courseResult, c.State().Put("CourseResult."+courseResultId, courseResult)
// }

func GetTranscript(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("Transcript."+id, &state.Transcript{})
}
