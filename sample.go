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
	Thing string `json:"thing"`
	MadeBy string `json:"madeBy"`
	Amount string `json:"amount"`
  CreatedAt string `json:"createdAt"`
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

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

fmt.Println("Initializing data")
var blank []string
	blankBytes, _ := json.Marshal(&blank)
	err := stub.PutState("key1", blankBytes)
    if err != nil {
        fmt.Println("Failed to initialize")
    }

	fmt.Println("Initialization complete")
	 return nil, nil

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("Invoke is running " + function +", with args",len(args))

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown invoke function name")
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error
	
	data := Info{
	Thing:args[0],
	MadeBy:args[1],
	Amount:args[2],
	CreatedAt:args[3]}
	
	fmt.Println(data)
	jsonAsBytes, _ := json.Marshal(data)
	fmt.Println(jsonAsBytes)
	err = stub.PutState("key1", jsonAsBytes)
	if err != nil {
		return nil, errors.New("Error putting data on ledger")
}
	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {
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
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) (Info, error) {
    var key string
    var jsonResp string
    var err error
    var info Info
    
    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
		fmt.Println("Error retrieving info " + key)
		return info, errors.New("Error retrieving info " + key)
	}
	err = json.Unmarshal(valAsbytes, &info)
	if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return info, errors.New(jsonResp)
    }
	
    return info, nil
}
