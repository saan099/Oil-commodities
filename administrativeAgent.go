package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) InitAdministrativeAgent(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New(`wrong number of arguments`)
	}
	adminId := args[0]
	name := args[1]
	email := args[2]

	adminAcc := administrativeAgent{}
	/////////////////////////////////////////////
	//        administrator parsing
	/////////////////////////////////////////////
	adminAcc.Id = adminId
	adminAcc.Name = name
	adminAcc.Email = email

	borrowerAsbytes, _ := json.Marshal(adminAcc)
	err := stub.PutState(adminId, borrowerAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}
	return borrowerAsbytes, nil
}
