package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) PutAmountforCase(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, errors.New(`wrong number of arguments`)
	}

	return nil, nil
}
