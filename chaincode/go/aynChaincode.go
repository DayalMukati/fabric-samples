package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for a Employee
type SmartContract struct {
	contractapi.Contract
}

// Employee describes basic details of what makes up a Employee
type Employee struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	City 	  string `json:"city"`
	Country   string `json:"country"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Employee
}


func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	fmt.Sprintf("EmployeeCC Successfully Init")

	return nil
}


// CreateEmployee adds a new employee to the world state with given details
func (s *SmartContract) CreateEmployee(ctx contractapi.TransactionContextInterface, employeeNumber string, name string, gender string, city string, country string) error {
	employee := Employee{
		Name:    name,
		Gender:  gender,
		City:    city,
		Country: country,
	}

	employeeAsBytes, _ := json.Marshal(employee)

	return ctx.GetStub().PutState(employeeNumber, employeeAsBytes)
}

// QueryAllEmployees returns all employees found in world state
func (s *SmartContract) QueryAllEmployees(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		employee := new(Employee)
		_ = json.Unmarshal(queryResponse.Value, employee)

		queryResult := QueryResult{Key: queryResponse.Key, Record: employee}
		results = append(results, queryResult)
	}

	return results, nil
}


func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create employee chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting employee chaincode: %s", err.Error())
	}
}