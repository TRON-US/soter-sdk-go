package soter

import (
	"context"

	"github.com/TRON-US/soter-sdk-go/utils"
)

type autopayRawData struct {
	Autopay   bool  `json:"autopay"`
	Timestamp int64 `json:"timestamp"`
}

type autopayPayload struct {
	UserAddress string         `json:"user_address"`
	RawData     autopayRawData `json:"raw_data"`
	Signature   string         `json:"signature"`
}

func getAutopayPayload(enable bool, userAddress, privateKey string) (string, error) {
	rawData := autopayRawData{
		Autopay:   enable,
		Timestamp: utils.GetUnixTimeNow(),
	}
	rawString, err := utils.GetStructRawString(rawData)
	if err != nil {
		return "", nil
	}
	signature, err := utils.GetSignature(rawString, privateKey)
	if err != nil {
		return "", nil
	}
	payload := autopayPayload{
		UserAddress: userAddress,
		RawData:     rawData,
		Signature:   signature,
	}
	return utils.GetStructRawString(payload)
}

func (s *Shell) Autopay(enable bool) (SoterResponse, error) {
	payload, err := getAutopayPayload(enable, s.userAddress, s.privateKey)
	if err != nil {
		return SoterResponse{}, err
	}
	var out SoterResponse
	rb := s.Request("autopay")
	rb = rb.BodyString(payload)
	rb = rb.Header("Content-Type", "application/json")
	rb.SetMethod("POST")
	err = rb.Exec(context.Background(), &out)
	return out, err
}
