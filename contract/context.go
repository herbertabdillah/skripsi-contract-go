package contract

import (
	r "github.com/herbertabdillah/skripsi-contract-new/repository"
	s "github.com/herbertabdillah/skripsi-contract-new/service"
	"github.com/hyperledger-labs/cckit/router"
	"github.com/hyperledger-labs/cckit/router/param"
)

type Context struct {
	Service    s.Service
	Repository r.Repository
}

func NewContext(rc router.Context) Context {
	c := Context{}
	c.Repository = r.NewRepository(rc)
	c.Service = s.NewService(rc, c.Repository)

	return c
}

func NewCC() *router.Chaincode {
	r := router.New(`contract`)

	r.Init(func(context router.Context) (i interface{}, e error) {
		return nil, nil
	})

	r.
		// Master data
		Invoke(`MasterData:init`, Init).
		Query(`MasterData:getFaculty`, GetFaculty, param.String("id")).
		Invoke(`MasterData:insertFaculty`, CreateFaculty, param.String("id"), param.String("name")).
		Query(`MasterData:getDepartment`, GetDepartment, param.String("id")).
		Invoke(`MasterData:insertDepartment`, CreateDepartment, param.String("id"), param.String("name"), param.String("facultyId")).
		Query(`MasterData:getCourse`, GetCourse, param.String("id")).
		Invoke(`MasterData:insertCourse`, CreateCourse, param.String("id"), param.String("departmentId"), param.String("name"), param.Int("credit"), param.String("kind")).
		// Users
		Query(`User:getLecturer`, GetLecturer, param.String("id")).
		Invoke(`User:insertLecturer`, CreateLecturer, param.String("id"), param.String("name"), param.String("nik")).
		Query(`User:getStudent`, GetStudent, param.String("id")).
		Invoke(`User:insertStudent`, CreateStudent, param.String("id"), param.String("name"), param.String("nim"), param.String("departmentId"), param.Int("entryYear"), param.String("status"), param.String("supervisorLecturerId")).
		// Administration
		Invoke(`Administration:startYear`, StartYear, param.Int("year"), param.String("semester")).
		Invoke(`Administration:endYear`, EndYear, param.Int("year"), param.String("semester")).
		Query(`Administration:getCourseYear`, GetCourseYear, param.Int("year"), param.String("semester")).
		Invoke(`Administration:insertCourseSemester`, InsertCourseSemester, param.String("id"), param.Int("year"), param.String("semester"), param.String("courseId"), param.String("lecturerId")).
		Query(`Administration:getCourseSemester`, GetCourseSemester, param.String("id")).
		// Course
		Invoke(`Course:insertCoursePlan`, InsertCoursePlan, param.String("id"), param.Int("year"), param.String("semester"), param.String("studentId"), param.String("status"), param.String("courseSemesterIds")).
		Invoke(`Course:updateCourseResult`, UpdateCourseResult, param.String("courseSemesterId"), param.String("courseResultId"), param.String("score")).
		Query(`Course:getCoursePlan`, GetCoursePlan, param.String("id")).
		Query(`Course:getCourseResult`, GetCourseResult, param.String("id")).
		// Graduation
		Query(`Graduation:getTranscript`, GetTranscript, param.String("id")).
		Query(`Graduation:graduate`, Graduate, param.String("id"))

	return router.NewChaincode(r)
}
