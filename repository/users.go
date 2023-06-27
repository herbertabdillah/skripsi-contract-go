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
	key := "StudentYear." + strconv.Itoa(year)
	res, err := r.context.State().Get(key, &state.StudentYear{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.StudentYear)

	return &obj, nil
}

func (r Repository) UpdateStudentYear(obj *state.StudentYear) (*state.StudentYear, error) {
	key := "StudentYear." + strconv.Itoa(obj.EntryYear)
	err := r.context.State().Put(key, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) InsertStudentYear(obj *state.StudentYear) (*state.StudentYear, error) {
	key := "StudentYear." + strconv.Itoa(obj.EntryYear)
	err := r.context.State().Insert(key, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateStudent(obj *state.Student) (*state.Student, error) {
	err := r.context.State().Put("Student."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) InsertStudent(obj *state.Student) (*state.Student, error) {
	err := r.context.State().Insert("Student."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) GetLecturer(id string) (*state.Lecturer, error) {
	res, err := r.context.State().Get("Lecturer."+id, &state.Lecturer{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Lecturer)

	return &obj, nil
}

func (r Repository) InsertLecturer(obj *state.Lecturer) (*state.Lecturer, error) {
	err := r.context.State().Insert("Lecturer."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
