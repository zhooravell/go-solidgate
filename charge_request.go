package solidgate

import (
	"net"
	"net/mail"
	"net/url"
)

// ChargeRequest structure to represent charge request data
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4817134/Charge+transaction
type ChargeRequest struct {
	amount                    int
	currency                  string
	orderID                   string
	orderDescription          string
	cardCvv                   string
	cardExpMonth              string
	cardExpYear               string
	cardHolder                string
	cardNumber                string
	customerEmail             *mail.Address
	geoCountry                string
	ipAddress                 *net.IP
	platform                  string
	statusURL                 *url.URL
	callbackURL               *url.URL
	chargebackNotificationURL *url.URL
}

// SetChargebackNotificationURL to set url
// Set URL of merchant page, which a customer will be redirected in case successful payment
func (rcv *ChargeRequest) SetChargebackNotificationURL(url *url.URL) *ChargeRequest {
	rcv.chargebackNotificationURL = url

	return rcv
}

// SetCallbackURL to set url
// Set URL of merchant page, where response with payment result will be sent
func (rcv *ChargeRequest) SetCallbackURL(url *url.URL) *ChargeRequest {
	rcv.callbackURL = url

	return rcv
}

// SetStatusURL to set url
// URL of merchant page, which a customer will be redirected in case successful payment
func (rcv *ChargeRequest) SetStatusURL(url *url.URL) *ChargeRequest {
	rcv.statusURL = url

	return rcv
}

// NewChargeRequest return ChargeRequest structure with mandatory parameters
func NewChargeRequest(
	amount int,
	currency string,
	cardCvv string,
	cardExpMonth string,
	cardExpYear string,
	cardHolder string,
	cardNumber string,
	customerEmail *mail.Address,
	geoCountry string,
	ipAddress *net.IP,
	orderDescription string,
	orderID string,
	platform string,
) *ChargeRequest {
	return &ChargeRequest{
		amount:           amount,
		currency:         currency,
		cardCvv:          cardCvv,
		cardExpMonth:     cardExpMonth,
		cardExpYear:      cardExpYear,
		cardHolder:       cardHolder,
		cardNumber:       cardNumber,
		customerEmail:    customerEmail,
		geoCountry:       geoCountry,
		ipAddress:        ipAddress,
		orderDescription: orderDescription,
		orderID:          orderID,
		platform:         platform,
	}
}
