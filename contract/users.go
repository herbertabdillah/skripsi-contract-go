package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func GetLecturer(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id = c.ParamString("id")

	return cc.Repository.GetLecturer(id)
}

func CreateLecturer(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id, name, nik = c.ParamString("id"), c.ParamString("name"), c.ParamString("nik")
	lecturer := &state.Lecturer{Name: name, Id: id, Nik: nik}

	return cc.Repository.InsertLecturer(lecturer)
}

func GetStudent(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id = c.ParamString("id")

	return cc.Repository.GetStudent(id)
}

func CreateStudent(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id, name, nim, departmentId, entryYear, status, supervisorLecturerId = c.ParamString("id"), c.ParamString("name"), c.ParamString("nim"), c.ParamString("departmentId"), c.ParamInt("entryYear"), c.ParamString("status"), c.ParamString("supervisorLecturerId")
	student := &state.Student{Name: name, Id: id, Nim: nim, DepartmentId: departmentId, EntryYear: entryYear, Status: status, SupervisorLecturerId: supervisorLecturerId}

	_, err := cc.Repository.GetDepartment(student.DepartmentId)
	if err != nil {
		return nil, err
	}

	_, err = cc.Repository.GetLecturer(student.SupervisorLecturerId)
	if err != nil {
		return nil, err
	}

	transcript := &state.Transcript{Id: id, StudentId: id, Score: 0, TranscriptResult: []state.TranscriptResult{}}
	_, err = cc.Repository.InsertTranscript(transcript)
	if err != nil {
		return nil, err
	}

	err = updateStudentYear(cc, student)
	if err != nil {
		return nil, err
	}

	return cc.Repository.InsertStudent(student)
}

func updateStudentYear(cc Context, student *state.Student) error {
	var updateErr error = nil

	studentYear, err := cc.Repository.GetStudentYear(student.EntryYear)
	if err != nil {
		studentYear = &state.StudentYear{EntryYear: student.EntryYear, StudentIds: []string{student.Id}}
		_, updateErr = cc.Repository.InsertStudentYear(studentYear)

	} else if err == nil {
		studentYear.StudentIds = append(studentYear.StudentIds, student.Id)
		_, updateErr = cc.Repository.UpdateStudentYear(studentYear)
	}

	return updateErr
}
