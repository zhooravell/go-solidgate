package solidgate

import (
	"net"
	"net/mail"
	"net/url"
	"testing"
)

func TestNewChargeRequest(t *testing.T) {
	amount := 2575
	currency := "USD"
	orderID := "777"
	orderDescription := "Premium package"
	geoCountry := "GBR"
	ipAddress := net.ParseIP("8.8.8.8")
	platform := "WEB"
	cardCvv := "123"
	cardExpMonth := "12"
	cardExpYear := "2019"
	cardHolder := "JOHN SNOW"
	cardNumber := "4111111111111111"
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewChargeRequest(
		amount,
		currency,
		cardCvv,
		cardExpMonth,
		cardExpYear,
		cardHolder,
		cardNumber,
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

	if cardCvv != r.cardCvv {
		t.Fatalf("cardCvv not equal")
	}

	if cardExpMonth != r.cardExpMonth {
		t.Fatalf("cardExpMonth not equal")
	}

	if cardExpYear != r.cardExpYear {
		t.Fatalf("cardExpYear not equal")
	}

	if cardHolder != r.cardHolder {
		t.Fatalf("cardHolder not equal")
	}

	if cardNumber != r.cardNumber {
		t.Fatalf("cardNumber not equal")
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
}

func TestNewChargeRequest_SetURLS(t *testing.T) {
	r := ChargeRequest{}

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
}
