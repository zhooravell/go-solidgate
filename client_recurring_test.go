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

func createRecurringRequest(t *testing.T) *RecurringRequest {
	ip := net.ParseIP("8.8.8.8")
	email, err := mail.ParseAddress("example.user@example-email.com")

	if err != nil {
		t.Error(err)
	}

	return NewRecurringRequest(2575, "USD", "7ats8da7sd8", email, &ip, "Premium package", "123", "WEB")
}

func TestSolidGateClient_Recurring(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/api/v1/recurring" != r.URL.String() {
			t.Fail()
		}

		if r.Header.Get("Merchant") != testMerchantID || r.Header.Get("Signature") == "" || r.Header.Get("Content-Type") != "application/json" {
			t.Fail()
		}

		payload := `{
  "transactions": {
    "1495123020887591dc450088f1": {
      "id": "1495123020887591dc450088f1",
      "operation": "recurring",
      "status": "created",
      "amount": 2575,
      "currency": "USD",
      "card": {
        "bank": "STATE BANK",
        "bin": 444455,
        "brand": "VISA",
        "country": "USA",
        "number": "444455XXXXXX6666",
        "card_exp_month": "03",
        "card_exp_year": 2025,
        "card_type": "DEBIT"
      },
      "card_token": {
        "token": "8b5ddb087c38496bc23fb5a1ce82c9640e1b8128aa9fe5bf54574d10c8a0493c71a1fc4da81e3782397fdc7c142ccdad52c0"
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
    "status": "processing",
    "total_fee_amount": 0,
    "refunded_amount": 0
  },
  "transaction": {
    "id": "1495123020887591dc450088f1",
    "operation": "recurring",
    "status": "created",
    "amount": 100,
    "currency": "USD",
    "card": {
      "bin": 344940,
      "bank": "NATIONWIDE BUILDING SOCIETY",
      "brand": "AMEX",
      "country": "USA",
      "number": "344940XXXXXX1496",
      "card_exp_month": "03",
      "card_exp_year": 2025,
      "card_type": "DEBIT"
    },
    "card_token": {
      "token": "8c6ecf6a69fc58f45ead35ed7c2a8f6edde8d7eb20e82263cba3702fcfb8faee8c1684ab910f8bb0ecb69d98386a998bf2fd"
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

	res, err := createClient(server.URL, t).Recurring(context.Background(), createRecurringRequest(t))

	if err != nil {
		t.Error(err)
	}

	if res.Order.Status != "processing" || res.PaymentAdviser.Advise != "pay" {
		t.Fail()
	}
}

func TestSolidGateClient_RecurringGateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"error":{"code":"2.01","messages":{"currency":["This value should not be blank."]}}}`)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).Recurring(context.Background(), createRecurringRequest(t))

	if err == nil || res != nil || err.Error() != "(2.01) currency: This value should not be blank." {
		t.Fail()
	}
}
