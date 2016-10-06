package main

import (
		"errors"
		"fmt"
		"github.com/hyperledger/fabric/core/chaincode/shim"
		"strconv"
)

// myChaincode example simple Chaincode implementation
type myChaincode struct {
}


//json data format
type Info struct {
	thing string `json:"thing"`
	madeBy string `json:"madeBy"`
	amount string `json:"amount"`
  createdAt string `json:"createdAt"`
}


func (t *myChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    fmt.Println("Invoke is running " + function +", with args",len(args))

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.Write(stub, args)
	}
	return nil, errors.New("Received unknown invoke function name")
}




// Init takes a string and int. These are stored as a key/value pair in the state
func (t *myChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

var info Info

info.thing="apple"
info.madeBy="me"
info.amount="1"
info.createdAt="2016"

	if len(args) != 1 {
			 return nil, errors.New("Incorrect number of arguments. Expecting 1")
	 }

	 jsonAsBytes, _ := json.Marshal(info)
	 err := stub.PutState("data", jsonAsBytes)
	 if err != nil {
			 return nil, err
	 }

	 return nil, nil

}

// Invoke is a no-op
func (t *myChaincode) Write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var data Info
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	data = args[0]
	fmt.Printf("data = %d\n", data)

	// Write the state to the ledger - this put is legal within Run
	jsonAsBytes, _ := json.Marshal(data)
	err = stub.PutState("data", jsonAsBytes)
	if err != nil {
		return nil, errors.New("Error putting data on ledger")
}
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

func (t *myChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)

    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}


func main() {
	err := shim.Start(new(myChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
