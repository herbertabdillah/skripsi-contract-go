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
