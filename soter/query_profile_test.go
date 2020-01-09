package soter

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestShell_QueryProfile(t *testing.T) {
	url := "http://127.0.0.1:8101"
	privateKey := "c8f0884e706c761e80a9227736a4a595f56b43660041920a5e6765a9b8ac3ab7"
	userAddress := "TTCXimHXjen9BdTFW5JvcLKGWNm3SSuECF"

	sh := NewShell(privateKey, userAddress, url)

	out, err := sh.QueryProfile(context.Background())
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
