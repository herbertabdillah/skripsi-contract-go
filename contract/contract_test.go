package contract_test

import (
	"testing"

	"github.com/herbertabdillah/skripsi-contract-new/config"
	"github.com/herbertabdillah/skripsi-contract-new/contract"
	"github.com/herbertabdillah/skripsi-contract-new/state"
	testcc "github.com/hyperledger-labs/cckit/testing"
	expectcc "github.com/hyperledger-labs/cckit/testing/expect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLifecycleEndToEnd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main suite")
}

var _ = Describe(`Lifecycle End to End (from master data to graduation)`, func() {
	var chaincode *testcc.MockStub

	It("run as expected", func() {
		chaincode = testcc.NewMockStub(`contract`, contract.NewCC())

		config.MAX_STUDENT_PER_CLASS = 3
		expectcc.ResponseOk(chaincode.Init())

		masterData(chaincode)
		userLecturer(chaincode)

		expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2010, "odd"))
		expectcc.ResponseOk(chaincode.Invoke("Administration:endYear", 2010, "odd"))
		expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "2", "Soekarno", "11100910000047", "1", 2010, "active", "2"))

		userStudent(chaincode)
		startYear2017(chaincode)
		courseSemester(chaincode)
		coursePlan(chaincode)
		updateCourseResult(chaincode)

		// haven't done all course
		expectcc.ResponseError(chaincode.Invoke("Graduation:graduate", "1"))

		endYear2017AndCheckDroupout(chaincode)

		expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2018, "even"))

		student1RetakeCourse(chaincode)
		graduateStudent1(chaincode)

		config.MAX_STUDENT_PER_CLASS = 40
	})
})

func masterData(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("MasterData:init"))
	expectcc.ResponseOk(chaincode.Invoke("MasterData:insertFaculty", "1", "Sains dan Teknologi"))

	queryResponse := chaincode.Query("MasterData:getFaculty", "1")
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

	queryResponse = chaincode.Query("MasterData:getDepartment", "1")
	department = expectcc.PayloadIs(queryResponse, &state.Department{}).(state.Department)
	Expect(department.CourseIds).To(Equal([]string{"1", "2", "3"}))
}

func userLecturer(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "1", "Donald Knuth", "19450001"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "2", "Ken Thompson", "19450002"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertLecturer", "3", "Thomas Cormen", "19450003"))

	queryResponse := chaincode.Query("User:getLecturer", "2")
	lecturer := expectcc.PayloadIs(queryResponse, &state.Lecturer{}).(state.Lecturer)
	Expect(lecturer.Name).To(Equal("Ken Thompson"))
	Expect(lecturer.Nik).To(Equal("19450002"))
}

func userStudent(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "1", "Herbert", "11170910000046", "1", 2017, "active", "2"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "3", "Natsir", "11170910000048", "1", 2017, "active", "2"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "4", "Muso", "11170910000049", "1", 2017, "active", "2"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "5", "Cokro", "11170910000050", "1", 2017, "active", "2"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "6", "Aidit", "11170910000051", "1", 2017, "active", "2"))
	expectcc.ResponseOk(chaincode.Invoke("User:insertStudent", "7", "Kahar", "11170910000052", "1", 2017, "active", "2"))

	queryResponse := chaincode.Query("User:getStudent", "2")
	student := expectcc.PayloadIs(queryResponse, &state.Student{}).(state.Student)
	Expect(student.Name).To(Equal("Soekarno"))
	Expect(student.Nim).To(Equal("11100910000047"))
	Expect(student.EntryYear).To(Equal(2010))
	Expect(student.Status).To(Equal("active"))

	queryResponse = chaincode.Query("Graduation:getTranscript", "2")
	transcript := expectcc.PayloadIs(queryResponse, &state.Transcript{}).(state.Transcript)
	Expect(transcript.Id).To(Equal("2"))
	Expect(transcript.Score).To(Equal(0.0))
	Expect(transcript.StudentId).To(Equal("2"))
	Expect(transcript.TranscriptResult).To(Equal([]state.TranscriptResult{}))
}

