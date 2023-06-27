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

	if err := cc.Service.Updatable(year, semester); err != nil {
		return nil, err
	}

	courseSemesterIds := []string{}
	json.Unmarshal([]byte(courseSemesterIdsRaw), &courseSemesterIds)
	coursePlan := &state.CoursePlan{Id: id, Year: year, Semester: semester, StudentId: studentId, Status: status, CourseSemesterIds: courseSemesterIds}

	// _, err := cc.Repository.InsertCoursePlan(coursePlan)
	// if err != nil {
	// 	return nil, err
	// }

	// err = classQuotaValidation(cc, coursePlan)
	// if err != nil {
	// 	return nil, err
	// }

	// err = creditValidation(cc, coursePlan)
	// if err != nil {
	// 	return nil, err
	// }

	// _, err = generateCourseResult(cc, coursePlan)
	// if err != nil {
	// 	return nil, err
	// }

	// return coursePlan, err
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

	if err := cc.Service.Updatable(courseResult.Year, courseResult.Semester); err != nil {
		return nil, err
	}

	isPass := false
	if score >= 2 {
		isPass = true
	}

	_, course, err := updateCourseResult(cc, courseResult, courseSemesterId, score, isPass)
	if err != nil {
		return nil, err
	}

	_, err = updateTranscript(cc, courseResult, course.Id, score, isPass, course.Credit)
	if err != nil {
		return nil, err
	}

	return cc.Repository.UpdateCourseResult(courseResult)
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

// func classQuotaValidation(cc Context, cp *state.CoursePlan) error {
// 	for _, courseSemesterId := range cp.CourseSemesterIds {
// 		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterId)
// 		if err != nil {
// 			return err
// 		}

// 		if courseSemester.StudentCount >= config.MAX_STUDENT_PER_CLASS {
// 			return errors.New("class full")
// 		}

// 		courseSemester.StudentCount += 1

// 		_, err = cc.Repository.UpdateCourseSemester(courseSemester)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func creditValidation(cc Context, cp *state.CoursePlan) error {
// 	totalCredit := 0
// 	lastSemesterScore := 0.0
// 	maxCredit := config.MAX_CREDIT_PER_SEMESTER
// 	for _, courseSemesterId := range cp.CourseSemesterIds {
// 		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterId)
// 		if err != nil {
// 			return err
// 		}

// 		course, err := cc.Repository.GetCourse(courseSemester.CourseId)
// 		if err != nil {
// 			return err
// 		}

// 		totalCredit += course.Credit
// 	}

// 	transcript, err := cc.Repository.GetTranscript(cp.StudentId)
// 	if err != nil {
// 		return err
// 	}

// 	if transcript == nil || len(transcript.TranscriptResult) == 0 {
// 		lastSemesterScore = 4.0
// 	} else {
// 		lastSemesterScore = transcript.TranscriptResult[len(transcript.TranscriptResult)-1].Score
// 	}

// 	if lastSemesterScore < 3 {
// 		maxCredit = 21
// 	} else if lastSemesterScore < 2 {
// 		maxCredit = 18
// 	}

// 	if totalCredit >= maxCredit {
// 		return errors.New("credit exceed " + strconv.Itoa(totalCredit) + " max: " + strconv.Itoa(maxCredit))
// 	}

// 	return nil
// }

// func generateCourseResult(cc Context, cp *state.CoursePlan) (*state.CourseResult, error) {
// 	result := []state.CourseSemeterResult{}

// 	for _, courseSemesterId := range cp.CourseSemesterIds {
// 		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		csr := state.CourseSemeterResult{CourseSemesterId: courseSemesterId, Score: 0, Pass: false, CourseId: courseSemester.CourseId}
// 		result = append(result, csr)
// 	}

// 	courseResult := &state.CourseResult{Id: cp.Id, Year: cp.Year, Semester: cp.Semester, StudentId: cp.StudentId, CoursePlanId: cp.Id, Result: result}
// 	_, err := cc.Repository.InsertCourseResult(courseResult)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return courseResult, nil
// }

func updateCourseResult(cc Context, cr *state.CourseResult, courseSemesterId string, score float64, isPass bool) (*state.CourseResult, *state.Course, error) {
	var currentCourse *state.Course
	totalScore := 0.0
	totalCredit := 0

	for i, courseSemesterResult := range cr.Result {
		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterResult.CourseSemesterId)
		if err != nil {
			return nil, nil, err
		}

		course, err := cc.Repository.GetCourse(courseSemester.CourseId)
		if err != nil {
			return nil, nil, err
		}

		if courseSemesterResult.CourseSemesterId == courseSemesterId {
			currentCourse = course

			courseSemesterResult.Score = score
			courseSemesterResult.Pass = isPass
			cr.Result[i] = courseSemesterResult
		}
		totalScore += courseSemesterResult.Score * float64(course.Credit)
		totalCredit += course.Credit
	}

	resultScore := totalScore / float64(totalCredit)
	cr.Score = resultScore

	_, err := cc.Repository.UpdateCourseResult(cr)

	if err != nil {
		return nil, currentCourse, err
	}

	return cr, currentCourse, nil
}

func updateTranscript(cc Context, cr *state.CourseResult, currentCourseId string, score float64, isPass bool, credit int) (*state.Transcript, error) {
	transcript, err := cc.Repository.GetTranscript(cr.StudentId)
	if err != nil {
		return nil, err
	}

	found := false
	totalTranscriptScore := 0.0
	totalTranscriptCredit := 0
	foundIndex := -1
	var foundResult state.TranscriptResult
	for i, transcriptResult := range transcript.TranscriptResult {
		if transcriptResult.CourseId == currentCourseId {
			found = true
			foundIndex = i
			foundResult = transcriptResult

			foundResult.Score = score
			foundResult.Pass = isPass
			foundResult.CourseResultId = cr.Id
			foundResult.Year = cr.Year
			foundResult.Semester = cr.Semester

			continue
		}

		course, err := cc.Repository.GetCourse(transcriptResult.CourseId)
		if err != nil {
			return nil, err
		}

		totalTranscriptCredit += course.Credit
		totalTranscriptScore += (float64(course.Credit) * transcriptResult.Score)
	}

	if found {
		transcript.TranscriptResult[foundIndex] = foundResult
	} else {
		newTranscriptResult := state.TranscriptResult{CourseResultId: cr.Id, CourseId: currentCourseId, Year: cr.Year, Semester: cr.Semester, Score: score, Pass: isPass}
		transcript.TranscriptResult = append(transcript.TranscriptResult, newTranscriptResult)
	}

	totalTranscriptCredit += credit
	totalTranscriptScore += (float64(credit) * score)

	transcript.Score = totalTranscriptScore / float64(totalTranscriptCredit)

	_, err = cc.Repository.UpdateTranscript(transcript)
	if err != nil {
		return nil, err
	}

	return transcript, nil
}
