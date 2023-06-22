package main

import (
	"github.com/herbertabdillah/skripsi-contract-new/contract"
	"github.com/hyperledger-labs/cckit/router"
	"github.com/hyperledger-labs/cckit/router/param"
)

func NewCC() *router.Chaincode {
	r := router.New(`contract`)

	r.Init(func(context router.Context) (i interface{}, e error) {
		return nil, nil
	})

	r.
		// Master data
		Query(`MasterData:getFaculty`, contract.GetFaculty, param.String("id")).
		Invoke(`MasterData:insertFaculty`, contract.CreateFaculty, param.String("id"), param.String("name")).
		Query(`MasterData:getDepartment`, contract.GetDepartment, param.String("id")).
		Invoke(`MasterData:insertDepartment`, contract.CreateDepartment, param.String("id"), param.String("name"), param.String("facultyId")).
		Query(`MasterData:getCourse`, contract.GetCourse, param.String("id")).
		Invoke(`MasterData:insertCourse`, contract.CreateCourse, param.String("id"), param.String("departmentId"), param.String("name"), param.Int("credit"), param.String("kind")).
		// Users
		Query(`User:getLecturer`, contract.GetLecturer, param.String("id")).
		Invoke(`User:insertLecturer`, contract.CreateLecturer, param.String("id"), param.String("name"), param.String("nik")).
		Query(`User:getStudent`, contract.GetStudent, param.String("id")).
		Invoke(`User:insertStudent`, contract.CreateStudent, param.String("id"), param.String("name"), param.String("nim"), param.String("departmentId"), param.Int("entryYear"), param.String("status"), param.String("supervisorLecturerId")).
		// Administration
		Invoke(`Administration:startYear`, contract.StartYear, param.Int("year"), param.String("semester")).
		Invoke(`Administration:endYear`, contract.EndYear, param.Int("year"), param.String("semester")).
		Query(`Administration:getCourseYear`, contract.GetCourseYear, param.Int("year"), param.String("semester")).
		Invoke(`Administration:insertCourseSemester`, contract.InsertCourseSemester, param.String("id"), param.Int("year"), param.String("semester"), param.String("courseId"), param.String("lecturerId")).
		Query(`Administration:getCourseSemester`, contract.GetCourseSemester, param.String("id")).
		// Course
		Invoke(`Course:insertCoursePlan`, contract.InsertCoursePlan, param.String("id"), param.Int("year"), param.String("semester"), param.String("studentId"), param.String("status"), param.String("courseSemesterIds")).
		Invoke(`Course:updateCourseResult`, contract.UpdateCourseResult, param.String("courseSemesterId"), param.String("courseResultId"), param.String("score")).
		Query(`Course:getCoursePlan`, contract.GetCoursePlan, param.String("id")).
		Query(`Course:getCourseResult`, contract.GetCourseResult, param.String("id")).
		// Graduation
		Query(`Graduation:getTranscript`, contract.GetTranscript, param.String("id")).
		Query(`Graduation:graduate`, contract.Graduate, param.String("id"))

	return router.NewChaincode(r)
}
