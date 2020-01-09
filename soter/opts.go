package soter

type SoterOpts = func(*RequestBuilder) error

func UserAddressOpts(userAddress string) SoterOpts {
	return func(rb *RequestBuilder) error {
		rb.Option("user_address", userAddress)
		return nil
	}
}

func RawDataOpts(rawData string) SoterOpts {
	return func(rb *RequestBuilder) error {
		rb.Option("raw_data", rawData)
		return nil
	}
}

func SignatureOpts(signature string) SoterOpts {
	return func(rb *RequestBuilder) error {
		rb.Option("signature", signature)
		return nil
	}
}

// response
type SoterResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