func startYear2017(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2017, "even"))
	expectcc.ResponseOk(chaincode.Invoke("Administration:endYear", 2017, "even"))
	expectcc.ResponseOk(chaincode.Invoke("Administration:startYear", 2017, "odd"))

	queryResponse := chaincode.Query("Administration:getCourseYear", 2017, "even")
	courseYear := expectcc.PayloadIs(queryResponse, &state.CourseYear{}).(state.CourseYear)
	Expect(courseYear.Year).To(Equal(2017))
	Expect(courseYear.Semester).To(Equal("even"))
	Expect(courseYear.Status).To(Equal("end"))
}

func courseSemester(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "1", "2017", "odd", "1", "1"))
	expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "2", "2017", "odd", "2", "2"))
	expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "3", "2017", "odd", "3", "3"))

	queryResponse := chaincode.Query("Administration:getCourseSemester", "2")
	courseSemester := expectcc.PayloadIs(queryResponse, &state.CourseSemester{}).(state.CourseSemester)
	Expect(courseSemester.Id).To(Equal("2"))
	Expect(courseSemester.Year).To(Equal(2017))
	Expect(courseSemester.Semester).To(Equal("odd"))
	Expect(courseSemester.LecturerId).To(Equal("2"))
	Expect(courseSemester.CourseId).To(Equal("2"))
}

func coursePlan(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "1", "2017", "odd", "1", "approved", `["1","2","3"]`))

	expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "1002", "2017", "odd", "5", "approved", `["1"]`))
	expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "1003", "2017", "odd", "6", "approved", `["1"]`))
	// class full
	expectcc.ResponseError(chaincode.Invoke("Course:insertCoursePlan", "1004", "2017", "odd", "7", "approved", `["1"]`))

	queryResponse := chaincode.Query("Course:getCoursePlan", "1")
	coursePlan := expectcc.PayloadIs(queryResponse, &state.CoursePlan{}).(state.CoursePlan)
	Expect(coursePlan.Id).To(Equal("1"))
	Expect(coursePlan.Status).To(Equal("approved"))
	Expect(coursePlan.CourseSemesterIds).To(Equal([]string{"1", "2", "3"}))

	queryResponse = chaincode.Query("Course:getCourseResult", "1")
	courseResult := expectcc.PayloadIs(queryResponse, &state.CourseResult{}).(state.CourseResult)
	Expect(courseResult.Id).To(Equal("1"))
	Expect(courseResult.Year).To(Equal(2017))
	Expect(courseResult.Semester).To(Equal("odd"))
	Expect(courseResult.Result).To(Equal([]state.CourseSemeterResult{
		{CourseSemesterId: "1", Score: 0, Pass: false, CourseId: "1"},
		{CourseSemesterId: "2", Score: 0, Pass: false, CourseId: "2"},
		{CourseSemesterId: "3", Score: 0, Pass: false, CourseId: "3"},
	}))
}

