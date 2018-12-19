package hmac

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
	"testing"
)

func TestAHappyCase(t *testing.T) {
	bunchOfBytes := make([]byte, 1024)
	io.ReadFull(rand.Reader, bunchOfBytes)

	secretKey := []byte("this is a top secret shared key")

	signature := Sign(secretKey, bunchOfBytes, sha256.New)
	valid := Verify(secretKey, bunchOfBytes, signature, sha256.New)

	if !valid {
		t.Fail()
	}
}

func TestUnhappyCase(t *testing.T) {
	bunchOfBytes := make([]byte, 1024)
	io.ReadFull(rand.Reader, bunchOfBytes)
	secretKey := []byte("this is a top secret shared key")

	otherKey := []byte("this is another made up key")
	otherBytes := []byte("An entirerly different bunch of bytes")

	signature := Sign(secretKey, bunchOfBytes, sha256.New)
	valid := Verify(otherKey, bunchOfBytes, signature, sha256.New)

	if valid {
		t.Fail()
	}

	valid = Verify(secretKey, otherBytes, signature, sha256.New)

	if valid {
		t.Fail()
	}
}
