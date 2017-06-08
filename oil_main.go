package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Oilchain struct {
}

var borrowersKey = `borrowersKey`

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

	return nil, nil
}

//Invoking functionality
func (t *Oilchain) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "init" {
		return t.Init(stub, function, args)
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
