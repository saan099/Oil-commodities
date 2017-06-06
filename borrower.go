package main


import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) InitBorrower(stub shim.ChaincodeStubInterface,args []string) ([]byte,error) {

  if len(args) !=4 {
    return nil, errors.New(`Wrong number of arguments.`);
  }
  var borrowerId=args[0]
  var name=args[1]
  var registrationId=args[2]
  var email=args[3]
  compCert:=[]complianceCertificate
  financialRep:=[]financialReport
  reserveRep:=[]reserveReport

  var borrowerAcc=borrower{}
  borrowerAcc.Id=borrowerId
  borrowerAcc.Name=name
  borrowerAcc.RegistratonID=registrationId
  borrowerAcc.Email=email
  borrowerAcc.ComplianceCertificates=compCert
  borrowerAcc.FinancialReports=financialRep
  borrowerAcc.ReserveReports=reserveRep
  borrowerAsbytes,_:=json.Marshal(borrowerAcc)
  err:=stub.PutState(borrowerId,borrowerAsbytes)

  if err!=nil {
    return nil, errors.New(`Problem in writing state.`)
  }


  return nil,nil
}
