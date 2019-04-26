package solidgate

import (
	"net/http"
	"net/url"
	"testing"
)

var (
	testMerchantID = "test"
	testPrivateKey = []byte("test")
)

func createClient(serverURL string, t *testing.T) Client {
	baseURL, err := url.Parse(serverURL)

	if err != nil {
		t.Error(err)
	}

	s := NewSha512Signer(testMerchantID, testPrivateKey)

	return NewSolidGateClient(testMerchantID, &http.Client{}, s, baseURL)
}

func TestNewSolidGateClient(t *testing.T) {
	c := createClient("https://pay.signedpay.com", t)

	if c == nil {
		t.Fail()
	}
}
