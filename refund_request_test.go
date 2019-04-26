package solidgate

import "testing"

func TestNewRefundRequest(t *testing.T) {
	amount := 2575
	orderID := "777"

	r := NewRefundRequest(orderID, amount)

	if amount != r.amount {
		t.Fatalf("amount not equal")
	}

	if orderID != r.orderID {
		t.Fatalf("orderID not equal")
	}
}
