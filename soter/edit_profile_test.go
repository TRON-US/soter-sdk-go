package soter

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShell_EditProfile(t *testing.T) {
	url := "https://sandbox.btfssoter.io/"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := NewShell(privateKey, userAddress, url)

	out, err := sh.EditProfile("2020-01-10-u2@gmail.com", "18890076713")
	if err != nil {
		t.Fatal(err)
	}
	if out.Code != 0 {
		t.Error("controller response code is not OK")
	}
	t.Log(fmt.Sprintf("response code: %v", out.Code))
	t.Log(fmt.Sprintf("response message: %v", out.Message))
	data, _ := json.Marshal(out.Data)
	t.Log(fmt.Sprintf("data: %v", string(data)))
}
