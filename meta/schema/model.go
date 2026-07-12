package schema

import (
	"encoding/json"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Kind []string

	Schema struct {
		Type                 Kind               `json:"type"`
		Properties           map[string]*Schema `json:"properties,omitempty"`
		AdditionalProperties any                `json:"additionalProperties,omitempty"` // bool or *Schema
		Pattern              string             `json:"pattern,omitempty"`
		Items                *Schema            `json:"items,omitempty"`
		Enum                 []string           `json:"enum,omitempty"`
		Required             []string           `json:"required,omitempty"`
		Default              any                `json:"default,omitempty"`
		Minimum              int64              `json:"minimum,omitempty"`
		Maximum              int64              `json:"maximum,omitempty"`
		MultipleOf           int64              `json:"multipleOf,omitempty"`
		MinLength            int64              `json:"minLength,omitempty"`
		MaxLength            int64              `json:"maxLength,omitempty"`
		MinItems             int64              `json:"minItems,omitempty"`
		MaxItems             int64              `json:"maxItems,omitempty"`
		Id                   string             `json:"$id,omitempty"`
		Defs                 map[string]*Schema `json:"$defs,omitempty"`
		Ref                  string             `json:"$ref,omitempty"`
	}

	// SchemaHook allows you to affect schema generation for structs. Due to reflection, this must not be bound to a pointer of your struct but the defined struct ie `func (MyStruct) Enchance(...)`
	SchemaHook interface {
		// Enhance is called once per field and allow you to mutate the fields schema to your hearts delight.
		Enhance(field string, schema *Schema)
	}
)

const (
	Object  = "object"
	Number  = "number"
	Integer = "interger"
	String  = "string"
	Boolean = "boolean"
	Array   = "array"
	Null    = "null"
)

func (k *Kind) UnmarshalJSON(data []byte) error {
	// try to unmarshal as array
	var keys []string
	err := json.Unmarshal(data, &keys)

	if err != nil {
		// else try unmarshal as single field
		var key string
		err = json.Unmarshal(data, &key)

		if err != nil {
			// finally fail
			return err
		}

		*k = Of(key)
		return nil
	}

	*k = Kind(keys)

	return nil
}

func (k Kind) MarshalJSON() ([]byte, error) {
	if len(k) > 1 {
		return json.Marshal([]string(k))
	} else if len(k) == 1 {
		return json.Marshal(k[0])
	} else {
		return json.Marshal("")
	}
}

func (k Kind) Contains(other string) bool {
	if len(other) == 0 {
		return false
	}

	return slice.Contains(k, other)
}

func Merge(it Kind, others ...string) Kind {
	return append(it, others...)
}

func Join(first, second string) Kind {
	return []string{first, second}
}

func Drop(source Kind, needle string) Kind {
	return slice.Filter(source, func(i string) bool {
		return i != needle
	})
}

func Of(value string) Kind {
	return Kind([]string{value})
}
