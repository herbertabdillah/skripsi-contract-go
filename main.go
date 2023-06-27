package main

import (
	"fmt"

	"github.com/herbertabdillah/skripsi-contract-new/contract"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	if err := shim.Start(contract.NewCC()); err != nil {
		fmt.Printf("Error starting hello world chaincode: %s", err)
	}
}
