//This is simply a basic chaincode from minifab converted to work with strings.
//Everything should be working
//The rest is developed based on this ss

package main

import (
	"fmt"
	"strings"

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
	var Aval, Bval string // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval = args[1]

	B = args[2]
	Bval = args[3]

	//fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
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
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}


// Transaction makes payment of X units from A to B
//'"invoke","a","b","uniqueid2"'
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("invoke method gets called")
	var A, B string    // Entities
	var Aval, Bval string // Asset holdings
	var X string          // Transaction unit unique id

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	if A == "add" {
		B = args[1]
		X = args[2]
		Avalbytes, err := stub.GetState(B)
		if err != nil {
			return shim.Error("Failed to get state")
		}
		if Avalbytes == nil {
			return shim.Error("A val bytes == nil, Entity not found")
		}
		Aval = string(Avalbytes)
		Aval = Aval + " " + X
	
		err = stub.PutState(B, []byte(Aval))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	} else {


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

	//Aval = arg1 = string = elements in A
	//Bval = arg2 = string = elements in B
	//X = arg 3 = string = unique id

	sliceA := strings.Split(Aval, " ")
	//sliceB := strings.Split(Aval, " ")

	//loop each elemnt and search for the uniqueid
	for i := range sliceA {
		res := strings.HasPrefix(sliceA[i], X)
		//res is true when the element begins with arg3
		if  res  {
		//actions with the element[i]
			Bval = Bval + " " + sliceA[i]
			sliceA = RemoveIndex(sliceA, i)
			break;
		}
	}
	Aval = strings.Join(sliceA, " ")

	//Aval = Aval - X
	//Bval = Bval + " " + X
	//fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

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
