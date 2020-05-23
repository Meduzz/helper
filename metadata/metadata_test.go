package metadata

import (
	"encoding/json"
	"testing"
)

func TestHowItLooks(t *testing.T) {
	m := NewMetadata(
		"example.com",
		"/test",
		"/cb",
		Policies(
			NewPolicy("a", POST("/a"), GET("/a")),
			NewPolicy("b", GET("/b"), POST("/b")),
		),
		Roles(
			NewRole("user", false, "a", "b"),
			NewRole("admin", true, "a", "b"),
		),
	)

	bs, _ := json.Marshal(m)

	println(string(bs))
}
