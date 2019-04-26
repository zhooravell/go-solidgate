package solidgate

import (
	"net"
	"net/mail"
	"net/url"
	"testing"
)

func TestNewInitPaymentRequest(t *testing.T) {
	amount := 2575
	currency := "USD"
	orderID := "777"
	orderDescription := "Premium package"
	geoCountry := "GBR"
	ipAddress := net.ParseIP("8.8.8.8")
	platform := "WEB"
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewInitPaymentRequest(
		amount,
		currency,
		customerEmail,
		geoCountry,
		&ipAddress,
		orderDescription,
		orderID,
		platform,
	)

	if amount != r.amount {
		t.Fatalf("amount not equal")
	}

	if currency != r.currency {
		t.Fatalf("currency not equal")
	}

	if orderID != r.orderID {
		t.Fatalf("orderID not equal")
	}

	if orderDescription != r.orderDescription {
		t.Fatalf("orderDescription not equal")
	}

	if geoCountry != r.geoCountry {
		t.Fatalf("geoCountry not equal")
	}

	if &ipAddress != r.ipAddress {
		t.Fatalf("ipAddress not equal")
	}

	if platform != r.platform {
		t.Fatalf("platform not equal")
	}

	if customerEmail != r.customerEmail {
		t.Fatalf("customerEmail not equal")
	}

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

	if failURL != r.failURL {
		t.Fatalf("failURL not equal")
	}

	if successURL != r.successURL {
		t.Fatalf("successURL not equal")
	}

	if callbackURL != r.callbackURL {
		t.Fatalf("callbackURL not equal")
	}

	if chargebackNotificationURL != r.chargebackNotificationURL {
		t.Fatalf("chargebackNotificationURL not equal")
	}
}
