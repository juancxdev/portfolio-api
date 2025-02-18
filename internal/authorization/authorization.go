package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Authorize(sign, secret, body, path, method string) bool {
	signerRequest := ""
	if method == "POST" || method == "PUT" {
		signerRequest = Signer(body, secret)
		return sign == signerRequest
	}

	signerRequest = Signer(path, secret)
	return sign == signerRequest
}

func Signer(data, secret string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))
	// Write Data to it
	h.Write([]byte(data))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
