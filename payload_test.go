package solidgate

import (
	"encoding/json"
	"net"
	"net/mail"
	"net/url"
	"testing"
)

func TestNewStatusRequestPayload(t *testing.T) {
	p := newStatusRequestPayload(NewStatusRequest("123443334"))

	j, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	expected := `{"order_id":"123443334"}`
	actual := string(j)

	if expected != actual {
		t.Error("status request bad format")
	}
}

func TestNewRefundRequestPayload(t *testing.T) {
	p := newRefundRequestPayload(NewRefundRequest("777", 2575))

	j, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	expected := `{"order_id":"777","amount":2575}`
	actual := string(j)

	if expected != actual {
		t.Error("refund request bad format")
	}
}

func TestNewInitPaymentRequestPayload(t *testing.T) {
	ipAddress := net.ParseIP("8.8.8.8")
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewInitPaymentRequest(
		2575,
		"USD",
		customerEmail,
		"GBR",
		&ipAddress,
		"Premium package",
		"777",
		"WEB",
	)

	failURL, err := url.Parse("http://merchant.example/fail")

	if err != nil {
		t.Error(err)
	}

	successURL, err := url.Parse("http://merchant.example/success")

	if err != nil {
		t.Error(err)
	}

	callbackURL, err := url.Parse("http://merchant.example/callback")

	if err != nil {
		t.Error(err)
	}

	chargebackNotificationURL, err := url.Parse("http://merchant.example/chargeback")

	if err != nil {
		t.Error(err)
	}

	r.SetFailURL(failURL).
		SetSuccessURL(successURL).
		SetCallbackURL(callbackURL).
		SetChargebackNotificationURL(chargebackNotificationURL)

	p := newInitPaymentRequestPayload(r)

	j, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	expected := `{"amount":2575,"currency":"USD","order_id":"777","order_description":"Premium package","customer_email":"\u003cexample.user@example-email.com\u003e","geo_country":"GBR","ip_address":"8.8.8.8","platform":"WEB","fail_url":"http://merchant.example/fail","success_url":"http://merchant.example/success","callback_url":"http://merchant.example/callback","chargeback_notification_url":"http://merchant.example/chargeback"}`
	actual := string(j)

	if expected != actual {
		t.Error("init payment request bad format")
	}
}

func TestNewChargeRequestPayload(t *testing.T) {
	ipAddress := net.ParseIP("8.8.8.8")
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewChargeRequest(
		2575,
		"USD",
		"123",
		12,
		2019,
		"JOHN SNOW",
		4111111111111111,
		customerEmail,
		"GBR",
		&ipAddress,
		"Premium package",
		"777",
		"WEB",
	)

	statusURL, err := url.Parse("http://merchant.example/status")

	if err != nil {
		t.Error(err)
	}

	callbackURL, err := url.Parse("http://merchant.example/callback")

	if err != nil {
		t.Error(err)
	}

	chargebackNotificationURL, err := url.Parse("http://merchant.example/chargeback")

	if err != nil {
		t.Error(err)
	}

	r.SetCallbackURL(callbackURL).
		SetChargebackNotificationURL(chargebackNotificationURL).
		SetStatusURL(statusURL)

	if statusURL != r.statusURL {
		t.Fatalf("statusURL not equal")
	}

	if callbackURL != r.callbackURL {
		t.Fatalf("callbackURL not equal")
	}

	if chargebackNotificationURL != r.chargebackNotificationURL {
		t.Fatalf("chargebackNotificationURL not equal")
	}

	p := newChargeRequestPayload(r)

	j, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	expected := `{"amount":2575,"currency":"USD","order_id":"777","order_description":"Premium package","customer_email":"\u003cexample.user@example-email.com\u003e","geo_country":"GBR","ip_address":"8.8.8.8","platform":"WEB","status_url":"http://merchant.example/status","callback_url":"http://merchant.example/callback","chargeback_notification_url":"http://merchant.example/chargeback","card_cvv":"123","card_exp_month":12,"card_exp_year":2019,"card_holder":"JOHN SNOW","card_number":4111111111111111}`
	actual := string(j)

	if expected != actual {
		t.Error("init payment request bad format")
	}
}

func TestNewRecurringRequestPayload(t *testing.T) {
	ipAddress := net.ParseIP("8.8.8.8")
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewRecurringRequest(
		2575,
		"USD",
		"777",
		customerEmail,
		&ipAddress,
		"Premium package",
		"777",
		"WEB",
	)

	statusURL, err := url.Parse("http://merchant.example/status")

	if err != nil {
		t.Error(err)
	}

	callbackURL, err := url.Parse("http://merchant.example/callback")

	if err != nil {
		t.Error(err)
	}

	chargebackNotificationURL, err := url.Parse("http://merchant.example/chargeback")

	if err != nil {
		t.Error(err)
	}

	r.SetCallbackURL(callbackURL).
		SetChargebackNotificationURL(chargebackNotificationURL).
		SetStatusURL(statusURL)

	p := newRecurringRequestPayload(r)

	j, err := json.Marshal(p)

	if err != nil {
		t.Error(err)
	}

	expected := `{"amount":2575,"currency":"USD","order_id":"777","order_description":"Premium package","customer_email":"\u003cexample.user@example-email.com\u003e","ip_address":"8.8.8.8","platform":"WEB","status_url":"http://merchant.example/status","callback_url":"http://merchant.example/callback","chargeback_notification_url":"http://merchant.example/chargeback"}`
	actual := string(j)

	if expected != actual {
		t.Error("init payment request bad format")
	}
}
