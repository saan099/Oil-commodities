package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strconv"
	"time"

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

func (t *Oilchain) RequestFinancialStatement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, error.New("worng number of arguments")
	}
	requestId := args[0]
	borrowerId := args[1]
	auditorId := args[2]
	currentT := time.Now().Local()
	date := currentT.Format("02-01-2006")

	var req request
	req.Id = requestId
	req.BorrowerId = borrowerId
	req.Date = date
	req.Status = "pending"

	borrowerAcc := borrower{}
	borrowerAsbytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)
	borrowerAcc.Requests = append(borrowerAcc.Requests, req)

	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if err != nil {
		return nil, errors.New("didnt write state")
	}

	auditorAcc := auditor{}
	auditorAsbytes, _ := stub.Getstate(auditorId)
	_ = json.Unmarshal(auditorAsbytes, &auditorAcc)
	auditorAcc.Requests = append(auditorAcc.Requests, req)
	newAuditorAsbytes := json.Marshal(auditorAcc)

	err = stub.PutState(auditorId, newAuditorAsbytes)
	if err != nil {
		return nil, errors.New("didnt write state")
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
	if len(args) < 6 {
		return nil, errors.New("wrong number of arguments")
	}

	borrowerId := args[0]
	loanPackageId := args[1]
	amountRequested := args[2]
	//financialStatementNumber := strconv.Atoi(args[3])
	complianceId := args[3]
	requestTo := args[4]
	numOfDocs, _ := strconv.Atoi(args[5])

	borrowerAcc := borrower{}
	requestId, borrowerAsbytes := RequestReserveReport(stub, requestTo, borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)

	var docs []document
	loanPack := loanPackage{}
	////////////////////////////////////////////////////
	//          loan package parsing
	////////////////////////////////////////////////////
	loanPack.Id = loanPackageId
	loanPack.AmountRequested, _ = strconv.ParseFloat(amountRequested, 64)
	loanPack.BorrowerId = borrowerId

	for i := 0; i < len(borrowerAcc.ComplianceCertificates); i++ {
		if borrowerAcc.ComplianceCertificates[i].Id == complianceId {
			loanPack.ComplianceCertificate = borrowerAcc.ComplianceCertificates[i]
		}
	}

	loanPack.BorrowerName = borrowerAcc.Name
	loanPack.Status = `pending`
	var i int
	for i = 6; i < numOfDocs*2; i++ {
		doc := document{}
		doc.DocName = args[i]
		i = i + 1
		doc.Id = args[i]
		docs = append(docs, doc)
	}
	loanPack.Documents = docs
	loanPack.RequestReserveReport.RequestTo = requestTo
	for i := range borrowerAcc.Requests {
		if borrowerAcc.Requests[i].Id == requestId {
			loanPack.RequestReserveReport = borrowerAcc.Requests[i]
		}
	}
	for j := range borrowerAcc.FinancialReports {
		loanPack.FinancialReports = append(loanPack.FinancialReports, borrowerAcc.FinancialReports[j])
	}

	borrowerAcc.LoanPacks = append(borrowerAcc.LoanPacks, loanPack)
	newBorrowerAsbytes, _ := json.Marshal(borrowerAcc)
	err := stub.PutState(borrowerId, newBorrowerAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}

	return nil, nil
}

func RequestReserveReport(stub shim.ChaincodeStubinterface, to string, borrowerId string) (string, []byte) {
	req := request{}
	req.BorrowerId = borrowerId
	req.RequestTo = to
	req.Status = `pending`
	req.Type = `reserveReport`
	currentT := time.Now().Local()
	date := currentT.Format("02-01-2006")
	req.Date = date
	h := sha256.New()
	h.Write([]byte(currentT))
	req.Id = h.Sum(nil)

	engineerAcc := engineer{}
	engineerAsbytes, _ := stub.GetState(to)
	_ = json.Unmarshal(engineerAsbytes, &engineerAcc)
	engineerAcc.Requests = append(engineerAcc.Requests, req)
	newEngineerAsbytes, _ := json.Marshal(engineerAcc)
	err := stub.PutState(to, newEngineerAsbytes)
	if err != nil {
		return nil, errors.New("didnt write state")
	}

	borrowerAcc := borrower{}
	borrowerAsbytes, _ := stub.GetState(borrowerId)
	_ = json.Unmarshal(borrowerAsbytes, &borrowerAcc)
	borrowerAcc.Requests = append(borrowerAcc.Requests, req)
	borrowerAsbytes, _ := json.Marshal(borrowerAcc)
	return h.Sum(nil), borrowerAsbytes
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
