// This is attepmt to store maps on the ledger, it is a valid chaincode,
// but when querying or transfering its acting funny, not suitable
// to convert map to a byte array. The idea was to map
// uiqueid:cid, and have them as different elements which we can transfer
// from entity to entity, but im not sure that this is possible with the
// chaincode functionality we have, as it requires the data to be converted in
// a byte array, before appending it to the Hyperledger

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init method gets called")
	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	Aval:= make(map[string]string)
	Bval := make(map[string]string)
	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	// a key value b key value
	// 0 1     2   3   4    5
	// Initialize the chaincode
	A = args[0]
	keyA := args[1]
	Aval[keyA] = args[2]

	B = args[3]
	keyB := args[4]
	Bval[keyB] = args[5]

	AvalBytes, _ := json.Marshal(Aval[keyA])
	BvalBytes, _ := json.Marshal(Bval[keyB])



	for key, value := range Aval {
		fmt.Println("key:", key, "value:", value)
		}
	for key, value := range Bval {
			fmt.Println("key:", key, "value:", value)
			}



	// Write the state to the ledger
	err = stub.PutState(A, AvalBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, BvalBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke method gets called")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	}  else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("invoke method gets called")
	var A, B string    // Entities
	var Aval, Bval string // Asset holdings
	var X string          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("A val bytes == nil, Entity not found")
	}
	Aval = string(Avalbytes)

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("B val bytes == nil, Entity not found")
	}
	Bval= string(Bvalbytes)

	// Perform the execution
	X = args[2]
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	//Aval = Aval - X
	Bval = Bval + " " + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(Aval))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(Bval))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("query method gets called")
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
