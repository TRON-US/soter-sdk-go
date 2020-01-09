package soter

import (
	"context"
	"encoding/json"
	"github.com/TRON-US/soter-sdk-golang/utils"
)

type balanceRawData struct {
	UserAddress string `json:"user_address"`
	Timestamp   int64  `json:"timestamp"`
}

func GetBalanceRawData(userAddress string) (string, error) {
	reqRaw := balanceRawData{
		UserAddress: userAddress,
		Timestamp:   utils.GetUnixTimeNow(),
	}

	rawBytes, err := json.Marshal(reqRaw)
	if err != nil {
		return "", nil
	}
	return string(rawBytes), nil
}

func (s *Shell) Balance(ctx context.Context, options ...SoterOpts) (SoterResponse, error) {
	var out SoterResponse
	rb := s.Request("balance")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err := rb.Exec(ctx, &out)
	return out, err
}
