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

func (t *Oilchain) UpdateLoanPackage(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New(`worng number of arguments`)
	}
	administrativeAgentId := args[0]
	loanPackageId := args[1]
	status := args[2]

	adminAgentAcc := administrativeAgent{}
	adminAgentAsbytes, _ := stub.GetState(administrativeAgentId)
	_ = json.Unmarshal(adminAgentAsbytes, &adminAgentAcc)
	loanPack := loanPackage{}
	var borrowerId string
	for i := range adminAgentAcc.LoanPackage {
		if adminAgentAcc.LoanPackage[i].Id == loanPackageId {
			adminAgentAcc.LoanPackage[i].Status = status
			loanPack = adminAgentAcc.LoanPackage[i]
			borrowerId = adminAgentAcc.LoanPackage[i].BorrowerId
		}
	}
	if borrowerId == nil {
		return nil, nil
	}
	borrowerAcc := borrower{}
	borrowerAsbytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)

	for i := range borrowerAcc.LoanPacks {
		if borrowerAcc.LoanPacks[i].Id == loanPackageId {
			borrowerAcc.LoanPacks[i].Status = status
		}
	}

	newBorrowerAbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}

	newAdminagentAsbytes, _ := json.Marshal(adminAgentAcc)

	err = stub.PutState(administrativeAgentId, newAdminagentAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}
	var loanStack []loanPackage
	loansAsbytes, _ := stub.GetState(loanStackKey)
	_ = json.Unmarshal(loansAsbytes, &loanStack)
	if loanPack.Status == `verified` {
		loanStack = append(loanStack, loanPack)
	}

	return nil, nil
}
