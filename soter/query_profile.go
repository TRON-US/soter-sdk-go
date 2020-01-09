package soter

import (
	"context"

	"github.com/TRON-US/soter-sdk-golang/utils"
)

type profileRawData struct {
	UserAddress string `json:"user_address"`
	Timestamp   int64  `json:"timestamp"`
}

func getProfileRawData(userAddress string) (string, error) {
	reqRaw := profileRawData{
		UserAddress: userAddress,
		Timestamp:   utils.GetUnixTimeNow(),
	}
	return utils.GetStructRawString(reqRaw)
}

func (s *Shell) QueryProfile() (SoterResponse, error) {
	rawData, err := getProfileRawData(s.userAddress)
	if err != nil {
		return SoterResponse{}, err
	}
	signature, err := utils.GetSignature(rawData, s.privateKey)
	if err != nil {
		return SoterResponse{}, err
	}
	options := []SoterOpts{
		UserAddressOpts(s.userAddress),
		RawDataOpts(rawData),
		SignatureOpts(signature),
	}

	var out SoterResponse
	rb := s.Request("get_profile")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err = rb.Exec(context.Background(), &out)
	return out, err
}
