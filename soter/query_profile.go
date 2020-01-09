package soter

import (
	"context"

	"github.com/TRON-US/soter-sdk-golang/utils"
)

type profileRawData struct {
	UserAddress string `json:"user_address"`
	Timestamp   int64  `json:"timestamp"`
}

func GetProfileRawData(userAddress string) (string, error) {
	reqRaw := profileRawData{
		UserAddress: userAddress,
		Timestamp:   utils.GetUnixTimeNow(),
	}

	return utils.GetStructRawString(reqRaw)
}

func (s *Shell) QueryProfile(ctx context.Context, options ...SoterOpts) (SoterResponse, error) {
	var out SoterResponse
	rb := s.Request("get_profile")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err := rb.Exec(ctx, &out)
	return out, err
}
