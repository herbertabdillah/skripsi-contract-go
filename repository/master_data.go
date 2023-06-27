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

func (r Repository) InsertCourse(obj *state.Course) (*state.Course, error) {
	err := r.context.State().Insert("Course."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) GetDepartment(id string) (*state.Department, error) {
	res, err := r.context.State().Get("Department."+id, &state.Department{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Department)

	return &obj, nil
}

func (r Repository) InsertDepartment(obj *state.Department) (*state.Department, error) {
	err := r.context.State().Insert("Department."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) GetFaculty(id string) (*state.Faculty, error) {
	res, err := r.context.State().Get("Faculty."+id, &state.Faculty{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Faculty)

	return &obj, nil
}

func (r Repository) InsertFaculty(obj *state.Faculty) (*state.Faculty, error) {
	err := r.context.State().Insert("Faculty."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateDepartment(obj *state.Department) (*state.Department, error) {
	err := r.context.State().Put("Department."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) GetApplicationConfig() (*state.ApplicationConfig, error) {
	res, err := r.context.State().Get("ApplicationConfig", &state.ApplicationConfig{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.ApplicationConfig)

	return &obj, nil
}

func (r Repository) InsertApplicationConfig(appConfig *state.ApplicationConfig) (*state.ApplicationConfig, error) {
	err := r.context.State().Insert("ApplicationConfig", appConfig)
	if err != nil {
		return nil, err
	}

	return appConfig, nil
}
