package service

import (
	"errors"
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/config"
	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (s Service) InsertCoursePlan(coursePlan *state.CoursePlan) (*state.CoursePlan, error) {
	if err := s.Updatable(coursePlan.Year, coursePlan.Semester); err != nil {
		return nil, err
	}

	err := s.classQuotaValidation(coursePlan)
	if err != nil {
		return nil, err
	}

	err = s.creditValidation(coursePlan)
	if err != nil {
		return nil, err
	}

	_, err = s.repository.InsertCoursePlan(coursePlan)
	if err != nil {
		return nil, err
	}

	_, err = s.generateCourseResult(coursePlan)
	if err != nil {
		return nil, err
	}

	return coursePlan, err
}

func (s Service) classQuotaValidation(cp *state.CoursePlan) error {
	for _, courseSemesterId := range cp.CourseSemesterIds {
		courseSemester, err := s.repository.GetCourseSemester(courseSemesterId)
		if err != nil {
			return err
		}

		if courseSemester.StudentCount >= config.MAX_STUDENT_PER_CLASS {
			return errors.New("class full")
		}

		courseSemester.StudentCount += 1

		_, err = s.repository.UpdateCourseSemester(courseSemester)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) creditValidation(cp *state.CoursePlan) error {
	totalCredit := 0
	lastSemesterScore := 0.0
	maxCredit := config.MAX_CREDIT_PER_SEMESTER
	for _, courseSemesterId := range cp.CourseSemesterIds {
		courseSemester, err := s.repository.GetCourseSemester(courseSemesterId)
		if err != nil {
			return err
		}

		course, err := s.repository.GetCourse(courseSemester.CourseId)
		if err != nil {
			return err
		}

		totalCredit += course.Credit
	}

	transcript, err := s.repository.GetTranscript(cp.StudentId)
	if err != nil {
		return err
	}

	if transcript == nil || len(transcript.TranscriptResult) == 0 {
		lastSemesterScore = 4.0
	} else {
		lastSemesterScore = transcript.TranscriptResult[len(transcript.TranscriptResult)-1].Score
	}

	if lastSemesterScore < 3 {
		maxCredit = 21
	} else if lastSemesterScore < 2 {
		maxCredit = 18
	}

	if totalCredit >= maxCredit {
		return errors.New("credit exceed " + strconv.Itoa(totalCredit) + " max: " + strconv.Itoa(maxCredit))
	}

	return nil
}

func (s Service) generateCourseResult(cp *state.CoursePlan) (*state.CourseResult, error) {
	result := []state.CourseSemeterResult{}

	for _, courseSemesterId := range cp.CourseSemesterIds {
		courseSemester, err := s.repository.GetCourseSemester(courseSemesterId)
		if err != nil {
			return nil, err
		}
		csr := state.CourseSemeterResult{CourseSemesterId: courseSemesterId, Score: 0, Pass: false, CourseId: courseSemester.CourseId}
		result = append(result, csr)
	}

	courseResult := &state.CourseResult{Id: cp.Id, Year: cp.Year, Semester: cp.Semester, StudentId: cp.StudentId, CoursePlanId: cp.Id, Result: result}
	_, err := s.repository.InsertCourseResult(courseResult)
	if err != nil {
		return nil, err
	}

	return courseResult, nil
}
