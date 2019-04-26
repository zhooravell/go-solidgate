package solidgate

type gateError struct {
	Error errorBlock `json:"error"`
}

type errorBlock struct {
	Code     string              `json:"code"`
	Messages map[string][]string `json:"messages"`
}

// check is error structure is empty
func (rcv *gateError) hasError() bool {
	return rcv.Error.Code != "" && len(rcv.Error.Messages) != 0
}