func updateCourseResult(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("Course:updateCourseResult", "1", "1", "1"))
	expectcc.ResponseOk(chaincode.Invoke("Course:updateCourseResult", "2", "1", "3.4"))
	expectcc.ResponseOk(chaincode.Invoke("Course:updateCourseResult", "3", "1", "4"))

	queryResponse := chaincode.Query("Course:getCourseResult", "1")
	courseResult := expectcc.PayloadIs(queryResponse, &state.CourseResult{}).(state.CourseResult)
	Expect(courseResult.Id).To(Equal("1"))
	Expect(courseResult.Year).To(Equal(2017))
	Expect(courseResult.Semester).To(Equal("odd"))
	Expect(courseResult.Score).To(Equal(2.35))
	Expect(courseResult.Result).To(Equal([]state.CourseSemeterResult{
		{CourseSemesterId: "1", Score: 1, Pass: false, CourseId: "1"},
		{CourseSemesterId: "2", Score: 3.4, Pass: true, CourseId: "2"},
		{CourseSemesterId: "3", Score: 4, Pass: true, CourseId: "3"},
	}))

	queryResponse = chaincode.Query("Graduation:getTranscript", "1")
	transcript := expectcc.PayloadIs(queryResponse, &state.Transcript{}).(state.Transcript)
	Expect(transcript.Id).To(Equal("1"))
	Expect(transcript.Score).To(Equal(2.35))
	Expect(transcript.StudentId).To(Equal("1"))
	Expect(transcript.TranscriptResult).To(Equal([]state.TranscriptResult{
		{CourseResultId: "1", CourseId: "1", Year: 2017, Semester: "odd", Score: 1, Pass: false},
		{CourseResultId: "1", CourseId: "2", Year: 2017, Semester: "odd", Score: 3.4, Pass: true},
		{CourseResultId: "1", CourseId: "3", Year: 2017, Semester: "odd", Score: 4, Pass: true},
	}))
}

func endYear2017AndCheckDroupout(chaincode *testcc.MockStub) {
	queryResponse := chaincode.Query("User:getStudent", "2")
	student := expectcc.PayloadIs(queryResponse, &state.Student{}).(state.Student)
	Expect(student.Status).To(Equal("active"))

	expectcc.ResponseOk(chaincode.Invoke("Administration:endYear", 2017, "odd"))

	queryResponse = chaincode.Query("User:getStudent", "2")
	student = expectcc.PayloadIs(queryResponse, &state.Student{}).(state.Student)
	Expect(student.Status).To(Equal("drop_out"))
	Expect(student.ExitYear).To(Equal(2017))
}

func student1RetakeCourse(chaincode *testcc.MockStub) {
	expectcc.ResponseOk(chaincode.Invoke("Administration:insertCourseSemester", "4", "2018", "even", "1", "1"))
	expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "2", "2018", "even", "1", "approved", `["4"]`))
	expectcc.ResponseOk(chaincode.Invoke("Course:updateCourseResult", "4", "2", "4"))
}

func graduateStudent1(chaincode *testcc.MockStub) {
	queryResponse := chaincode.Query("Graduation:getTranscript", "1")
	transcript := expectcc.PayloadIs(queryResponse, &state.Transcript{}).(state.Transcript)
	Expect(transcript.Score).To(Equal(3.85))
	Expect(transcript.TranscriptResult).To(Equal([]state.TranscriptResult{
		{CourseResultId: "2", CourseId: "1", Year: 2018, Semester: "even", Score: 4, Pass: true},
		{CourseResultId: "1", CourseId: "2", Year: 2017, Semester: "odd", Score: 3.4, Pass: true},
		{CourseResultId: "1", CourseId: "3", Year: 2017, Semester: "odd", Score: 4, Pass: true},
	}))

	expectcc.ResponseOk(chaincode.Invoke("Graduation:graduate", "1"))

	queryResponse = chaincode.Query("User:getStudent", "1")
	student := expectcc.PayloadIs(queryResponse, &state.Student{}).(state.Student)
	Expect(student.Status).To(Equal("graduated"))
	Expect(student.ExitYear).To(Equal(2018))
}

func SetupInitialData(chaincode *testcc.MockStub) {
	chaincode.Init()

	masterData(chaincode)
	userLecturer(chaincode)

	chaincode.Invoke("Administration:startYear", 2010, "odd")
	chaincode.Invoke("Administration:endYear", 2010, "odd")
	chaincode.Invoke("User:insertStudent", "2", "Soekarno", "11100910000047", "1", 2010, "active", "2")

	userStudent(chaincode)
	startYear2017(chaincode)
	courseSemester(chaincode)
}
