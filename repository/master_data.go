package repository

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (r Repository) GetCourse(id string) (*state.Course, error) {
	res, err := r.context.State().Get("Course."+id, &state.Course{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Course)

	return &obj, nil
}

func (r Repository) GetDepartment(id string) (*state.Department, error) {
	res, err := r.context.State().Get("Department."+id, &state.Department{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Department)

	return &obj, nil
}
