package solidgate

import (
	"net"
	"net/mail"
	"net/url"
)

// RecurringRequest structure to represent recurring request data
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4686126/Recurring+transaction
type RecurringRequest struct {
	amount                    int
	currency                  string
	orderID                   string
	orderDescription          string
	recurringToken            string
	customerEmail             *mail.Address
	ipAddress                 *net.IP
	platform                  string
	statusURL                 *url.URL
	callbackURL               *url.URL
	chargebackNotificationURL *url.URL
}

// SetChargebackNotificationURL to set url
// Set URL of merchant page, which a customer will be redirected in case successful payment
func (rcv *RecurringRequest) SetChargebackNotificationURL(url *url.URL) *RecurringRequest {
	rcv.chargebackNotificationURL = url

	return rcv
}

// SetCallbackURL to set url
// Set URL of merchant page, where response with payment result will be sent
func (rcv *RecurringRequest) SetCallbackURL(url *url.URL) *RecurringRequest {
	rcv.callbackURL = url

	return rcv
}

// SetStatusURL to set url
// URL of merchant page, which a customer will be redirected in case successful payment
func (rcv *RecurringRequest) SetStatusURL(url *url.URL) *RecurringRequest {
	rcv.statusURL = url

	return rcv
}

// NewRecurringRequest return RecurringRequest with mandatory parameters
func NewRecurringRequest(
	amount int,
	currency string,
	recurringToken string,
	customerEmail *mail.Address,
	ipAddress *net.IP,
	orderDescription string,
	orderID string,
	platform string,
) *RecurringRequest {
	return &RecurringRequest{
		amount:           amount,
		currency:         currency,
		recurringToken:   recurringToken,
		customerEmail:    customerEmail,
		ipAddress:        ipAddress,
		orderDescription: orderDescription,
		orderID:          orderID,
		platform:         platform,
	}
}
