package solidgate

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const version = "0.0.2"

// https://solidgate.atlassian.net/wiki/spaces/API/pages/4718729/Host+to+host+API
type Client interface {
	InitPayment(ctx context.Context, r *InitPaymentRequest) (*InitPaymentResponse, error)
	Status(ctx context.Context, r *StatusRequest) (*StatusResponse, error)
	Recurring(ctx context.Context, r *RecurringRequest) (*RecurringResponse, error)
	Charge(ctx context.Context, r *ChargeRequest) (*ChargeResponse, error)
	Refund(ctx context.Context, r *RefundRequest) (*RefundResponse, error)
}

// SolidGate HTTP client
type solidGateClient struct {
	merchantID string
	httpClient *http.Client
	signer     Signer
	baseURL    *url.URL
}

// Operation of order initiating. It shall be used while working with payment form.
// This operation deems to create order in system and return token being applied while calling payment form.
func (rcv *solidGateClient) InitPayment(ctx context.Context, r *InitPaymentRequest) (*InitPaymentResponse, error) {
	res, err := rcv.request(ctx, "/api/v1/init-payment", newInitPaymentRequestPayload(r))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("init payment not OK response")
	}

	body, err := extractBody(res)

	if err != nil {
		return nil, err
	}

	if err := extractErrorFromResponse(body); err != nil {
		return nil, err
	}

	result := new(InitPaymentResponse)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Request for getting current order status
func (rcv *solidGateClient) Status(ctx context.Context, r *StatusRequest) (*StatusResponse, error) {
	res, err := rcv.request(ctx, "/api/v1/status", newStatusRequestPayload(r))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status not OK response")
	}

	body, err := extractBody(res)

	if err != nil {
		return nil, err
	}

	result := new(StatusResponse)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Operation of recurring payment.
// In contrast to Charge, token previously received is to be sent in request instead of cardholder data.
// Upon successful payment, two scenarios on this order are possible - refund or chargeback.
func (rcv *solidGateClient) Recurring(ctx context.Context, r *RecurringRequest) (*RecurringResponse, error) {
	res, err := rcv.request(ctx, "/api/v1/recurring", newRecurringRequestPayload(r))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("recurring not OK response")
	}

	body, err := extractBody(res)

	if err != nil {
		return nil, err
	}

	if err := extractErrorFromResponse(body); err != nil {
		return nil, err
	}

	result := new(RecurringResponse)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Charge request - basic operation of withdrawal amounts from cardholder's account.
// Upon successful withdrawal, two scenarios on this order are possible - refund or chargeback.
// This operation can be made via 3DS. When operation successfully done,
// cardholder data are tokenized so that the subsequent payments can be effected by token (recurring payments).
func (rcv *solidGateClient) Charge(ctx context.Context, r *ChargeRequest) (*ChargeResponse, error) {
	res, err := rcv.request(ctx, "/api/v1/charge", newChargeRequestPayload(r))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("charge not OK response")
	}

	body, err := extractBody(res)

	if err != nil {
		return nil, err
	}

	if err := extractErrorFromResponse(body); err != nil {
		return nil, err
	}

	result := new(ChargeResponse)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// This is a request for transferring funds back to cardholder.
// Refunds can made only on successfully paid order.
func (rcv *solidGateClient) Refund(ctx context.Context, r *RefundRequest) (*RefundResponse, error) {
	res, err := rcv.request(ctx, "/api/v1/refund", newRefundRequestPayload(r))

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("refund not OK response")
	}

	body, err := extractBody(res)

	if err != nil {
		return nil, err
	}

	if err := extractErrorFromResponse(body); err != nil {
		return nil, err
	}

	result := new(RefundResponse)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Prepare request
func (rcv *solidGateClient) request(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	u := rcv.baseURL.ResolveReference(&url.URL{Path: path})
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u.String(), buf)

	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("SolidGateGoLangClient/%s", version)) // For detecting request made by this client
	req.Header.Set("Merchant", rcv.merchantID)

	sing, err := rcv.signer.Sine(buf.Bytes())

	if err != nil {
		return nil, err
	}

	req.Header.Set("Signature", string(sing))

	res, err := rcv.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// extract errors
func extractErrorFromResponse(body []byte) error {
	var data gateError

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	if data.hasError() {
		for k, m := range data.Error.Messages {
			for _, v := range m {
				return fmt.Errorf("(%s) %s: %s", data.Error.Code, k, v)
			}
		}
	}

	return nil
}

// extract body content
func extractBody(res *http.Response) ([]byte, error) {
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	return ioutil.ReadAll(res.Body)
}

// Constructor for solidGateClient
func NewSolidGateClient(merchantID string, httpClient *http.Client, signer Signer, baseURL *url.URL) Client {
	return &solidGateClient{
		merchantID: merchantID,
		httpClient: httpClient,
		signer:     signer,
		baseURL:    baseURL,
	}
}
