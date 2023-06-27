package util

import (
	testcc "github.com/hyperledger-labs/cckit/testing"
)

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

func masterData(chaincode *testcc.MockStub) {
	err := chaincode.Invoke("MasterData:init")
	print(err.Message)
	chaincode.Invoke("MasterData:insertFaculty", "1", "Sains dan Teknologi")
	chaincode.Invoke("MasterData:insertDepartment", "1", "Teknik Informatika", "1")
	chaincode.Invoke("MasterData:insertCourse", "1", "1", "Dasar dasar pemrograman", 6, "class")
	chaincode.Invoke("MasterData:insertCourse", "2", "1", "Matematika dasar", 3, "class")
	chaincode.Invoke("MasterData:insertCourse", "3", "1", "Matematika diskrit", 3, "class")
}

func userLecturer(chaincode *testcc.MockStub) {
	chaincode.Invoke("User:insertLecturer", "1", "Donald Knuth", "19450001")
	chaincode.Invoke("User:insertLecturer", "2", "Ken Thompson", "19450002")
	chaincode.Invoke("User:insertLecturer", "3", "Thomas Cormen", "19450003")
}

func userStudent(chaincode *testcc.MockStub) {
	chaincode.Invoke("User:insertStudent", "1", "Herbert", "11170910000046", "1", 2017, "active", "2")
	chaincode.Invoke("User:insertStudent", "3", "Natsir", "11170910000048", "1", 2017, "active", "2")
	chaincode.Invoke("User:insertStudent", "4", "Muso", "11170910000049", "1", 2017, "active", "2")
	chaincode.Invoke("User:insertStudent", "5", "Cokro", "11170910000050", "1", 2017, "active", "2")
	chaincode.Invoke("User:insertStudent", "6", "Aidit", "11170910000051", "1", 2017, "active", "2")
	chaincode.Invoke("User:insertStudent", "7", "Kahar", "11170910000052", "1", 2017, "active", "2")
}

func startYear2017(chaincode *testcc.MockStub) {
	chaincode.Invoke("Administration:startYear", 2017, "even")
	chaincode.Invoke("Administration:endYear", 2017, "even")
	chaincode.Invoke("Administration:startYear", 2017, "odd")
}

func courseSemester(chaincode *testcc.MockStub) {
	chaincode.Invoke("Administration:insertCourseSemester", "1", "2017", "odd", "1", "1")
	chaincode.Invoke("Administration:insertCourseSemester", "2", "2017", "odd", "2", "2")
	chaincode.Invoke("Administration:insertCourseSemester", "3", "2017", "odd", "3", "3")
}
