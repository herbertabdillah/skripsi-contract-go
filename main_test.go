package main_test

import (
	"testing"

	main "github.com/herbertabdillah/skripsi-contract-new"
	"github.com/herbertabdillah/skripsi-contract-new/state"
	testcc "github.com/hyperledger-labs/cckit/testing"
	expectcc "github.com/hyperledger-labs/cckit/testing/expect"

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
			expectcc.ResponseOk(chaincode.Invoke("MasterData:insertFaculty", "1", "Sains dan Teknologi"))

			queryResponse := chaincode.Query("MasterData:getFaculty", "1")
			faculty := expectcc.PayloadIs(queryResponse, &state.Faculty{}).(state.Faculty)
			Expect(faculty.Name).To(Equal("Sains dan Teknologi"))
		})
	})
})
