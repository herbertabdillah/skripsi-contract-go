package service

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (s Service) UpdateCourseResult(courseResult *state.CourseResult, courseSemesterId string, score float64) (*state.CourseResult, error) {
	if err := s.Updatable(courseResult.Year, courseResult.Semester); err != nil {
		return nil, err
	}

	courseResult, err := s.repository.GetCourseResult(courseResult.Id)
	if err != nil {
		return nil, err
	}

	isPass := false
	if score >= 2 {
		isPass = true
	}

	_, course, err := s.updateCourseResult(courseResult, courseSemesterId, score, isPass)
	if err != nil {
		return nil, err
	}

	_, err = s.updateTranscript(courseResult, course.Id, score, isPass, course.Credit)
	if err != nil {
		return nil, err
	}

	return s.repository.UpdateCourseResult(courseResult)
}

func (s Service) updateCourseResult(cr *state.CourseResult, courseSemesterId string, score float64, isPass bool) (*state.CourseResult, *state.Course, error) {
	var currentCourse *state.Course
	totalScore := 0.0
	totalCredit := 0

	for i, courseSemesterResult := range cr.Result {
		courseSemester, err := s.repository.GetCourseSemester(courseSemesterResult.CourseSemesterId)
		if err != nil {
			return nil, nil, err
		}

		course, err := s.repository.GetCourse(courseSemester.CourseId)
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

	_, err := s.repository.UpdateCourseResult(cr)

	if err != nil {
		return nil, currentCourse, err
	}

	return cr, currentCourse, nil
}

func (s Service) updateTranscript(cr *state.CourseResult, currentCourseId string, score float64, isPass bool, credit int) (*state.Transcript, error) {
	transcript, err := s.repository.GetTranscript(cr.StudentId)
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

		course, err := s.repository.GetCourse(transcriptResult.CourseId)
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

	_, err = s.repository.UpdateTranscript(transcript)
	if err != nil {
		return nil, err
	}

	return transcript, nil
}
