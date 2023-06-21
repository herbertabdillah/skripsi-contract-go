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
		Query(`MasterData:getFaculty`, contract.GetFaculty, param.String("id")).
		Invoke(`MasterData:insertFaculty`, contract.CreateFaculty, param.String("id"), param.String("name"))

	return router.NewChaincode(r)
}
