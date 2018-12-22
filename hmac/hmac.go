package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func Sign(key, body []byte, provider func() hash.Hash) []byte {
	hasher := hmac.New(provider, key)
	hasher.Write(body)
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}

func Verify(key, body, signature []byte, provider func() hash.Hash) bool {
	self := Sign(key, body, provider)
	return hmac.Equal(self, signature)
}
