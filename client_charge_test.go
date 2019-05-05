package solidgate

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"testing"
)

func createChargeRequest(t *testing.T) *ChargeRequest {
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

	return NewChargeRequest(amount, currency, cardCvv, cardExpMonth, cardExpYear, cardHolder, cardNumber, customerEmail, geoCountry, &ipAddress, orderDescription, orderID, platform)
}

func TestSolidGateClient_Charge(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/api/v1/charge" != r.URL.String() {
			t.Fail()
		}

		if r.Header.Get("Merchant") != testMerchantID || r.Header.Get("Signature") == "" || r.Header.Get("Content-Type") != "application/json" {
			t.Fail()
		}

		payload := `{
  "transactions": {
    "00016857481e16b07fc": {
      "id": "00016857481e16b07fc",
      "operation": "pay",
      "status": "created",
      "amount": 2575,
      "currency": "USD",
      "card": {
        "bank": "STATE BANK",
        "bin": "411111",
        "brand": "VISA",
        "country": "USA",
        "number": "411111XXXXXX1111",
        "card_exp_month": "03",
        "card_exp_year": 2025,
        "card_type": "DEBIT"
      }
    }
  },
  "order": {
    "order_id": "777",
    "amount": 2575,
    "currency": "USD",
    "fraudulent": true,
    "marketing_amount": 2575,
    "marketing_currency": "USD",
    "status": "created",
    "refunded_amount": 0,
    "total_fee_amount": 0
  },
  "transaction": {
    "id": "00016857481e16b07fc",
    "operation": "pay",
    "status": "created",
    "amount": 2575,
    "currency": "USD",
    "card": {
      "bin": "411111",
      "brand": "VISA",
      "country": "USA",
      "number": "411111XXXXXX1111",
      "card_exp_month": "03",
      "card_exp_year": 2025,
      "card_type": "DEBIT"
    }
  },
  "payment_adviser": {
    "advise": "pay"
  }
}`

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(payload)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).Charge(context.Background(), createChargeRequest(t))

	if err != nil {
		t.Error(err)
	}

	if res.Order.Status != "created" || res.Transaction.ID != "00016857481e16b07fc" {
		t.Fail()
	}
}

func TestSolidGateClient_ChargeGateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"error":{"code":"2.01","messages":{"currency":["This value should not be blank."]}}}`)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).Charge(context.Background(), createChargeRequest(t))

	if err == nil || res != nil || err.Error() != "(2.01) currency: This value should not be blank." {
		t.Fail()
	}
}
