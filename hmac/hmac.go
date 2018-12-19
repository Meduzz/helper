package hmac

import (
	"crypto/hmac"
	"hash"
)

func Sign(key, body []byte, provider func() hash.Hash) []byte {
	hasher := hmac.New(provider, key)
	hasher.Write(body)
	return hasher.Sum(nil)
}

func Verify(key, body, signature []byte, provider func() hash.Hash) bool {
	self := Sign(key, body, provider)
	return hmac.Equal(self, signature)
}
