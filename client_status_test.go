package solidgate

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSolidGateClient_Status(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "/api/v1/status" != r.URL.String() {
			t.Fail()
		}

		if r.Header.Get("Merchant") != testMerchantID || r.Header.Get("Signature") == "" || r.Header.Get("Content-Type") != "application/json" {
			t.Fail()
		}

		payload := `{
  "transactions": {
    "1499079426595a2302db936": {
      "id": "1499079426595a2302db936",
      "operation": "pay",
      "status": "success",
      "descriptor": "descriptor",
      "amount": 10000,
      "currency": "EUR",
      "fee": {
        "amount": 353,
        "currency": "USD"
      },
      "card": {
        "bank": "STATE BANK",
        "bin": "453245",
        "brand": "VISA",
        "country": "USA",
        "number": "453245XXXXXX2692",
        "card_exp_month": "03",
        "card_exp_year": 2025,
        "card_type": "DEBIT",
        "card_token": {
          "token": "64edcaea1441e2bde37b8afd0d0160dbd8d9db0abedf685af7b4fd0e529c8caf19abf3acd0b9b7e0bb884c85d7a21c0926a4"
        }
      }
    },
    "1499079426595b4ade6e24b": {
      "id": "1499079426595b4ade6e24b",
      "operation": "refund",
      "status": "success",
      "descriptor": "descriptor",
      "amount": 10000,
      "currency": "EUR",
      "fee": {
        "amount": 0,
        "currency": "USD"
      }
    }
  },
  "chargebacks": {
    "2": {
      "id": 2,
      "dispute_date": "2017-07-04",
      "settlement_date": "2017-07-04",
      "amount": 1000,
      "currency": "EUR",
      "reason_code": "222",
      "status": "approved"
    }
  },
  "order": {
    "order_id": "1499079426",
    "status": "refunded",
    "amount": 10000,
    "refunded_amount": 10000,
    "currency": "EUR",
    "marketing_amount": 11415,
    "marketing_currency": "USD",
    "processing_amount": 10000,
    "processing_currency": "EUR",
    "descriptor": "qwerty",
    "fraudulent": true,
    "total_fee_amount": 353,
    "fee_currency": "USD"
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

	res, err := createClient(server.URL, t).Status(context.Background(), NewStatusRequest("777"))

	if err != nil {
		t.Error(err)
	}

	if res.Order.Status != "refunded" {
		t.Fail()
	}

	if len(res.Transactions) != 2 {
		t.Fail()
	}
}
