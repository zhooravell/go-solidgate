package solidgate

import (
	"net"
	"net/mail"
	"net/url"
	"testing"
)

func TestNewRecurringRequest(t *testing.T) {
	amount := 2575
	currency := "USD"
	orderID := "777"
	orderDescription := "Premium package"
	ipAddress := net.ParseIP("8.8.8.8")
	platform := "WEB"
	recurringToken := "7ats8da7sd8-a66dfa7-a9s9das89t"
	customerEmail, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	r := NewRecurringRequest(
		amount,
		currency,
		recurringToken,
		customerEmail,
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

	if recurringToken != r.recurringToken {
		t.Fatalf("recurringToken not equal")
	}

	if orderID != r.orderID {
		t.Fatalf("orderID not equal")
	}

	if orderDescription != r.orderDescription {
		t.Fatalf("orderDescription not equal")
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

func TestRecurringRequest_SetURLs(t *testing.T) {
	r := RecurringRequest{}

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
