package solidgate

// StatusRequest structure to represent status request data
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4849863/Check+Order+Status
type StatusRequest struct {
	orderID string // Order ID specified in the merchant system
}

// NewStatusRequest return StatusRequest with mandatory parameters
func NewStatusRequest(orderID string) *StatusRequest {
	return &StatusRequest{
		orderID: orderID,
	}
}
