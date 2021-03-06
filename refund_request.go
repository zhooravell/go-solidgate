package solidgate

// RefundRequest structure to represent refund request data
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4784391/Refund+transaction
type RefundRequest struct {
	orderID string // Order ID specified in merchant system.
	amount  int    // Refund amount - integer without fractional component (i.e cents). For instance, 1020 (USD) means 10 USD and 20 cents.
}

// NewRefundRequest return RefundRequest with mandatory parameters
func NewRefundRequest(orderID string, amount int) *RefundRequest {
	return &RefundRequest{
		orderID: orderID,
		amount:  amount,
	}
}
