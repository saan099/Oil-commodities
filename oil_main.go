package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Oilchain struct {
}

var borrowersKey = `borrowersKey`
var casestack = `caseStack`

func main() {

	err := shim.Start(new(Oilchain))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *Oilchain) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 0 {
		return nil, errors.New("error:A01 wrong number of aguments in initialization")
	}
	var loans []Case
	loansAsbytes, _ := json.Marshal(loans)
	err := stub.PutState(casestack, loansAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state.`)
	}
	return nil, nil
}

//Invoking functionality
func (t *Oilchain) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "initBorrower" {
		return t.InitBorrower(stub, args)
	} else if function == "addFinancialStatement" {
		return t.AddFinancialStatement(stub, args)
	} else if function == "initEngineer" {
		return t.InitEngineer(stub, args)
	} else if function == "makeReserveReport" {
		return t.MakeReserveReport(stub, args)
	} else if function == "addComplianceCertificate" {
		return t.AddComplianceCertificate(stub, args)
	} else if function == "createCase" {
		return t.CreateCase(stub, args)
	} else if function == "initAdministrativeAgent" {
		return t.InitAdministrativeAgent(stub, args)
	} else if function == "initAuditor" {
		return t.InitAuditor(stub, args)
	} else if function == "updateLoanPackage" {
		return t.UpdateLoanPackage(stub, args)
	} else if function == "requestFinancialStatement" {
		return t.RequestFinancialStatement(stub, args)
	} else if function == "initLender" {
		return t.InitLender(stub, args)
	} else if function == "makeProposals" {
		return t.MakeProposals(stub, args)
	} else if function == "makeLoanPackage" {
		return t.MakeLoanPackage(stub, args)
	} else if function == "makeCreditAgreement" {
		return t.MakeCreditAgreement(stub, args)
	} else if function == "updateReserveRep" {
		return t.UpdateReserveRep(stub, args)
	}

	return nil, errors.New("error:C01 No function called")

}

// Query data
func (t *Oilchain) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "read" {
		return t.Read(stub, args)
	} else if function == "readAllBorrowers" {
		return t.ReadAllBorrowers(stub, args)
	}

	return nil, errors.New("error:C02 No function called")
}

func (t *Oilchain) Read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("error:A04 Wrong numer of arguments")
	}

	valAsbytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, err
	}
	return valAsbytes, nil

}

func (t *Oilchain) ReadAllBorrowers(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 0 {
		return nil, errors.New("error:A05 wrong number of arguments")
	}

	return nil, nil
}
