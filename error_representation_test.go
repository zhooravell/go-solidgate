package solidgate

import "testing"

func TestGateError_hasErrorTrue(t *testing.T) {
	e := gateError{}
	e.Error.Code = "2.01"
	e.Error.Messages = make(map[string][]string)
	e.Error.Messages["currency"] = make([]string, 1)
	e.Error.Messages["currency"][0] = "This value should not be blank."

	if !e.hasError() {
		t.Fail()
	}
}

func TestGateError_hasErrorFalse(t *testing.T) {
	e := gateError{}

	if e.hasError() {
		t.Fail()
	}
}
