package service_test

// import (
// 	"testing"

// 	"github.com/herbertabdillah/skripsi-contract-new/contract"
// 	"github.com/herbertabdillah/skripsi-contract-new/state"
// 	t "github.com/herbertabdillah/skripsi-contract-new/test"
// 	"github.com/hyperledger-labs/cckit/router"
// 	"github.com/hyperledger-labs/cckit/serialize"
// 	testcc "github.com/hyperledger-labs/cckit/testing"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// func TestInsertCoursePlan(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Insert course plan service suite")
// }

// var _ = Describe(`Insert Course Plan`, func() {
// 	var ctx contract.Context
// 	var stub *testcc.MockStub
// 	var stubRouterCtx router.Context

// 	BeforeEach(func() {
// 		stub = testcc.NewMockStub(`contract`, contract.NewCC())
// 		stub.Init()
// 		stubRouterCtx = router.NewContext(stub, serialize.DefaultSerializer, router.NewLogger("test"))

// 		ctx = contract.NewContext(stubRouterCtx)

// 		setupInitialData(stub, ctx)
// 		stub.MockTransactionStart("1")
// 	})

// 	AfterEach(func() {
// 		stub.MockTransactionEnd("1")
// 	})

// 	It("run as expected", func() {
// 		coursePlan := &state.CoursePlan{Id: "1", Year: 2017, Semester: "odd", StudentId: "1", Status: "approved", CourseSemesterIds: []string{"1"}}
// 		_, err := ctx.Service.InsertCoursePlan(coursePlan)

// 		Expect(err).To(BeNil())
// 	})
// })

// func setupInitialData(stub *testcc.MockStub, ctx contract.Context) {
// 	t.SetupInitialData(stub)
// 	// stub.MockTransactionStart("1")
// 	// ctx.Repository.InsertApplicationConfig(&state.ApplicationConfig{Year: 2017, Semester: "odd"})
// 	// ctx.Repository.InsertFaculty(&state.Faculty{Id: "1", Name: "Sains dan Teknologi"})
// 	// ctx.Repository.InsertDepartment(&state.Department{Id: "1", Name: "Teknik Informatika", FacultyId: "1"})
// 	// ctx.Repository.InsertCourseYear(&state.CourseYear{Year: 2017, Semester: "odd", Status: "active"})
// 	// stub.MockTransactionEnd("1")
// }
