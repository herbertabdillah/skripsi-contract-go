package repository

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
)

func (r Repository) GetCourseResult(id string) (*state.CourseResult, error) {
	res, err := r.context.State().Get("CourseResult."+id, &state.CourseResult{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.CourseResult)

	return &obj, nil
}

func (r Repository) GetCoursePlan(id string) (*state.CoursePlan, error) {
	res, err := r.context.State().Get("CoursePlan."+id, &state.CoursePlan{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.CoursePlan)

	return &obj, nil
}

func (r Repository) InsertCoursePlan(obj *state.CoursePlan) (*state.CoursePlan, error) {
	err := r.context.State().Insert("CoursePlan."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) InsertCourseResult(obj *state.CourseResult) (*state.CourseResult, error) {
	err := r.context.State().Insert("CourseResult."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateCourseResult(obj *state.CourseResult) (*state.CourseResult, error) {
	err := r.context.State().Put("CourseResult."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateCourseSemester(obj *state.CourseSemester) (*state.CourseSemester, error) {
	err := r.context.State().Put("CourseSemester."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
