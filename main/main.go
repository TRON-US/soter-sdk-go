package main

import (
	"context"
	"fmt"

	"github.com/TRON-US/soter-sdk-golang/soter"
	"github.com/TRON-US/soter-sdk-golang/utils"
)

func main() {
	//TestSetAutopay()
	TestBalance()
	//TestAddFile()
}

func TestAddFile()  {
	url := "http://127.0.0.1:8101"
	sh := soter.NewShell(url)

	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	out, err := sh.AddFile(context.Background(), userAddress, userAddress, privateKey, "go.mod")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}


func TestSetAutopay() {
	url := "http://127.0.0.1:8101"
	sh := soter.NewShell(url)

	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"
	payload, err := soter.GetAutopayPayload(true, userAddress, privateKey)
	if err != nil {
		panic(err)
	}

	out, err := sh.Autopay(context.Background(), payload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}

func TestBalance() {
	url := "http://127.0.0.1:8101"
	sh := soter.NewShell(url)

	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"
	rawData, err := soter.GetBalanceRawData(userAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println(rawData)
	signature, err := utils.GetSignature(rawData, privateKey)
	if err != nil {
		panic(err)
	}

	balanceOpts := []soter.SoterOpts{
		soter.UserAddressOpts(userAddress),
		soter.RawDataOpts(rawData),
		soter.SignatureOpts(signature),
	}
	out, err := sh.Balance(context.Background(), balanceOpts...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}
