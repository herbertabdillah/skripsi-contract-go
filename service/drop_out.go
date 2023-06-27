package service

import (
	"github.com/herbertabdillah/skripsi-contract-new/config"
)

func (s Service) DroupOut(year int) error {
	studentYear, _ := s.repository.GetStudentYear(year - config.MAX_STUDY_YEAR)

	if studentYear == nil {
		return nil
	}

	for _, studentId := range studentYear.StudentIds {
		student, err := s.repository.GetStudent(studentId)
		if err != nil {
			return err
		}

		if student.Status != "active" {
			continue
		}

		student.Status = "drop_out"
		student.ExitYear = year

		_, err = s.repository.UpdateStudent(student)
		if err != nil {
			return err
		}
	}

	return nil
}
