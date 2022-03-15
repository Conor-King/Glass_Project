/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
        "os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// GlassChainCode implementation
type GlassChainCode struct {
}

type GlassResource struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	CID       string `json:"cid"`    //the fieldtags are needed to keep case from bouncing around
	URI       string `json:"uri"`
}

type GlassResourceKey struct { 
	ObjectType string `json:"docType"` 
	CID       string `json:"cid"`    
	Key   string `json:"key"` //We assume this will be a symmetric key (AES-CBC 256) for now.
}



// Init initializes chaincode
// ===========================
func (t *GlassChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *GlassChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {
	case "createGlassResource":
		//Create new Glass Resource Entry
		return t.createGlassResource(stub, args)
	case "readGlassResource":
		//read Glass Resource 
		return t.readGlassResource(stub, args)
	case "readGlassResourceKey":
		//read Glass Resource Key
		return t.readGlassResourceKey(stub, args)
	case "deleteGlassResource": //TODO: implement delete GlassResourceKey?
		//delete Glass Resource
		return t.delete(stub, args)
	case "getGlassResourceHash":
		// get Glass Resource hash for collectionGlassResources
		return t.getGlassResourceHash(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

// ============================================================
// createGlassResource - create new Glass resource and store into chaincode state
// ============================================================
func (t *GlassChainCode) createGlassResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	type GlassResourceTransientInput struct {
		CID string `json:"cid"`    //the fieldtags are needed to keep case from bouncing around
		Key string `json:"key"`
		URI string `json:"URI"`
	}

	// ==== Input sanitation ====
	fmt.Println("- start Glass Resource")

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Private data must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	GlassResourceJsonBytes, ok := transMap["GlassResource"]
	if !ok {
		return shim.Error("GlassResource must be a key in the transient map")
	}

	if len(GlassResourceJsonBytes) == 0 {
		return shim.Error("GlassResource value in the transient map must be a non-empty JSON string")
	}

	var GlassResourceInput GlassResourceTransientInput
	err = json.Unmarshal(GlassResourceJsonBytes, &GlassResourceInput)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(GlassResourceJsonBytes))
	}

	if len(GlassResourceInput.CID) == 0 {
		return shim.Error("CID field must be a non-empty string")
	}
	if len(GlassResourceInput.Key) == 0 {
		return shim.Error("Key field must be a non-empty string")
	}
	if len(GlassResourceInput.URI) == 0 {
		return shim.Error("URI field must be a non-empty string")
	}

	// ==== Check if GlassResource already exists ====
	GlassResourceAsBytes, err := stub.GetPrivateData("collectionGlassResources", GlassResourceInput.CID)
	if err != nil {
		return shim.Error("Failed to get GlassResource: " + err.Error())
	} else if GlassResourceAsBytes != nil {
		fmt.Println("This GlassResource already exists: " + GlassResourceInput.CID)
		return shim.Error("This GlassResource already exists: " + GlassResourceInput.CID)
	}

	// ==== Create GlassResource object and marshal to JSON ====
	GlassResource := &GlassResource{
		ObjectType: "GlassResource",
		CID:       GlassResourceInput.CID,
		URI:      GlassResourceInput.URI,
	}
	GlassResourceJSONasBytes, err := json.Marshal(GlassResource)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save GlassResource to state ===
	err = stub.PutPrivateData("collectionGlassResources", GlassResourceInput.CID, GlassResourceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Create Glass Resource Key object with fdom, marshal to JSON, and save to state ====
	GlassResourceKey := &GlassResourceKey{
		ObjectType: "GlassResourceKey",
		CID:       GlassResourceInput.CID,
		Key:      GlassResourceInput.Key,
	}
	
	GlassResourceKeyBytes, err := json.Marshal(GlassResourceKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutPrivateData("collectionGlassResourcesKeys", GlassResourceInput.CID, GlassResourceKeyBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== GlassResource saved and indexed. Return success ====
	fmt.Println("- end init metrics")
	return shim.Success(nil)
}

// ===============================================
// readGlassResource - read Glass resource from chaincode state
// ===============================================
func (t *GlassChainCode) readGlassResource(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var cid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting CID of the Glass resource entry to query")
	}

	cid = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionGlassResources", cid) //get the metrics from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + cid + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Glass Resource entry does not exist: " + cid + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readGlassResourceKey - read Glass resource (decryption) key from chaincode state
// ===============================================
func (t *GlassChainCode) readGlassResourceKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var cid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting cid of the Glass resource entry to query")
	}

	cid = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionGlassResourcesKeys", cid) //get the GlassResourceKeys from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + cid + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Glass Resource Key entry does not exist: " + cid + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// getGlassResourceHash - get Glass resource entry private data hash for collectionGlassResources from chaincode state (untested)
// ===============================================
func (t *GlassChainCode) getGlassResourceHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the GlassResource to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetPrivateDataHash("collectionGlassResources", name)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get GlassResource private data hash for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"GlassResource data hash does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove GlassResource Entry key/value pair from state (untested)
// ==================================================
func (t *GlassChainCode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("- start delete GlassResource entry")

	type GlassResourceDeleteTransientInput struct {
		Name string `json:"name"`
	}

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. GlassResource name must be passed in transient map.")
	}

	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	GlassResourceDeleteJsonBytes, ok := transMap["GlassResource_delete"]
	if !ok {
		return shim.Error("GlassResource_delete must be a key in the transient map")
	}

	if len(GlassResourceDeleteJsonBytes) == 0 {
		return shim.Error("GlassResource_delete value in the transient map must be a non-empty JSON string")
	}

	var GlassResourceDeleteInput GlassResourceDeleteTransientInput
	err = json.Unmarshal(GlassResourceDeleteJsonBytes, &GlassResourceDeleteInput)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(GlassResourceDeleteJsonBytes))
	}

	if len(GlassResourceDeleteInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}

	valAsbytes, err := stub.GetPrivateData("collectionGlassResources", GlassResourceDeleteInput.Name) //get the metrics from chaincode state

	var GlassResourceToDelete GlassResource
	err = json.Unmarshal([]byte(valAsbytes), &GlassResourceToDelete)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(valAsbytes))
	}

	// delete the metrics from state
	err = stub.DelPrivateData("collectionGlassResources", GlassResourceDeleteInput.Name)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}


func main() {
	err := shim.Start(&GlassChainCode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting GlassChainCode chaincode: %s", err)
		os.Exit(2)
	}
}
