package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) InitAuditor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New(`wring number of arguments`)
	}
	auditorId := args[0]
	auditorName := args[1]
	auditorAcc := auditor{}
	//////////////////////////////////////////
	//      auditor parsing
	//////////////////////////////////////////
	auditorAcc.Id = auditorId
	auditorAcc.Name = auditorName

	auditorAsbytes, _ := json.Marshal(auditorAcc)
	err := stub.PutState(auditorId, auditorAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}

	return nil, nil
}
