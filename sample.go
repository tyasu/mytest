package main

import (
	  "encoding/json"
		"errors"
		"fmt"
		"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//json data format
type Info struct {
	thing string `json:"thing"`
	madeBy string `json:"madeBy"`
	amount string `json:"amount"`
  createdAt string `json:"createdAt"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

fmt.Println("Initializing data")
var blank []string
	blankBytes, _ := json.Marshal(&blank)
	err := stub.PutState("data", blankBytes)
    if err != nil {
        fmt.Println("Failed to initialize")
    }

	fmt.Println("Initialization complete")
	 return nil, nil

}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    fmt.Println("Invoke is running " + function +", with args",len(args))

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown invoke function name")
}


// Invoke is a no-op
func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var data Info
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	data.thing = args[0]
	data.madeBy = args[1]
	data.amount = args[2]
	data.createdAt = args[3]


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
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {													//read a variable
		info, err := t.read(stub, args)
		if err != nil {
			fmt.Println("Error Getting info")
			return nil, err
		} else {
			infoBytes, err1 := json.Marshal(&info)
			if err1 != nil {
				fmt.Println("Error marshalling the info")
				return nil, err1
			}	
			fmt.Println("All success, returning the info")
			return infoBytes, nil		 
		}
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) (Info, error) {
    var key, jsonResp string
    var err error
		var info Info

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
		err = json.Unmarshal(valAsbytes, &info)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }

    return info, nil
}
