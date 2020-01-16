# soter-sdk-go
> This is an unofficial go interface to soter HTTP API.

## Install
```bash
go get -u github.com/TRON-US/soter-sdk-go
```
## Usage
Soter provides a list of HTTP endpoints/services for users to store files in BTFS, but it could be
kind of complex if a user is not familiar with Soter. This soter-sdk-go provides users with a handy 
way to interact with Soter.

### Example
#### Add a file
Add a file named as "hello.txt"
```go
package main

import (
	"fmt"

	"github.com/TRON-US/soter-sdk-go/soter"
)

func main() {
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
```
#### Query user balance
```go
package main

import (
	"fmt"

	"github.com/TRON-US/soter-sdk-go/soter"
)

func main() {
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
```
#### Set autopay subscription
An example about how to set autopay subscription can be checked 
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/set_autopay_test.go).
#### Update user information
An example about how to update user information can be checked 
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/edit_profile_test.go).
#### Query user deposit history
An example about how to query user deposit history can be checked
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/deposit_history_test.go).
#### Query user order list
An example about how to query user order list can be checked 
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/order_list_test.go).
#### Query user uploaded files
An example about how to query user uploaded files can be checked
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/file_list_test.go).
#### Query order details
An example about how to query order details can be checked
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/order_details_test.go).
#### Query user profile
An example about how to query user profile can be checked 
[here](https://github.com/TRON-US/soter-sdk-go/blob/master/soter/query_profile_test.go).

## License
MIT