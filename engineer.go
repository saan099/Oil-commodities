package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *Oilchain) InitEngineer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, errors.New(`wrong number of arguments`)
	}
	engineerId := args[0]
	engineerName := args[1]
	engineerEmail := args[2]
	registrationId := args[3]

	engineerAcc := engineer{}
	////////////////////////////////////////////////
	//       engineer account parsing
	////////////////////////////////////////////////
	engineerAcc.Id = engineerId
	engineerAcc.Name = engineerName
	engineerAcc.Email = engineerEmail
	engineerAcc.RegistratonID = registrationId

	engineerAsbytes, _ := json.Marshal(engineerAcc)
	err := stub.PutState(engineerId, engineerAsbytes)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *Oilchain) MakeReserveReport(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 5 {
		return nil, errors.New(`wrong number of arguments`)
	}

	engineerId := args[0]
	borrowerid := args[1]
	reportId := args[2]
	date := args[3]
	developedCrude := args[4]
	undevelopedCrude := args[5]

	reserveRep := reserveReport{}
	////////////////////////////////////////////////
	//       reserve report  parsing
	////////////////////////////////////////////////
	reserveRep.Id = engineerId
	reserveRep.BorrowerId = borrowerid
	reserveRep.EngineerId = engineerId
	reserveRep.Date = date
	reserveRep.DevelopedCrude = strconv.ParseFloat(developedCrude, 64)
	reserveRep.UndevelopedCrude = strconv.ParseFloat(undevelopedCrude, 64)

	engineerAcc := engineer{}
	engineerAsbytes := stub.GetState(engineerId)
	_ = json.Unmarshal(engineerAsbytes, &engineerAcc)
	engineerAcc.ReserveReports = append(engineerAcc.ReserveReports, reserveRep)

	newEngineerAsbytes, _ := json.Marshal(engineerAcc)
	err := stub.PutState(engineerId, newEngineerAsbytes)
	if err != nil {
		return nil, errors.New(`didnt write state`)
	}

	return nil, nil
}
