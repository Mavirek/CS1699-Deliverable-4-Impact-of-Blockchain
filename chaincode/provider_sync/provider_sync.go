package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Provider struct {
	ID      string `json:"id"`
	First   string `json:"first"`
	Last    string `json:"last"`
	Type    string `json:"type"`
	Address string `json:"address"`
	State   string `json:"state"`
	ZIP     string `json:"zip"`
	Payers  string `json:"payers"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryProvider" {
		return s.query(APIstub, args)
	} else if function == "init" {
		return s.Init(APIstub)
	} else if function == "addProvider" {
		return s.add(APIstub, args)
	} else if function == "updateProvider" {
		return s.updateProvider(APIstub, args)
	} else if function == "queryUpTo" {
		return s.queryUpTo(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function")
}

//Usage in query.js
//fcn: 'queryProvider'
//args: 'id'

func (s *SmartContract) query(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	prov, _ := APIstub.GetState(args[0])
	return shim.Success(prov)
}

// Usage in invoke.js
// fcn: 'addProvider'
// args: 'first','last','type','address','state','zip',payers'

func (s *SmartContract) add(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var provider = Provider{ID: args[0], First: args[1], Last: args[2], Type: args[3], Address: args[4], State: args[5], ZIP: args[6], Payers: args[7]}

	prov, _ := json.Marshal(provider)
	APIstub.PutState(args[0], prov)

	return shim.Success(nil)
}

// Usage in invoke.js - will do an add() if the provider doesn't already exist
// fcn: 'updateProvider'
// args: 'first','last','type','address','state','zip',payers'

func (s *SmartContract) updateProvider(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	prov, _ := APIstub.GetState(args[0])
	//all fields must be provided in args
	if prov != nil {
		//provider exists
		provider := Provider{}
		json.Unmarshal(prov, &provider)
		provider.ID = args[0]
		provider.First = args[1]
		provider.Last = args[2]
		provider.Type = args[3]
		provider.Address = args[4]
		provider.State = args[5]
		provider.ZIP = args[6]
		provider.Payers = args[7]
		prov, _ = json.Marshal(provider)
		APIstub.PutState(args[0], prov)
	} else {
		//provider doesn't exist
		s.add(APIstub, args)
	}

	return shim.Success(nil)
}

//Usage in query.js
// fcn: 'queryUpTo'
// args: '<number of providers to query up to>'

func (s *SmartContract) queryUpTo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	start := "1"           //inclusive
	end := string(args[0]) //exclusive

	resultsIter, err := APIstub.GetStateByRange(start, end)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buff bytes.Buffer
	buff.WriteString("[")

	writtenArrMember := false
	for resultsIter.HasNext() {
		queryResp, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if writtenArrMember == true {
			buff.WriteString(",")
		}
		buff.WriteString("{\"Key\":")
		buff.WriteString("\"")
		buff.WriteString(queryResp.Key)
		buff.WriteString("\"")
		buff.WriteString(", \"Record\":")
		buff.WriteString(string(queryResp.Value))
		buff.WriteString("}")
		writtenArrMember = true
	}
	buff.WriteString("]")

	fmt.Printf(buff.String())

	return shim.Success(buff.Bytes())
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
