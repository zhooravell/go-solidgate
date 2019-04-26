package solidgate

import "testing"

func TestNewStatusRequest(t *testing.T) {
	id := "123443334"

	r := NewStatusRequest(id)

	if id != r.orderID {
		t.Fail()
	}
}
