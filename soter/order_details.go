package soter

import (
	"context"
	"github.com/TRON-US/soter-sdk-go/utils"
)

type orderDetailsRawData struct {
	RequestId string `json:"request_id"`
	Timestamp int64  `json:"timestamp"`
}

func getOrderDetailsRawData(reqId string) (string, error) {
	rawData := orderDetailsRawData{
		RequestId: reqId,
		Timestamp: utils.GetUnixTimeNow(),
	}
	return utils.GetStructRawString(rawData)
}

func (s *Shell) QueryOrderDetails(requestId string) (SoterResponse, error) {
	rawData, err := getOrderDetailsRawData(requestId)
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
	rb := s.Request("order_details")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err = rb.Exec(context.Background(), &out)
	return out, err
}