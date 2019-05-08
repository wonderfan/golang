/**
 * One chain code and multiple smart contracts
 * File and Invoice operations
 * Interface definitions

type FileAsset interface {
	Store(uid string,fileName string, digest string, owner string) pb.Response
	Query(uid string, owner string) pb.Response
	Source(uid string, amount string, payDate string, seller string, buyer string, fileUid string) pb.Response
}

type InvoiceAsset interface {
	Generate(uid string, amount string, holder string) pb.Response
	Transfer(uid string, holder string) pb.Response
	Divide(originalUid string,uid string, amount string, holder string) pb.Response
	Mortgage(uid string) pb.Response
	Close(uid string) pb.Response
}

 *
 */

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"strconv"
)

const (
	tableSuffix = "uid"
	keyDelimiter = "~"
	tableFile = "file"
	tableInvoice = "invoice"
	tableAsset = "asset"
	tableFileCompositeKey = tableFile + keyDelimiter + tableSuffix
	tableInvoiceCompositeKey = tableInvoice + keyDelimiter + tableSuffix
	tableAssetCompositeKey = tableAsset + keyDelimiter + tableSuffix
	statusCreated = "created"
	statusMortgaged = "mortgaged"
	statusClosed = "closed"
)

type File struct {
	uid string
	fileName string
	digest string
	owner string
}

type Invoice struct {
	uid string
	amount string
	payDate string
	seller string
	buyer string
	fileUid string
}

type Asset struct {
	uid string
	amount string
	balance string
	status string
	holder string
	invoiceUid string
}

type SupplyChain struct {

}

func (supplyChain *SupplyChain) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// add the initialization logic and process
	return shim.Success(nil)
}

func (supplyChain *SupplyChain) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// dispatch the function invocation to different methods
	function, args := stub.GetFunctionAndParameters()

	if args == nil || len(args) < 0 {
		return shim.Error("Invalid function parameters.")
	}

	if function == "store" {
		return supplyChain.Store(args, stub)
	} else if function == "query" {
		return supplyChain.Store(args, stub)
	} else if function == "source" {
		return supplyChain.Source(args, stub)
	} else if function == "generate" {
		return supplyChain.Generate(args, stub)
	} else if function == "transfer" {
		return supplyChain.Transfer(args, stub)
	} else if function == "divide" {
		return supplyChain.Divide(args, stub)
	} else if function == "mortgage" {
		return supplyChain.Mortgage(args, stub)
	}else if function == "close" {
		return supplyChain.Close(args, stub)
	}

	return shim.Error("Invalid function name.")
}

func (supplyChain *SupplyChain) Store(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	file := new(File)
	file.uid = args[0]
	file.fileName = args[1]
	file.digest = args[2]
	file.owner = args[3]

	compositeKey, err := stub.CreateCompositeKey(tableFileCompositeKey,[]string{tableFile,file.uid})

	if err != nil {
		return shim.Error(err.Error())
	}

	jsonValue, err := json.Marshal(file)
	if err != nil {
		return shim.Error(err.Error())
	}

	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Query(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	uid := args[0]
	compositeKey, err := stub.CreateCompositeKey(tableFileCompositeKey,[]string{tableFile,uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	jsonValue, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Source(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	invoice := new(Invoice)
	invoice.uid = args[0]
	invoice.amount = args[1]
	invoice.payDate = args[2]
	invoice.seller = args[3]
	invoice.buyer = args[4]
	invoice.fileUid = args[5]

	compositeKey, err := stub.CreateCompositeKey(tableInvoiceCompositeKey,[]string{tableInvoice,invoice.uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	jsonValue, err := json.Marshal(invoice)
	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Generate(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	asset := new(Asset)
	asset.uid = args[0]
	asset.amount = args[1]
	asset.holder = args[2]
	asset.invoiceUid = args[3]
	asset.balance = asset.amount
	asset.status = statusCreated

	invoiceCompositeKey, err := stub.CreateCompositeKey(tableInvoiceCompositeKey,[]string{tableInvoice,asset.invoiceUid})
	if err != nil {
		return shim.Error(err.Error())
	}

	invoiceRecord, err := stub.GetState(invoiceCompositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}
    if invoiceRecord == nil {
		return shim.Error(fmt.Sprintf("original data does not exist : %s", asset.invoiceUid))
	}
	compositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,asset.uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	jsonValue, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Transfer(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	uid := args[0]
	holder := args[1]
	compositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	dbData, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if dbData == nil {
		return shim.Error(fmt.Sprintf("asset does not exist: %s",uid))
	}

	asset := new(Asset)

	err = json.Unmarshal(dbData,asset)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset.holder = holder
	jsonValue, err := json.Marshal(asset)

	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Divide(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	uid := args[0]
	derivedAsset := new(Asset)
	derivedAsset.uid = args[2]
	derivedAsset.amount = args[3]
	derivedAsset.holder = args[4]
	derivedAsset.invoiceUid = args[5]

	compositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	dbData, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if dbData == nil {
		return shim.Error(fmt.Sprintf("asset does not exist: %s",uid))
	}

	originalAsset := new(Asset)

	err = json.Unmarshal(dbData,originalAsset)

	if err != nil {
		return shim.Error(err.Error())
	}
	amount, err := strconv.Atoi(originalAsset.amount)
	if err != nil {
		return shim.Error(err.Error())
	}
	dividedAmount, err := strconv.Atoi(derivedAsset.amount)
	if err != nil {
		return shim.Error(err.Error())
	}
	originalAsset.balance = strconv.Itoa(amount - dividedAmount)

	derivedCompositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,derivedAsset.uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	jsonValue, err := json.Marshal(derivedAsset)
	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(derivedCompositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}

	origValue, err := json.Marshal(originalAsset)
	if err != nil {
		return shim.Error(err.Error())
	}
	errState = stub.PutState(compositeKey,origValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
    return shim.Success(origValue)
}

func (supplyChain *SupplyChain) Mortgage(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	uid := args[0]
	compositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	dbData, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if dbData == nil {
		return shim.Error(fmt.Sprintf("asset does not exist: %s",uid))
	}

	asset := new(Asset)

	err = json.Unmarshal(dbData,asset)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset.status = statusMortgaged
	jsonValue, err := json.Marshal(asset)

	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func (supplyChain *SupplyChain) Close(args []string, stub shim.ChaincodeStubInterface) pb.Response {
	uid := args[0]
	compositeKey, err := stub.CreateCompositeKey(tableAssetCompositeKey,[]string{tableAsset,uid})
	if err != nil {
		return shim.Error(err.Error())
	}
	dbData, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if dbData == nil {
		return shim.Error(fmt.Sprintf("asset does not exist: %s",uid))
	}

	asset := new(Asset)

	err = json.Unmarshal(dbData,asset)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset.status = statusClosed
	jsonValue, err := json.Marshal(asset)

	if err != nil {
		return shim.Error(err.Error())
	}
	errState := stub.PutState(compositeKey,jsonValue)
	if errState != nil {
		return shim.Error(errState.Error())
	}
	return shim.Success(jsonValue)
}

func main() {
	err := shim.Start(new(SupplyChain))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
