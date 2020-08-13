package hashing

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	text := "test"
	result := Hash([]byte(text), sha1.New())

	if result != "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
		fmt.Printf("result was not equal to ... %s", result)
		t.Fail()
	}
}

func TestHashWithSalt(t *testing.T) {
	text := "test"
	salt := "salt"
	result := HashWithSalt([]byte(text), []byte(salt), sha1.New())

	if result != "f438229716cab43569496f3a3630b3727524b81b" {
		fmt.Printf("result was not equal to ... %s", result)
		t.Fail()
	}
}
