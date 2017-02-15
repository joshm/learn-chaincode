/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

  err := stub.PutState("hello_world", []byte(args[0]))
  if err != nil {
    return shim.Error(err.Error())
  }

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
  if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub)
	} else if function == "write" {
    return t.write(stub, args)
  }

	fmt.Println("invoke did not find func: " + function)					//error

	return shim.Error("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  fmt.Println("ex02 write")
  var key, value string

  if len(args) != 2 {
    return shim.Error("Incorrect number of args. Expecting 2")
  }

  key = args[0]
  value = args[1]
  err := stub.PutState(key, []byte(value))
  if err != nil {
    return shim.Error(err.Error())
  }

  return shim.Success(nil)


}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {											//read a variable
    return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

  return shim.Error("Received unkown function query: " + function)
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error

  if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	key = args[0]
  Avalbytes, err := stub.GetState(key)
  if err != nil {
    jsonResp = "{\"Error\":\"Falied to get state for " + key + "\"}"
    return shim.Error(jsonResp)
  }

  if Avalbytes == nil {
    jsonResp = "{\"Error\":\"Nil amount for " + key + "\"}"
    return shim.Error(jsonResp)
  }

  return shim.Success(Avalbytes)
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


