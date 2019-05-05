package solidgate

// JSON body representation for refund
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4784391/Refund+transaction
type refundRequestPayload struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}

// Create refund payload from refund request
func newRefundRequestPayload(r *RefundRequest) *refundRequestPayload {
	return &refundRequestPayload{
		Amount:  r.amount,
		OrderID: r.orderID,
	}
}

// JSON body representation for status
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4849863/Check+Order+Status
type statusRequestPayload struct {
	OrderID string `json:"order_id"`
}

// Create status payload from status request
func newStatusRequestPayload(r *StatusRequest) *statusRequestPayload {
	return &statusRequestPayload{
		OrderID: r.orderID,
	}
}

// JSON body representation for init payment
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4718775/InitPayment+transaction
type initPaymentRequestPayload struct {
	Amount           int    `json:"amount"`
	Currency         string `json:"currency"`
	OrderID          string `json:"order_id"`
	OrderDescription string `json:"order_description"`
	CustomerEmail    string `json:"customer_email"`
	GeoCountry       string `json:"geo_country"`
	IPAddress        string `json:"ip_address"`
	Platform         string `json:"platform"`
	urlPayload
}

// Create init payment payload from init payment request
func newInitPaymentRequestPayload(r *InitPaymentRequest) *initPaymentRequestPayload {
	p := &initPaymentRequestPayload{
		Amount:           r.amount,
		Currency:         r.currency,
		OrderID:          r.orderID,
		OrderDescription: r.orderDescription,
		CustomerEmail:    r.customerEmail.String(),
		GeoCountry:       r.geoCountry,
		IPAddress:        r.ipAddress.String(),
		Platform:         r.platform,
	}

	if r.failURL != nil {
		p.FailURL = r.failURL.String()
	}

	if r.successURL != nil {
		p.SuccessURL = r.successURL.String()
	}

	if r.callbackURL != nil {
		p.CallbackURL = r.callbackURL.String()
	}

	if r.chargebackNotificationURL != nil {
		p.ChargeBackNotificationURL = r.chargebackNotificationURL.String()
	}

	return p
}

// JSON body representation for init charge
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4817134/Charge+transaction
type chargeRequestPayload struct {
	initPaymentRequestPayload
	CardCvv      string `json:"card_cvv"`
	CardExpMonth string `json:"card_exp_month"`
	CardExpYear  string `json:"card_exp_year"`
	CardHolder   string `json:"card_holder"`
	CardNumber   string `json:"card_number"`
}

// Create charge payload from init charge request
func newChargeRequestPayload(r *ChargeRequest) *chargeRequestPayload {
	p := &chargeRequestPayload{
		initPaymentRequestPayload: initPaymentRequestPayload{
			Amount:           r.amount,
			Currency:         r.currency,
			OrderID:          r.orderID,
			OrderDescription: r.orderDescription,
			CustomerEmail:    r.customerEmail.Address,
			GeoCountry:       r.geoCountry,
			IPAddress:        r.ipAddress.String(),
			Platform:         r.platform,
		},
		CardCvv:      r.cardCvv,
		CardExpMonth: r.cardExpMonth,
		CardExpYear:  r.cardExpYear,
		CardHolder:   r.cardHolder,
		CardNumber:   r.cardNumber,
	}

	if r.statusURL != nil {
		p.StatusURL = r.statusURL.String()
	}

	if r.callbackURL != nil {
		p.CallbackURL = r.callbackURL.String()
	}

	if r.chargebackNotificationURL != nil {
		p.ChargeBackNotificationURL = r.chargebackNotificationURL.String()
	}

	return p
}

// JSON body representation for init recurring
// https://solidgate.atlassian.net/wiki/spaces/API/pages/4686126/Recurring+transaction
type recurringRequestPayload struct {
	Amount           int    `json:"amount"`
	Currency         string `json:"currency"`
	OrderID          string `json:"order_id"`
	OrderDescription string `json:"order_description"`
	CustomerEmail    string `json:"customer_email"`
	IPAddress        string `json:"ip_address"`
	Platform         string `json:"platform"`
	urlPayload
}

// Create recurring payment payload from recurring payment request
func newRecurringRequestPayload(r *RecurringRequest) *recurringRequestPayload {
	p := &recurringRequestPayload{
		Amount:           r.amount,
		Currency:         r.currency,
		OrderID:          r.orderID,
		OrderDescription: r.orderDescription,
		CustomerEmail:    r.customerEmail.String(),
		IPAddress:        r.ipAddress.String(),
		Platform:         r.platform,
	}

	if r.statusURL != nil {
		p.StatusURL = r.statusURL.String()
	}

	if r.callbackURL != nil {
		p.CallbackURL = r.callbackURL.String()
	}

	if r.chargebackNotificationURL != nil {
		p.ChargeBackNotificationURL = r.chargebackNotificationURL.String()
	}

	return p
}

// JSON body representation for callback urls
type urlPayload struct {
	FailURL                   string `json:"fail_url,omitempty"`
	StatusURL                 string `json:"status_url,omitempty"`
	SuccessURL                string `json:"success_url,omitempty"`
	CallbackURL               string `json:"callback_url,omitempty"`
	ChargeBackNotificationURL string `json:"chargeback_notification_url,omitempty"`
}
