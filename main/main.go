package main

import (
	"context"
	"fmt"

	"github.com/TRON-US/soter-sdk-golang/soter"
)

func main() {
	TestSetAutopay()
	//TestBalance()
	//TestAddFile()
}

func TestAddFile()  {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := soter.NewShell(privateKey, userAddress, url)

	out, err := sh.AddFile(context.Background(), userAddress, userAddress, privateKey, "go.mod")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}


func TestSetAutopay() {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"
	sh := soter.NewShell(privateKey, userAddress, url)

	out, err := sh.Autopay(context.Background(), false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}

func TestBalance() {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := soter.NewShell(privateKey, userAddress, url)

	out, err := sh.Balance(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}
