package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

// CandidateInfoStore is a high level smart contract that CandidateInfoStores together business artifact based smart contracts
type CandidateInfoStore struct {

}

type CandidateDetails struct{
    CandidateId string `json:"candidateId"`
	Title string `json:"title"`
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	DOB string `json:"dob"`
	EmailID string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
	UniqueIdType string `json:"uniqueIdType"`
	UniqueIdNumber string `json:"uniqueIdNumber"`
	Nationality string `json:"nationality"`
	Address string `json:"address"`
	Country string `json:"country"`
	City string `json:"city"`
	Zip string `json:"zip"`
	State string `json:"state"`
}

func (t *CandidateInfoStore) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("CandidateDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create application Table
	err = stub.CreateTable("CandidateDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "candidateId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "emailId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phoneNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "uniqueIdType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "uniqueIdNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "State", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating CandidateDetails.")
	}
	// setting up the users role
	stub.PutState("user_type1_1", []byte("Govt"))
	stub.PutState("user_type1_2", []byte("IBM"))
	stub.PutState("user_type1_3", []byte("CTS"))
	stub.PutState("user_type1_4", []byte("????"))	
	
	return nil, nil
}

//registerUser to register a user
func (t *CandidateInfoStore) CandidateRegister(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 16 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 16. Got: %d.", len(args))
		}
		candidateId:=args[0]
		title:=args[1]
		gender:=args[2]
		firstName:=args[3]
		lastName:=args[4]
		dob:=args[5]
		emailId:=args[6]
		phoneNumber:=args[7]
		uniqueIdType:=args[8]
		uniqueIdNumber:=args[9]
		nationality:=args[10]
		address :=args[11]
		country:=args[12]
		city:=args[13]
		zip:=args[14]
		State:=args[15]
		
		assignerOrg, err := stub.GetState(candidateId)
		if assignerOrg !=nil{
			return nil, fmt.Errorf("Candidate already registered %s",candidateId)
		} else if err !=nil{
			return nil, fmt.Errorf("System error")
		}
		
		
		// Insert a row
		ok, err := stub.InsertRow("CandidateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: candidateId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: emailId}},
				&shim.Column{Value: &shim.Column_String_{String_: phoneNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdType}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: State}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

func (t *CandidateInfoStore) getAllCandidate(stub shim.ChaincodeStubInterface , args []string) ([]byte, error) {

	// Get the row pertaining to this candidateId
	var columns []shim.Column
	rows, err := stub.GetRows("CandidateDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application \"}"
		return nil, errors.New(jsonResp)
	}

	res2E := []*CandidateDetails{}
	
	for row := range rows{
	newCan := new (CandidateDetails)
	newCan.CandidateId = row.Columns[0].GetString_()
	newCan.Title = row.Columns[1].GetString_()
	newCan.Gender = row.Columns[2].GetString_()
	newCan.FirstName = row.Columns[3].GetString_()
	newCan.LastName = row.Columns[4].GetString_()
	newCan.DOB = row.Columns[5].GetString_()
	newCan.EmailID = row.Columns[6].GetString_()
	newCan.PhoneNumber = row.Columns[7].GetString_()
	newCan.UniqueIdType = row.Columns[8].GetString_()
	newCan.UniqueIdNumber = row.Columns[9].GetString_()
	newCan.Nationality = row.Columns[10].GetString_()
	newCan.Address = row.Columns[11].GetString_()
	newCan.Country = row.Columns[12].GetString_()
	newCan.City = row.Columns[13].GetString_()
	newCan.Zip = row.Columns[14].GetString_()
	newCan.State = row.Columns[15].GetString_()
	res2E=append(res2E,newCan)
}

    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))

	return mapB, nil

}

func (t *CandidateInfoStore) getCandidate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting candidateId to query")
	}

	candidateId := args[0]


	// Get the row pertaining to this candidateId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: candidateId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("CandidateDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + candidateId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + candidateId + "\"}"
		return nil, errors.New(jsonResp)
	}


	res2E := CandidateDetails{}
	res2E.CandidateId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.DOB = row.Columns[5].GetString_()
	res2E.EmailID = row.Columns[6].GetString_()
	res2E.PhoneNumber = row.Columns[7].GetString_()
	res2E.UniqueIdType = row.Columns[8].GetString_()
	res2E.UniqueIdNumber = row.Columns[9].GetString_()
	res2E.Nationality = row.Columns[10].GetString_()
	res2E.Address = row.Columns[11].GetString_()
	res2E.Country = row.Columns[12].GetString_()
	res2E.City = row.Columns[13].GetString_()
	res2E.Zip = row.Columns[14].GetString_()
	res2E.State = row.Columns[15].GetString_()

    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))

	return mapB, nil

}

// Invoke invokes the chaincode
func (t *CandidateInfoStore) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "CandidateRegister" {
		t := CandidateInfoStore{}
		return t.CandidateRegister(stub, args)	
	} 

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *CandidateInfoStore) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getCandidate" { 
		t := CandidateInfoStore{}
		return t.getCandidate(stub, args)
	}else if function == "getAllCandidate"{
        t := CandidateInfoStore{}
		return t.getAllCandidate(stub, args)  
    } 
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(CandidateInfoStore))
	if err != nil {
		fmt.Printf("Error starting CandidateInfoStore: %s", err)
	}
} 
