package utils

import (
	"encoding/hex"
	"encoding/json"
	"github.com/TRON-US/chaos/crypto"
	"time"
)

/*
	SignRawData returns the signature signed by a given private key against a given raw data
*/
func GetSignature(rawData, privateKey string) (string, error) {
	digest, err := crypto.Signature(true, rawData, privateKey)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(digest), nil
}

/*
	GetUnixTimeNow returns the unix timestamp in millisecond since 00:00:00 UTC on 1 January 1970
*/
func GetUnixTimeNow() int64 {
	return time.Now().Unix() * 1000
}

/*
	GetStructRawString returns the string representation of a marshaled struct instance
 */

func GetStructRawString(i interface{}) (string, error)  {
	rawBytes, err := json.Marshal(i)
	if err != nil {
		return "", nil
	}
	return string(rawBytes), nil
}