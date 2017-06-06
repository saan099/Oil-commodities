package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) InitBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 4 {
		return nil, errors.New(`Wrong number of arguments.`)
	}
	var borrowerId = args[0]
	var name = args[1]
	var registrationId = args[2]
	var email = args[3]
	var compCert []complianceCertificate
	var financialRep []financialReport
	var reserveRep []reserveReport
	//////////////////////////////////////////////////
	//  borrower account data parsing
	//////////////////////////////////////////////////
	var borrowerAcc = borrower{}
	borrowerAcc.Id = borrowerId
	borrowerAcc.Name = name
	borrowerAcc.RegistratonID = registrationId
	borrowerAcc.Email = email
	borrowerAcc.ComplianceCertificates = compCert
	borrowerAcc.FinancialReports = financialRep
	borrowerAcc.ReserveReports = reserveRep
	borrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, borrowerAsbytes)

	if err != nil {
		return nil, errors.New(`Problem in writing state.`)
	}

	return nil, nil
}

func (t *Oilchain) AddFinancialStatement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 5 {
		return nil, errors.New(`wrong number of arguments`)
	}

	borrowerId := args[0]
	reportId := args[1]
	creditDays := args[2]
	date := args[3]
	loanAmount := args[4]

	financialrep := financialReport{}
	//////////////////////////////////////////////////
	//  financialrep account data parsing
	//////////////////////////////////////////////////
	financialrep.Id = reportId
	financialrep.CreditPeriod = creditDays
	financialrep.Date = date
	financialrep.LoanAmount = loanAmount
	financialrep.Status = `pending`

	borrowerAcc := borrower{}
	borrowerAsytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsytes, &borrowerAcc)
	borrowerAcc.FinancialReports = append(borrowerAcc.FinancialReports, financialrep)

	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if er != nil {
		return nil, errors.New(`couldnt write state`)
	}

	return nil, nil
}
