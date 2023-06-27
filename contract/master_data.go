package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func Init(c router.Context) (interface{}, error) {
	appConfig := &state.ApplicationConfig{}
	return appConfig, c.State().Insert("ApplicationConfig", appConfig)
}

func GetFaculty(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id = c.ParamString("id")

	return cc.Repository.GetFaculty(id)
}

func CreateFaculty(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id, name = c.ParamString("id"), c.ParamString("name")
	var faculty = &state.Faculty{Name: name, Id: id}

	return cc.Repository.InsertFaculty(faculty)
}

func GetDepartment(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id = c.ParamString("id")

	return cc.Repository.GetDepartment(id)
}

func CreateDepartment(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id, name, facultyId = c.ParamString("id"), c.ParamString("name"), c.ParamString("facultyId")
	var department = &state.Department{Name: name, Id: id, FacultyId: facultyId}

	_, err := cc.Repository.GetFaculty(department.FacultyId)
	if err != nil {
		return nil, err
	}

	return cc.Repository.InsertDepartment(department)
}

func GetCourse(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id = c.ParamString("id")

	return cc.Repository.GetCourse(id)
}

func CreateCourse(c router.Context) (interface{}, error) {
	cc := NewContext(c)
	var id, departmentId, name, credit, kind = c.ParamString("id"), c.ParamString("departmentId"), c.ParamString("name"), c.ParamInt("credit"), c.ParamString("kind")
	course := &state.Course{Id: id, DepartmentId: departmentId, Name: name, Credit: credit, Kind: kind}

	err := insertCourseToDepartment(cc, course)
	if err != nil {
		return nil, err
	}

	return cc.Repository.InsertCourse(course)
}

func insertCourseToDepartment(cc Context, course *state.Course) error {
	department, err := cc.Repository.GetDepartment(course.DepartmentId)
	if err != nil {
		return err
	}

	department.CourseIds = append(department.CourseIds, course.Id)

	_, err = cc.Repository.UpdateDepartment(department)

	return err
}
