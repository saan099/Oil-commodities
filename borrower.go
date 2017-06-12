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
	borrowerAcc := borrower{}
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
	financialrep.BorrowerId = borrowerId

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

func (t *Oilchain) CreateLoanPackage(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) < 4 {
		return nil, errors.New("wrong number of arguments")
	}

	borrowerId := args[0]
	loanPackageId := args[1]
	amountRequested := args[2]
	financialId := args[3]
	complianceId := args[4]
	reserveId := args[5]
	numOfDocs, _ := strconv.Atoi(args[6])
	administrativeAgentId := args[7]

	borrowerAcc := borrower{}
	borrowerAsbytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)

	var docs []document
	loanPack := loanPackage{}
	////////////////////////////////////////////////////
	//          loan package parsing
	////////////////////////////////////////////////////
	loanPack.Id = loanPackageId
	loanPack.AmountRequested, _ = strconv.ParseFloat(amountRequested, 64)
	loanPack.BorrowerId = borrowerId
	for i := 0; i < len(borrowerAcc.FinancialReports); i++ {
		if borrowerAcc.FinancialReports[i].Id == financialId {
			loanPack.FinancialReport = borrowerAcc.FinancialReports[i]
		}
	}
	for i := 0; i < len(borrowerAcc.ComplianceCertificates); i++ {
		if borrowerAcc.ComplianceCertificates[i].Id == complianceId {
			loanPack.ComplianceCertificate = borrowerAcc.ComplianceCertificates[i]
		}
	}
	for i := 0; i < len(borrowerAcc.ReserveReports); i++ {
		if borrowerAcc.ReserveReports[i].Id == reserveId {
			loanPack.ReserveReport = borrowerAcc.ReserveReports[i]
		}
	}

	loanPack.BorrowerName = borrowerAcc.Name
	loanPack.Status = `pending`
	for i := 7; i < numOfDocs*2; i++ {
		doc := document{}
		doc.DocName = args[i]
		i = i + 1
		doc.Id = args[i]
		docs = append(docs, doc)
	}
	loanPack.Documents = docs

	borrowerAcc.LoanPacks = append(borrowerAcc.LoanPacks, loanPack)
	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}
	adminAcc := administrativeAgent{}
	adminAsbytes, _ := stub.getState(administrativeAgentId)
	_ = json.Unmarshal(adminAsbytes, &adminAcc)
	adminAcc.LoanPackage = append(adminAcc.LoanPackage, loanPack)
	newAdminAsbytes, _ := json.Marshal(adminAcc)
	err = stub.PutState(administrativeAgentId, newAdminAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
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
