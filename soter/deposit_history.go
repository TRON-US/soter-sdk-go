package soter

import (
	"context"

	"github.com/TRON-US/soter-sdk-go/utils"
)

type depositRawData struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
	Type      int32 `json:"type"`
	Timestamp int64 `json:"timestamp"`
}

func getDepositRawData(start, end int64, offset, limit int32) (string, error) {
	rawData := depositRawData{
		StartDate: start,
		EndDate: end,
		Offset: offset,
		Limit: limit,
		Type: 0,
		Timestamp: utils.GetUnixTimeNow(),
	}
	return utils.GetStructRawString(rawData)
}

func (s *Shell) QueryDepositHistory(start, end int64, offset, limit int32) (SoterResponse, error) {
	rawData, err := getDepositRawData(start, end, offset, limit)
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
	rb := s.Request("history")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err = rb.Exec(context.Background(), &out)
	return out, err
}
