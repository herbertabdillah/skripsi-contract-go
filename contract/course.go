package contract

import (
	"encoding/json"
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func InsertCoursePlan(c router.Context) (interface{}, error) {
	id, year, semester, studentId, status, courseSemesterIdsRaw := c.ParamString("id"), c.ParamInt("year"), c.ParamString("semester"), c.ParamString("studentId"), c.ParamString("status"), c.ParamString("courseSemesterIds")
	courseSemesterIds := []string{}
	json.Unmarshal([]byte(courseSemesterIdsRaw), &courseSemesterIds)
	coursePlan := &state.CoursePlan{Id: id, Year: year, Semester: semester, StudentId: studentId, Status: status, CourseSemesterIds: courseSemesterIds}

	var err = c.State().Insert("CoursePlan."+id, coursePlan)
	if err != nil {
		return nil, err
	}

	result := []state.CourseSemeterResult{}

	for _, courseSemesterId := range courseSemesterIds {
		courseSemesterRes, err := c.State().Get("CourseSemester."+courseSemesterId, &state.CourseSemester{})
		courseSemester := courseSemesterRes.(state.CourseSemester)
		if err != nil {
			return nil, err
		}
		csr := state.CourseSemeterResult{CourseSemesterId: courseSemesterId, Score: 0, Pass: false, CourseId: courseSemester.CourseId}
		result = append(result, csr)
	}

	courseResult := &state.CourseResult{Id: id, Year: year, Semester: semester, StudentId: studentId, CoursePlanId: id, Result: result}
	err = c.State().Insert("CourseResult."+id, courseResult)

	return coursePlan, err
}

func UpdateCourseResult(c router.Context) (interface{}, error) {
	var err error
	courseSemesterId, courseResultId, scoreStr := c.ParamString("courseSemesterId"), c.ParamString("courseResultId"), c.ParamString("score")
	score, _ := strconv.ParseFloat(scoreStr, 64)
	result, err := c.State().Get("CourseResult."+courseResultId, &state.CourseResult{})
	if err != nil {
		return nil, err
	}
	courseResult := result.(state.CourseResult)
	totalScore := 0.0
	totalCredit := 0
	isPass := false
	var currentCourseId string
	var currentCredit int
	// var course state.Course
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
		courseSemesterRes, err := c.State().Get("CourseSemester."+courseSemesterResult.CourseSemesterId, &state.CourseSemester{})
		if err != nil {
			return nil, err
		}
		courseSemester := courseSemesterRes.(state.CourseSemester)
		courseRes, err := c.State().Get("Course."+courseSemester.CourseId, &state.Course{})
		if err != nil {
			return nil, err
		}
		course := courseRes.(state.Course)

		if courseSemesterResult.CourseSemesterId == courseSemesterId {
			currentCredit = course.Credit
		}
		totalScore += courseSemesterResult.Score * float64(course.Credit)
		totalCredit += course.Credit
	}
	resultScore := totalScore / float64(totalCredit)
	courseResult.Score = resultScore

	transcriptRes, err := c.State().Get("Transcript."+courseResult.StudentId, &state.Transcript{})
	if err != nil {
		return nil, err
	}
	transcript := transcriptRes.(state.Transcript)

	found := false
	totalTranscriptScore := 0.0
	totalTranscriptCredit := 0
	foundIndex := -1
	var foundResult state.TranscriptResult
	for i, transcriptResult := range transcript.TranscriptResult {
		// totalTranscriptScore += float64(transcriptResult.Score)
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
			courseRes, err := c.State().Get("Course."+transcriptResult.CourseId, &state.Course{})
			if err != nil {
				return nil, err
			}
			course := courseRes.(state.Course)

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

	err = c.State().Put("Transcript."+courseResult.StudentId, transcript)
	if err != nil {
		return nil, err
	}

	return courseResult, c.State().Put("CourseResult."+courseResultId, courseResult)
}

func GetCoursePlan(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("CoursePlan."+id, &state.CoursePlan{})
}

func GetCourseResult(c router.Context) (interface{}, error) {
	id := c.ParamString("id")

	return c.State().Get("CourseResult."+id, &state.CourseResult{})
}
