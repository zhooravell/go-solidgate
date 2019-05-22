package solidgate

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSolidGateClient_Refund(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/api/v1/refund" != r.URL.String() {
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
    "marketing_amount": 2575,
    "marketing_currency": "USD",
    "processing_amount": 2575,
    "processing_currency": "USD",
    "status": "refunded",
    "refunded_amount": 50,
    "total_fee_amount": 48,
    "fee_currency": "USD"
  },
  "transaction": {
    "id": "149310695572758ff0116380f7",
    "operation": "refund",
    "status": "success",
    "descriptor": "DESCRIPTOR",
    "amount": 2575,
    "currency": "USD",
    "fee": {
        "amount": 10,
        "currency": "USD"
      }
  },
  "transactions": {
    "149310695572758ff0116380f7": {
      "id": "149310695572758ff0116380f7",
      "operation": "refund",
      "status": "success",
      "descriptor": "DESCRIPTOR",
      "amount": 2575,
      "currency": "USD",
      "fee": {
        "amount": 10,
        "currency": "USD"
      }
    },
    "149310695572758ff010f5d352": {
      "id": "149310695572758ff010f5d352",
      "operation": "pay",
      "status": "success",
      "descriptor": "DESCRIPTOR",
      "amount": 2575,
      "currency": "USD",
      "fee": {
        "amount": 38,
        "currency": "USD"
      },
      "card": {
        "bin": 444455,
        "brand": "VISA",
        "country": "USA",
        "number": "444455XXXXXX6666",
        "card_exp_month": "03",
        "card_exp_year": 2025,
        "card_type": "DEBIT",
        "card_token": {
          "token": "140dc2f1bd02c...f806d4db669ef"
        }
      }
    }
  }
}`

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(payload)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).Refund(context.Background(), NewRefundRequest("777", 2525))

	if err != nil {
		t.Error(err)
	}

	if res.Order.Status != "refunded" || res.Transaction.ID != "149310695572758ff0116380f7" {
		t.Fail()
	}
}

func TestSolidGateClient_RefundGateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"error":{"code":"2.01","messages":{"currency":["This value should not be blank."]}}}`)); err != nil {
			log.Println(err)
		}
	}))

	defer server.Close()

	res, err := createClient(server.URL, t).Refund(context.Background(), NewRefundRequest("777", 2525))

	if err == nil || res != nil || err.Error() != "(2.01) currency: This value should not be blank." {
		t.Fail()
	}
}
