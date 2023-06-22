package contract

import (
	"errors"

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

func Graduate(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	studentRes, err := c.State().Get("Student."+id, &state.Student{})
	if err != nil {
		return nil, err
	}
	student := studentRes.(state.Student)

	departmentRes, err := c.State().Get("Department."+student.DepartmentId, &state.Department{})
	if err != nil {
		return nil, err
	}
	department := departmentRes.(state.Department)
	courseIds := department.CourseIds

	// transcript := transcriptRes.(state.Transcript)
	transcriptRes, err := c.State().Get("Transcript."+id, &state.Transcript{})
	if err != nil {
		return nil, err
	}
	transcript := transcriptRes.(state.Transcript)

	for _, transcriptResult := range transcript.TranscriptResult {
		if transcriptResult.Pass {
			for j, courseId := range courseIds {
				if transcriptResult.CourseId == courseId {
					copy(courseIds[j:], courseIds[j+1:])     // Shift a[i+1:] left one index.
					courseIds[len(courseIds)-1] = ""         // Erase last element (write zero value).
					courseIds = courseIds[:len(courseIds)-1] // Truncate slice.
				}
			}

		}
	}

	if len(courseIds) == 0 {
		student.Status = "graduated"
		err = c.State().Put("Student."+id, student)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("haven't done all course")
	}
	return student, nil
}

func GetTranscript(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("Transcript."+id, &state.Transcript{})
}
