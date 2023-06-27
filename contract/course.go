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

	_, err := cc.Repository.InsertCoursePlan(coursePlan)
	if err != nil {
		return nil, err
	}

	result := []state.CourseSemeterResult{}

	for _, courseSemesterId := range courseSemesterIds {
		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterId)
		if err != nil {
			return nil, err
		}
		csr := state.CourseSemeterResult{CourseSemesterId: courseSemesterId, Score: 0, Pass: false, CourseId: courseSemester.CourseId}
		result = append(result, csr)
	}

	courseResult := &state.CourseResult{Id: id, Year: year, Semester: semester, StudentId: studentId, CoursePlanId: id, Result: result}
	_, err = cc.Repository.InsertCourseResult(courseResult)

	return coursePlan, err
}

func UpdateCourseResult(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var err error
	courseSemesterId, courseResultId, scoreStr := c.ParamString("courseSemesterId"), c.ParamString("courseResultId"), c.ParamString("score")
	score, _ := strconv.ParseFloat(scoreStr, 64)
	courseResult, err := cc.Repository.GetCourseResult(courseResultId)
	if err != nil {
		return nil, err
	}
	totalScore := 0.0
	totalCredit := 0
	isPass := false
	var currentCourseId string
	var currentCredit int
	for i, courseSemesterResult := range courseResult.Result {
		if courseSemesterResult.CourseSemesterId == courseSemesterId {
			currentCourseId = courseSemesterResult.CourseId
			courseSemesterResult.Score = score
			if score >= 2 {
				courseSemesterResult.Pass = true
				isPass = true
			} else {
				courseSemesterResult.Pass = false
			}
			courseResult.Result[i] = courseSemesterResult
		}
		courseSemester, err := cc.Repository.GetCourseSemester(courseSemesterResult.CourseSemesterId)
		if err != nil {
			return nil, err
		}

		course, err := cc.Repository.GetCourse(courseSemester.CourseId)
		if err != nil {
			return nil, err
		}

		if courseSemesterResult.CourseSemesterId == courseSemesterId {
			currentCredit = course.Credit
		}
		totalScore += courseSemesterResult.Score * float64(course.Credit)
		totalCredit += course.Credit
	}
	resultScore := totalScore / float64(totalCredit)
	courseResult.Score = resultScore

	transcript, err := cc.Repository.GetTranscript(courseResult.StudentId)
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
			foundResult.CourseResultId = courseResultId
			foundResult.Year = courseResult.Year
			foundResult.Semester = courseResult.Semester
		} else {
			course, err := cc.Repository.GetCourse(transcriptResult.CourseId)
			if err != nil {
				return nil, err
			}

			totalTranscriptCredit += course.Credit
			totalTranscriptScore += (float64(course.Credit) * transcriptResult.Score)
		}
	}

	if found {
		transcript.TranscriptResult[foundIndex] = foundResult
	} else {
		newTranscriptResult := state.TranscriptResult{CourseResultId: courseResultId, CourseId: currentCourseId, Year: courseResult.Year, Semester: courseResult.Semester, Score: score, Pass: isPass}
		transcript.TranscriptResult = append(transcript.TranscriptResult, newTranscriptResult)
	}
	totalTranscriptCredit += currentCredit
	totalTranscriptScore += (float64(currentCredit) * score)

	transcript.Score = totalTranscriptScore / float64(totalTranscriptCredit)

	_, err = cc.Repository.UpdateTranscript(transcript)
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
