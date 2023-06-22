package main_test

import (
	"testing"

	main "github.com/herbertabdillah/skripsi-contract-new"
	"github.com/herbertabdillah/skripsi-contract-new/state"
	testcc "github.com/hyperledger-labs/cckit/testing"
	expectcc "github.com/hyperledger-labs/cckit/testing/expect"

	"github.com/hyperledger/fabric-protos-go/peer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContract(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commercial Paper Suite")
}

var _ = Describe(`contract`, func() {
	chaincode := testcc.NewMockStub(`contract`, main.NewCC())

	BeforeSuite(func() {
		expectcc.ResponseOk(chaincode.Init())
	})

	Describe("Lifecycle", func() {
		It("true", func() {
			var queryResponse peer.Response
			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertFaculty", "1", "Sains dan Teknologi"))
			queryResponse = chaincode.Query("MasterData:getFaculty", "1")
			faculty := expectcc.PayloadIs(queryResponse, &state.Faculty{}).(state.Faculty)
			Expect(faculty.Name).To(Equal("Sains dan Teknologi"))

			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertDepartment", "1", "Teknik Informatika", "1"))
			queryResponse = chaincode.Query("MasterData:getDepartment", "1")
			department := expectcc.PayloadIs(queryResponse, &state.Department{}).(state.Department)
			Expect(department.Name).To(Equal("Teknik Informatika"))
			Expect(department.FacultyId).To(Equal("1"))

			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertCourse", "1", "1", "Dasar dasar pemrograman", 6, "class"))
			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertCourse", "2", "1", "Matematika dasar", 3, "class"))
			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertCourse", "3", "1", "Matematika diskrit", 3, "class"))
			queryResponse = chaincode.Query("MasterData:getCourse", "2")
			course := expectcc.PayloadIs(queryResponse, &state.Course{}).(state.Course)
			Expect(course.Name).To(Equal("Matematika dasar"))
			Expect(course.Credit).To(Equal(3))
			Expect(course.DepartmentId).To(Equal("1"))
			Expect(course.Kind).To(Equal("class"))

			expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "1", "Donald Knuth", "19450001"))
			expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "2", "Ken Thompson", "19450002"))
			expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "3", "Thomas Cormen", "19450003"))
			queryResponse = chaincode.Query("User:getLecturer", "2")
			lecturer := expectcc.PayloadIs(queryResponse, &state.Lecturer{}).(state.Lecturer)
			Expect(lecturer.Name).To(Equal("Ken Thompson"))
			Expect(lecturer.Nik).To(Equal("19450002"))

			expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "1", "Herbert", "11170910000046", "1", 2017, "active", "2"))
			expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "2", "Soekarno", "11150910000047", "1", 2015, "active", "2"))
			expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "3", "Natsir", "11170910000048", "1", 2017, "active", "2"))
			queryResponse = chaincode.Query("User:getStudent", "2")
			student := expectcc.PayloadIs(queryResponse, &state.Student{}).(state.Student)
			Expect(student.Name).To(Equal("Soekarno"))
			Expect(student.Nim).To(Equal("11150910000047"))
			Expect(student.EntryYear).To(Equal(2015))
			Expect(student.Status).To(Equal("active"))

			expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2017, "even"))
			expectcc.ResponseOk(chaincode.Invoke("Administration:endYear", 2017, "even"))
			expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2017, "odd"))
			queryResponse = chaincode.Query("Administration:getCourseYear", 2017, "even")
			courseYear := expectcc.PayloadIs(queryResponse, &state.CourseYear{}).(state.CourseYear)
			Expect(courseYear.Year).To(Equal(2017))
			Expect(courseYear.Semester).To(Equal("even"))
			Expect(courseYear.Status).To(Equal("end"))

			expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "1", "2017", "odd", "1", "1"))
			expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "2", "2017", "odd", "2", "2"))
			expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "3", "2017", "odd", "3", "3"))
			queryResponse = chaincode.Query("Administration:getCourseSemester", "2")
			courseSemester := expectcc.PayloadIs(queryResponse, &state.CourseSemester{}).(state.CourseSemester)
			Expect(courseSemester.Id).To(Equal("2"))
			Expect(courseSemester.Year).To(Equal(2017))
			Expect(courseSemester.Semester).To(Equal("odd"))
			Expect(courseSemester.LecturerId).To(Equal("2"))
			Expect(courseSemester.CourseId).To(Equal("2"))
		})
	})
})
