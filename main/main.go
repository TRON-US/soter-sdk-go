package main

import (
	"fmt"

	"github.com/TRON-US/soter-sdk-go/soter"
)

func main() {
	//TestSetAutopay()
	//TestBalance()
	//TestAddFile()
	TestOrderList()
}

func TestOrderList() {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := soter.NewShell(privateKey, userAddress, url)

	start := 1561826420000
	end := 1581826420000
	offset := 0
	limit := 100

	out, err := sh.QueryOrderList(int64(start), int64(end), int32(offset), int32(limit))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}

func TestAddFile()  {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := soter.NewShell(privateKey, userAddress, url)

	out, err := sh.AddFile(userAddress, "go.mod")
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

	out, err := sh.Autopay(false)
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

	out, err := sh.Balance()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}
