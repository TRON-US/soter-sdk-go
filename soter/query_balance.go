package soter

import (
	"context"
	"github.com/TRON-US/soter-sdk-golang/utils"
)

type balanceRawData struct {
	UserAddress string `json:"user_address"`
	Timestamp   int64  `json:"timestamp"`
}

func getBalanceRawData(userAddress string) (string, error) {
	reqRaw := balanceRawData{
		UserAddress: userAddress,
		Timestamp:   utils.GetUnixTimeNow(),
	}

	return utils.GetStructRawString(reqRaw)
}

func (s *Shell) Balance(ctx context.Context) (SoterResponse, error) {
	rawData, err := getBalanceRawData(s.userAddress)
	if err != nil {
		return SoterResponse{}, err
	}
	signature, err := utils.GetSignature(rawData, s.privateKey)
	if err != nil {
		return SoterResponse{}, err
	}

	balanceOpts := []SoterOpts{
		UserAddressOpts(s.userAddress),
		RawDataOpts(rawData),
		SignatureOpts(signature),
	}
	
	var out SoterResponse
	rb := s.Request("balance")
	for _, option := range balanceOpts {
		option(rb)
	}
	rb.SetMethod("GET")
	err = rb.Exec(ctx, &out)
	return out, err
}
