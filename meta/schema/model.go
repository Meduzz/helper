package schema

import "encoding/json"

type (
	Kind string

	Schema struct {
		Type       []Kind             `json:"kind"`
		Properties map[string]*Schema `json:"properties,omitempty"`
		Pattern    string             `json:"pattern,omitempty"`
		Items      *Schema            `json:"items,omitempty"`
		Enum       []string           `json:"enum,omitempty"`
		Required   []string           `json:"required,omitempty"`
		Default    any                `json:"default,omitempty"`
		Minimum    int64              `json:"minimum,omitempty"`
		Maximum    int64              `json:"maximum,omitempty"`
		MultipleOf int64              `json:"multipleOf,omitempty"`
		MinLength  int64              `json:"minLength,omitempty"`
		MaxLength  int64              `json:"maxLength,omitempty"`
		MinItems   int64              `json:"minItems,omitempty"`
		MaxItems   int64              `json:"maxItems,omitempty"`
		Id         string             `json:"$id,omitempty"`
		Defs       map[string]*Schema `json:"$defs,omitempty"`
		json.RawMessage
	}
)

const (
	Object  = Kind("object")
	Number  = Kind("number")
	Integer = Kind("interger")
	String  = Kind("string")
	Boolean = Kind("boolean")
	Array   = Kind("array")
	Null    = Kind("null")
)
