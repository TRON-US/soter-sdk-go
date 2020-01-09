package soter

import (
	"bytes"
	"context"
	"github.com/TRON-US/soter-sdk-golang/utils"
	"io"
	"mime/multipart"
	"os"

	"github.com/satori/go.uuid"
)

type AddFileRawData struct {
	RequestUser string `json:"request_user"`
	SignedUser string `json:"signed_user"`
	RequestId string `json:"request_id"`
	Timestamp int64 `json:"timestamp"`
}

func (s *Shell) AddFile(ctx context.Context, requestUser, signedUser, privateKey, filePath string) (SoterResponse, error) {
	var out SoterResponse
	rb := s.Request("add")
	rb.SetMethod("POST")

	// prepare add file raw data
	rawData := AddFileRawData{
		RequestUser: requestUser,
		SignedUser: signedUser,
		RequestId: uuid.NewV4().String(),
		Timestamp: utils.GetUnixTimeNow(),
	}
	rawString, err := utils.GetStructRawString(rawData)
	if err != nil {
		return SoterResponse{}, err
	}
	signature, err := utils.GetSignature(rawString, privateKey)
	if err != nil {
		return SoterResponse{}, nil
	}

	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// set this following header is crucial
	rb = rb.Header("Content-Type", w.FormDataContentType())

	// fill the form
	err = w.WriteField("raw_data", rawString)
	if err != nil {
		return SoterResponse{}, err
	}
	err = w.WriteField("signature", signature)
	if err != nil {
		return SoterResponse{}, err
	}
	// Add your  file
	f, err := os.Open(filePath)
	if err != nil {
		return SoterResponse{}, err
	}
	fw, err := w.CreateFormFile("file", filePath)
	if err != nil {
		return SoterResponse{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return SoterResponse{}, err
	}
	_ = f.Close()
	_ = w.Close()

	// execute the request
	err = rb.Body(&b).Exec(context.Background(), &out)
	return out, err
}
