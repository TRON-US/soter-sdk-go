package soter

import (
	"context"
	"github.com/TRON-US/soter-sdk-go/utils"
)

type updateProfileRawData struct {
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Timestamp   int64  `json:"timestamp"`
}

type updateProfilePayload struct {
	UserAddress string               `json:"user_address"`
	RawData     updateProfileRawData `json:"raw_data"`
	Signature   string               `json:"signature"`
}

func getUpdateProfilePayload(email, phoneNumber, userAddress, privateKey string) (string, error) {
	rawData := updateProfileRawData{
		Email: email,
		PhoneNumber: phoneNumber,
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
	payload := updateProfilePayload{
		UserAddress: userAddress,
		RawData: rawData,
		Signature: signature,
	}

	return utils.GetStructRawString(payload)
}

func (s *Shell) EditProfile(email, phoneNumber string) (SoterResponse, error) {
	payload, err := getUpdateProfilePayload(email, phoneNumber, s.userAddress, s.privateKey)
	if err != nil {
		return SoterResponse{}, err
	}
	var out SoterResponse
	rb := s.Request("edit_profile")
	rb = rb.BodyString(payload)
	rb = rb.Header("Content-Type", "application/json")
	rb.SetMethod("POST")
	err = rb.Exec(context.Background(), &out)
	return out, err
}
