package contract_test

import (
	"testing"

	"github.com/herbertabdillah/skripsi-contract-new/config"
	"github.com/herbertabdillah/skripsi-contract-new/contract"

	testcc "github.com/hyperledger-labs/cckit/testing"
	expectcc "github.com/hyperledger-labs/cckit/testing/expect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInsertCoursePlan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main suite")
}

var _ = Describe(`Insert course plan`, func() {
	var chaincode *testcc.MockStub

	BeforeEach(func() {
		chaincode = testcc.NewMockStub(`contract`, contract.NewCC())
		SetupInitialData(chaincode)
	})

	Describe("quota exceed", func() {
		It("error", func() {
			config.MAX_STUDENT_PER_CLASS = 1

			expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "1", "2017", "odd", "1", "approved", `["1","2","3"]`))
			expectcc.ResponseError(chaincode.Invoke("Course:insertCoursePlan", "1002", "2017", "odd", "5", "approved", `["1"]`))

			config.MAX_STUDENT_PER_CLASS = 40
		})
	})

	// Describe("last semester credit", func() {
	// 	It("error", func() {

	// 		expectcc.ResponseOk(chaincode.Invoke("Course:insertCoursePlan", "1", "2017", "odd", "1", "approved", `["1","2","3"]`))
	// 		expectcc.ResponseError(chaincode.Invoke("Course:insertCoursePlan", "1002", "2017", "odd", "5", "approved", `["1"]`))

	// 	})
	// })
})
