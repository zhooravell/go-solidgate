package solidgate

import (
	"net"
	"net/mail"
	"net/url"
)

// InitPaymentRequest structure to represent init payment request data
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4718775/InitPayment+transaction
type InitPaymentRequest struct {
	amount                    int
	currency                  string
	orderID                   string
	orderDescription          string
	customerEmail             *mail.Address
	geoCountry                string
	ipAddress                 *net.IP
	platform                  string
	failURL                   *url.URL
	successURL                *url.URL
	callbackURL               *url.URL
	chargebackNotificationURL *url.URL
}

// SetChargebackNotificationURL to set url
// Set URL of merchant page, which a customer will be redirected in case successful payment
func (rcv *InitPaymentRequest) SetChargebackNotificationURL(url *url.URL) *InitPaymentRequest {
	rcv.chargebackNotificationURL = url

	return rcv
}


// SetCallbackURL to set url
// Set URL of merchant page, where response with payment result will be sent
func (rcv *InitPaymentRequest) SetCallbackURL(url *url.URL) *InitPaymentRequest {
	rcv.callbackURL = url

	return rcv
}

// SetSuccessURL to set url
func (rcv *InitPaymentRequest) SetSuccessURL(url *url.URL) *InitPaymentRequest {
	rcv.successURL = url

	return rcv
}

// SetFailURL to set url
// Set URL of merchant page, which a customer will be redirected in case of not successful payment
func (rcv *InitPaymentRequest) SetFailURL(failURL *url.URL) *InitPaymentRequest {
	rcv.failURL = failURL

	return rcv
}

// NewInitPaymentRequest return InitPaymentRequest with mandatory parameters
func NewInitPaymentRequest(
	amount int,
	currency string,
	customerEmail *mail.Address,
	geoCountry string,
	ipAddress *net.IP,
	orderDescription string,
	orderID string,
	platform string,
) *InitPaymentRequest {
	return &InitPaymentRequest{
		amount:           amount,
		currency:         currency,
		customerEmail:    customerEmail,
		geoCountry:       geoCountry,
		ipAddress:        ipAddress,
		orderDescription: orderDescription,
		orderID:          orderID,
		platform:         platform,
	}
}
