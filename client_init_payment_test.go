package solidgate

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"testing"
	"time"
)

func createInitRequest(t *testing.T) *InitPaymentRequest {
	ip := net.ParseIP("8.8.8.8")
	email, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	return NewInitPaymentRequest(2575, "USD", email, "GBR", &ip, "Premium package", "123", "WEB")
}

func TestSolidGateClient_InitPayment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/api/v1/init-payment" != r.URL.String() {
			t.Fail()
		}

		if r.Header.Get("Merchant") != testMerchantID || r.Header.Get("Signature") == "" || r.Header.Get("Content-Type") != "application/json" {
			t.Fail()
		}

		payload := `{
  "order": {
    "order_id": "777",
    "amount": 2575,
    "currency": "USD",
    "fraudulent": true,
    "status": "created",
    "total_fee_amount": 0
  },
  "pay_form": {
    "token": "5849cf0468afc0ac4be4963c619b776eee75271d56e3c6621ab21",
    "design_name": "promo_christmas"
  }
}`
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(payload)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).InitPayment(context.Background(), createInitRequest(t))

	if err != nil {
		t.Error(err)
	}

	if res.Order.Status != "created" || res.PayForm.Token != "5849cf0468afc0ac4be4963c619b776eee75271d56e3c6621ab21" {
		t.Fail()
	}
}

func TestSolidGateClient_InitPaymentNotOkStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).InitPayment(context.Background(), createInitRequest(t))

	if err == nil || res != nil {
		t.Fail()
	}
}

func TestSolidGateClient_InitPaymentGateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"error":{"code":"2.01","messages":{"currency":["This value should not be blank."]}}}`)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).InitPayment(context.Background(), createInitRequest(t))

	if err == nil || res != nil || err.Error() != "(2.01) currency: This value should not be blank." {
		t.Fail()
	}
}

func TestSolidGateClient_InitPaymentTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusNoContent)

	}))

	defer server.Close()

	ctx, cf := context.WithTimeout(context.Background(), 10*time.Millisecond)

	if cf == nil {
		t.Fail()
	}

	res, err := createClient(server.URL, t).InitPayment(ctx, createInitRequest(t))

	if err == nil || res != nil {
		t.Fail()
	}
}
