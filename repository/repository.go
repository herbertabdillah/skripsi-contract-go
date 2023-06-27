package repository

import (
	"strconv"

	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

type Repository struct {
	context router.Context
}

func NewRepository(c router.Context) Repository {
	return Repository{context: c}
}

func (r Repository) GetApplicationConfig() (*state.ApplicationConfig, error) {
	res, err := r.context.State().Get("ApplicationConfig", &state.ApplicationConfig{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.ApplicationConfig)

	return &obj, nil
}

func (r Repository) GetCourseYear(year int, semester string) (*state.CourseYear, error) {
	var semesterNumber int
	if semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	key := "CourseYear." + strconv.Itoa(year) + strconv.Itoa(semesterNumber)
	res, err := r.context.State().Get(key, &state.CourseYear{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.CourseYear)

	return &obj, nil
}
func (r Repository) GetCourseSemester(id string) (*state.CourseSemester, error) {
	res, err := r.context.State().Get("CourseSemester."+id, &state.CourseSemester{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.CourseSemester)

	return &obj, nil
}

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

func (r Repository) GetCourse(id string) (*state.Course, error) {
	res, err := r.context.State().Get("Course."+id, &state.Course{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Course)

	return &obj, nil
}

func (r Repository) GetTranscript(id string) (*state.Transcript, error) {
	res, err := r.context.State().Get("Transcript."+id, &state.Transcript{})
	if err != nil {
		return nil, err
	}
	obj := res.(state.Transcript)

	return &obj, nil
}

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

func (r Repository) UpdateStudent(obj *state.Student) (*state.Student, error) {
	err := r.context.State().Put("Student."+obj.Id, obj)
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

func (r Repository) UpdateTranscript(obj *state.Transcript) (*state.Transcript, error) {
	err := r.context.State().Put("Transcript."+obj.Id, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateApplicationConfig(obj *state.ApplicationConfig) (*state.ApplicationConfig, error) {
	err := r.context.State().Put("ApplicationConfig", obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) UpdateCourseYear(obj *state.CourseYear) (*state.CourseYear, error) {
	var semesterNumber int
	if obj.Semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	key := "CourseYear." + strconv.Itoa(obj.Year) + strconv.Itoa(semesterNumber)
	err := r.context.State().Put(key, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r Repository) InsertCourseYear(obj *state.CourseYear) (*state.CourseYear, error) {
	var semesterNumber int
	if obj.Semester == "even" {
		semesterNumber = 1
	} else {
		semesterNumber = 2
	}

	key := "CourseYear." + strconv.Itoa(obj.Year) + strconv.Itoa(semesterNumber)
	err := r.context.State().Insert(key, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
