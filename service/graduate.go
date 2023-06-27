package service

import (
	"errors"

	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (s Service) Graduate(studentId string) (*state.Student, error) {
	student, err := s.repository.GetStudent(studentId)
	if err != nil {
		return nil, err
	}

	appConfig, err := s.repository.GetApplicationConfig()
	if err != nil {
		return nil, err
	}

	department, err := s.repository.GetDepartment(student.DepartmentId)
	if err != nil {
		return nil, err
	}

	transcript, err := s.repository.GetTranscript(student.Id)
	if err != nil {
		return nil, err
	}

	if student.Status != "active" {
		return nil, errors.New("student not active")
	}

	if !haveDoneCourses(department, transcript) {
		return nil, errors.New("haven't done all course")
	}

	student.Status = "graduated"
	student.ExitYear = appConfig.Year

	return s.repository.UpdateStudent(student)
}

func haveDoneCourses(department *state.Department, transcript *state.Transcript) bool {
	haventDoneCourseIds := department.CourseIds

	for _, transcriptResult := range transcript.TranscriptResult {
		if !transcriptResult.Pass {
			continue
		}

		for courseIndex, courseId := range haventDoneCourseIds {
			if transcriptResult.CourseId == courseId {
				haventDoneCourseIds = remove(haventDoneCourseIds, courseIndex)
			}
		}
	}

	return len(haventDoneCourseIds) <= 0
}

func remove(arr []string, index int) []string {
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = ""
	arr = arr[:len(arr)-1]

	return arr
}
