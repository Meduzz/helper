package jwtutil

import (
	"io"
	"strings"
	"testing"
)

var token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2MDI0MDUyNDgsImV4cCI6MTYzMzk0MTI0OCwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkVtYWlsIjoidGVzdGVyQGV4YW1wbGUuY29tIn0.Z-1SqW38WrtYtxyMRq4UC7sa9c7mRcOXzSmBUMYrzlM"

func TestExtractValue(t *testing.T) {
	val, err := ReadClaim(token, "Email")

	if err != nil {
		t.Fail()
		return
	}

	email, ok := val.(string)

	if !ok {
		t.Fail()
		return
	}

	if email != "tester@example.com" {
		t.Fail()
		return
	}
}

func TestNothingFound(t *testing.T) {
	_, err := ReadClaim(token, "test")

	if err != io.EOF {
		t.Fail()
		return
	}
}

func TestNotAJWT(t *testing.T) {
	_, err := ReadClaim("qwerty.asdf.zxcv", "qwerty")

	if !strings.HasPrefix(err.Error(), "invalid character") { // <- json parsing error :(
		t.Fail()
		return
	}
}

func TestEvenLessJWT(t *testing.T) {
	_, err := ReadClaim("asdf", "qwerty")

	if err.Error() != "JWT does not contain 3 pieces" {
		t.Fail()
		return
	}
}
