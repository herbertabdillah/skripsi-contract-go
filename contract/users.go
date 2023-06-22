package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func GetLecturer(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get("Lecturer."+id, &state.Lecturer{})
}

func CreateLecturer(c router.Context) (interface{}, error) {
	var id, name, nik = c.ParamString("id"), c.ParamString("name"), c.ParamString("nik")
	lecturer := &state.Lecturer{Name: name, Id: id, Nik: nik}

	return lecturer, c.State().Insert("Lecturer."+id, lecturer)
}

func GetStudent(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get("Student."+id, &state.Student{})
}

func CreateStudent(c router.Context) (interface{}, error) {
	var id, name, nim, departmentId, entryYear, status, supervisorLecturerId = c.ParamString("id"), c.ParamString("name"), c.ParamString("nim"), c.ParamString("departmentId"), c.ParamInt("entryYear"), c.ParamString("status"), c.ParamString("supervisorLecturerId")
	student := &state.Student{Name: name, Id: id, Nim: nim, DepartmentId: departmentId, EntryYear: entryYear, Status: status, SupervisorLecturerId: supervisorLecturerId}

	var err error
	_, err = c.State().Get("Department."+departmentId, &state.Department{})
	if err != nil {
		return nil, err
	}
	_, err = c.State().Get("Lecturer."+supervisorLecturerId, &state.Lecturer{})
	if err != nil {
		return nil, err
	}

	transcript := &state.Transcript{Id: id, StudentId: id, Score: 0, TranscriptResult: []state.TranscriptResult{}}
	err = c.State().Insert("Transcript."+id, transcript)
	if err != nil {
		return nil, err
	}

	return student, c.State().Insert("Student."+id, student)
}
