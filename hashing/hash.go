package hashing

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func Hash(data []byte, hasher hash.Hash) string {
	hasher.Write(data)
	bs := hasher.Sum(nil)

	return hex.EncodeToString(bs)
}

func HashWithSalt(data, salt []byte, hasher hash.Hash) string {
	hasher.Write(data)
	hasher.Write(salt)
	bs := hasher.Sum(nil)

	return hex.EncodeToString(bs)
}

func Token() string {
	randomBytes, err := Random(1024)

	if err != nil {
		return Token()
	}

	return Hash(randomBytes, sha1.New())
}

func Secret() string {
	randomBytes, err := Random(1024)

	if err != nil {
		return Secret()
	}

	return Hash(randomBytes, sha256.New())
}

func Random(size int) ([]byte, error) {
	bs := make([]byte, size)
	_, err := rand.Read(bs)

	if err != nil {
		return nil, err
	}

	return bs, nil
}
