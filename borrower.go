package main

import (
	"encoding/json"
	"errors"
	"strconv"

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

	var financialrep = financialReport{}
	//////////////////////////////////////////////////
	//  financialrep data parsing
	//////////////////////////////////////////////////
	financialrep.Id = reportId
	financialrep.CreditPeriod, _ = strconv.Atoi(creditDays)
	financialrep.Date = date
	financialrep.LoanAmount, _ = strconv.ParseFloat(loanAmount, 64)
	financialrep.Status = `pending`

	borrowerAcc := borrower{}
	borrowerAsytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsytes, &borrowerAcc)
	borrowerAcc.FinancialReports = append(borrowerAcc.FinancialReports, financialrep)

	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if err != nil {
		return nil, errors.New(`couldnt write state`)
	}

	return nil, nil
}

func (t *Oilchain) AddComplianceCertificate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.New(`wrong number of arguments`)
	}
	borrowerId := args[0]
	complianceRepId := args[1]
	date := args[2]

	complianceCert := complianceCertificate{}
	//////////////////////////////////////////////////
	//  complianceCert data parsing
	//////////////////////////////////////////////////
	complianceCert.BorrowerId = borrowerId
	complianceCert.Id = complianceRepId
	complianceCert.Date = date

	borrowerAcc := borrower{}
	borrowerAsbytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)
	borrowerAcc.ComplianceCertificates = append(borrowerAcc.ComplianceCertificates, complianceCert)
	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if err != nil {
		return nil, errors.New(`not written in state.`)
	}
	return nil, nil
}

/*
func (t *Oilchain) RequestReserveReport(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args)!=2{
    return nil, errors.New(`wrong number of arguments.`)
  }
  borrowerId:=args[0]
  engineerId:=args[1]



  return nil,nil
}
*/
