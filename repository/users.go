package repository

import (
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (r Repository) GetStudent(id string) (*state.Student, error) {
	res, err := r.context.State().Get("Student."+id, &state.Student{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Student)

	return &obj, nil
}

func (r Repository) GetStudentYear(year int) (*state.StudentYear, error) {
	res, err := r.context.State().Get("StudentYear."+strconv.Itoa(year), &state.StudentYear{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.StudentYear)

	return &obj, nil
}

func (r Repository) UpdateStudent(obj *state.Student) (*state.Student, error) {
	err := r.context.State().Put("Student."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
