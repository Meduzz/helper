package jwtutil

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// ReadClaim will allow you to read claims from JWT without validating the token. THIS IS BAD PRACTICE, but sometimes useful.
func ReadClaim(jwt, key string) (interface{}, error) {
	split := strings.Split(jwt, ".")

	if len(split) != 3 {
		return nil, fmt.Errorf("JWT does not contain 3 pieces")
	}

	bs, err := base64.RawStdEncoding.DecodeString(split[1])

	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(bs))
	found := false

	for {
		token, err := dec.Token()

		if err != nil {
			return nil, err
		}

		if found {
			return token, nil
		}

		val, ok := token.(string)

		if ok && val == key {
			found = true
		}
	}

	return nil, fmt.Errorf("Key %s was not found", key)
}
