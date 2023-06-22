package contract

import (
	"github.com/herbertabdillah/skripsi-contract-new/state"
	"github.com/hyperledger-labs/cckit/router"
)

func GetFaculty(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get("Faculty."+id, &state.Faculty{})
}

func CreateFaculty(c router.Context) (interface{}, error) {
	var id, name = c.ParamString("id"), c.ParamString("name")
	var faculty = &state.Faculty{Name: name, Id: id}

	return faculty, c.State().Insert("Faculty."+id, faculty)
}

func GetDepartment(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get("Department."+id, &state.Department{})
}

func CreateDepartment(c router.Context) (interface{}, error) {
	var id, name, facultyId = c.ParamString("id"), c.ParamString("name"), c.ParamString("facultyId")
	var department = &state.Department{Name: name, Id: id, FacultyId: facultyId}

	_, err := c.State().Get("Faculty."+facultyId, &state.Faculty{})
	if err != nil {
		return nil, err
	}

	return department, c.State().Insert("Department."+id, department)
}

func GetCourse(c router.Context) (interface{}, error) {
	var id = c.ParamString("id")

	return c.State().Get("Course."+id, &state.Course{})
}

func CreateCourse(c router.Context) (interface{}, error) {
	var id, departmentId, name, credit, kind = c.ParamString("id"), c.ParamString("departmentId"), c.ParamString("name"), c.ParamInt("credit"), c.ParamString("kind")
	course := &state.Course{Id: id, DepartmentId: departmentId, Name: name, Credit: credit, Kind: kind}

	_, err := c.State().Get("Department."+departmentId, &state.Department{})
	if err != nil {
		return nil, err
	}

	return course, c.State().Insert("Course."+id, course)
}
