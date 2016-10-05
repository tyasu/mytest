package main

import (
		"encoding/json"
		"errors"
		"fmt"
		"github.com/hyperledger/fabric/core/chaincode/shim"
		"strconv"
)

// myChaincode example simple Chaincode implementation
type myChaincode struct {
}


//json data format
type issueInfo struct {
	thing string `json:"currency"`
	madeBy string `json:"issuedBy"`
	ammount string `json:"ammount"`
  createdAt string `json:"createdAt"`
}

func (t *myChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    fmt.Println("Run is running " + function +", with args",len(args))

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "issue" {
		return t.issue(stub, args)
	}
	return nil, errors.New("Received unknown invoke function name")
}




// Init takes a string and int. These are stored as a key/value pair in the state
func (t *myChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var info issueInfo
	var ammount int
	var data []string
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	// Initialize the chaincode
	//data = args[0]
	//minfo, err := json.Unmarshal([]byte(data, &info)
	ammount, err = strconv.Atoi(args[2])
	//if err != nil {
	//	return nil, errors.New("Expecting integer value for asset holding")
	//}
	fmt.Printf("ammount = %d\n", ammount)
	fmt.Printf("%#v\n", minfo)

	// Write the state to the ledger - this put is legal within Run
	err = stub.PutState("abc", []byte(strconv.Itoa(amount)))
	if err != nil {
		return nil, errors.New("Error putting data on ledger")
}

var empty []string
	jsonAsBytes, _ := json.Marshal(empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(moneyIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
}

	return nil, nil
}

// Invoke is a no-op
func (t *myChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *myChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {													//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}


func main() {
	err := shim.Start(new(myChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
