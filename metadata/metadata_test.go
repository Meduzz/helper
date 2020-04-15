package metadata

import (
	"encoding/json"
	"testing"
)

func TestHowItLooks(t *testing.T) {
	m := NewMetadata(
		IDRef("full"),
		"example.com",
		"/test",
		"/sub",
		"/unsub",
		Policies(
			NewPolicy("a", POST("/a"), GET("/a")),
			NewPolicy("b", GET("/b"), POST("/b")),
		),
		Roles(
			NewRole("user", "a", "b"),
			NewRole("admin", "a", "b"),
		),
	)

	bs, _ := json.Marshal(m)

	println(string(bs))
}
