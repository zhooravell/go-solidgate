package solidgate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"github.com/pkg/errors"
)

// https://solidgate.atlassian.net/wiki/spaces/API/pages/4784199/Access+to+SolidGate+Gateway+API
type Signer interface {
	Sine(payload []byte) (string, error)
}

type sha512Signer struct {
	merchantID string
	privateKey []byte
}

// Constructor for sha512Signer
func NewSha512Signer(merchantID string, privateKey []byte) Signer {
	return &sha512Signer{merchantID: merchantID, privateKey: privateKey}
}

// Value of a signature is base64-coding of hash function SHA-512
func (rcv *sha512Signer) Sine(payload []byte) (string, error) {
	// merchantId + requestJsonData + merchantId
	w := new(bytes.Buffer)
	w.WriteString(rcv.merchantID)
	w.Write(payload)
	w.WriteString(rcv.merchantID)

	h := hmac.New(sha512.New, rcv.privateKey)

	if _, err := h.Write(w.Bytes()); err != nil {
		return "", errors.Wrap(err, "sha512 sine generator")
	}
	// to lowercase hexits
	sum := hex.EncodeToString(h.Sum(nil))

	return base64.StdEncoding.EncodeToString([]byte(sum)), nil
}
