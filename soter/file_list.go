package soter

import (
	"context"
	"github.com/TRON-US/soter-sdk-golang/utils"
)

type fileListRawData struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
	IsDeleted bool  `json:"is_deleted"`
	Timestamp int64 `json:"timestamp"`
}

func getFileListRawData(start, end int64, offset, limit int32, deleted bool) (string, error) {
	rawData := fileListRawData{
		StartDate: start,
		EndDate: end,
		Offset: offset,
		Limit: limit,
		IsDeleted: deleted,
		Timestamp: utils.GetUnixTimeNow(),
	}
	return utils.GetStructRawString(rawData)
}

func (s *Shell) QueryFileList(start, end int64, offset, limit int32, deleted bool) (SoterResponse, error) {
	rawData, err := getFileListRawData(start, end, offset, limit, deleted)
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
	rb := s.Request("files")
	for _, option := range options {
		option(rb)
	}
	rb.SetMethod("GET")
	err = rb.Exec(context.Background(), &out)
	return out, err
}