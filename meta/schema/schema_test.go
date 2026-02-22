package schema_test

import (
	"testing"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/meta/schema"
)

type (
	Person struct {
		Name  string `json:"name"`
		Age   int    `json:"omitempty"`
		Email string `json:"-"`
	}
)

var (
	_ schema.SchemaHook = Person{}
)

func TestSchema(t *testing.T) {
	t.Run("SchemaBuilder", func(t *testing.T) {
		subject := &schema.Schema{}
		builder := schema.NewSchemaBuilder(subject)

		builder.Default(1)

		if subject.Default != 1 {
			t.Errorf("default was not 1 but %v", subject.Default)
		}

		builder.Array(func(ab schema.ArrayBuilder) {})

		if !contains(subject.Type, schema.Array) {
			t.Error("type did not contain array")
		}

		builder.Boolean()

		if !contains(subject.Type, schema.Boolean) {
			t.Error("type did not contain boolean")
		}

		builder.String(func(sb schema.StringBuilder) {})

		if !contains(subject.Type, schema.String) {
			t.Error("type did not contain string")
		}

		builder.Number(func(nb schema.NumberBuilder) {})

		if !contains(subject.Type, schema.Number) {
			t.Error("type did not contain number")
		}

		builder.Object(func(ob schema.ObjectBuilder) {})

		if !contains(subject.Type, schema.Object) {
			t.Error("type did not contain object")
		}
	})

	t.Run("ObjectBuilder", func(t *testing.T) {
		subject := &schema.Schema{}
		builder := schema.NewObjectBuilder(subject)

		builder.Id("Object")

		if subject.Id != "Object" {
			t.Errorf("id was not Object but: %v", subject.Id)
		}

		builder.Property("name", true, false, func(sb schema.SchemaBuilder) {
			sb.String(func(sb schema.StringBuilder) {})
		})
		builder.Property("age", false, true, func(sb schema.SchemaBuilder) {
			sb.Number(func(nb schema.NumberBuilder) {})
		})

		if !contains(subject.Required, "name") {
			t.Error("name was not marked as required")
		}

		nameSchema, nameOk := subject.Properties["name"]
		ageSchema, ageOk := subject.Properties["age"]

		if !nameOk {
			t.Error("name waas not set")
		}

		if contains(nameSchema.Type, schema.Null) {
			t.Error("name was marked as nullable")
		}

		if !ageOk {
			t.Error("age was not set")
		}

		if !contains(ageSchema.Type, schema.Null) {
			t.Error("age was not marked as nullable")
		}
	})

	t.Run("ArrayBuidler", func(t *testing.T) {
		subject := &schema.Schema{}
		builder := schema.NewArrayBuilder(subject)

		builder.Minimum(1)
		builder.Maximum(10)
		builder.Items(func(sb schema.SchemaBuilder) {
			sb.String(func(sb schema.StringBuilder) {})
		})

		if subject.MaxItems != 10 {
			t.Error("maxItems was not set correctly")
		}

		if subject.MinItems != 1 {
			t.Error("minItems was not set correctly")
		}

		if subject.Items == nil {
			t.Error("items was not set")
		}

		if !contains(subject.Items.Type, schema.String) {
			t.Error("items.Type was not string")
		}
	})

	t.Run("NumberBuilder", func(t *testing.T) {
		subject := &schema.Schema{}
		builder := schema.NewNumberBuilder(subject)

		builder.Minimum(1)
		builder.Maximum(10)
		builder.MultipleOf(2)
		builder.Integer()

		if !contains(subject.Type, schema.Integer) {
			t.Error("type did not contain integeger")
		}

		if subject.Minimum != 1 {
			t.Error("minimum was not set correctly")
		}

		if subject.Maximum != 10 {
			t.Error("maximum was not set correctly")
		}

		if subject.MultipleOf != 2 {
			t.Error("multipleOf was not set correctly")
		}
	})

	t.Run("StringBuidler", func(t *testing.T) {
		subject := &schema.Schema{}
		builder := schema.NewStringBuilder(subject)

		builder.Pattern("^asdf$")
		builder.Minimum(1)
		builder.Maximum(10)
		builder.Enum("asdf")

		if subject.Pattern != "^asdf$" {
			t.Error("pattern was not set correctly")
		}

		if subject.MinLength != 1 {
			t.Error("minLength was not set correctly")
		}

		if subject.MaxLength != 10 {
			t.Error("maxLength was not set correctly")
		}

		if !contains(subject.Enum, "asdf") {
			t.Error("enum did not contain asdf")
		}
	})

	t.Run("SchemaFor", func(t *testing.T) {
		t.Run("structs", func(t *testing.T) {
			subject := schema.SchemaFor(&Person{})

			if subject.Id != "Person" {
				t.Error("schema id was not Person")
			}

			if !contains(subject.Type, schema.Object) {
				t.Error("schema was not marked as object")
			}

			if !contains(subject.Required, "name") {
				t.Error("name was not marked as required")
			}

			if len(subject.Properties) != 2 {
				t.Error("there's not excactly 2 fields in the properties")
			}

			nameSchema, nameOk := subject.Properties["name"]
			ageSchema, ageOk := subject.Properties["Age"] // struct field name (since no name in json tag)

			if !nameOk {
				t.Error("name waas not set")
			}

			if contains(nameSchema.Type, schema.Null) {
				t.Error("name was marked as nullable")
			}

			if !ageOk {
				t.Error("age was not set")
			}

			if !contains(ageSchema.Type, schema.Null) {
				t.Error("age was not marked as nullable")
			}

			if ageSchema.Minimum != 18 {
				t.Error("minimum was not set to 18")
			}

			if ageSchema.Maximum != 100 {
				t.Error("maximum was not set to 100")
			}
		})

		t.Run("Ints", func(t *testing.T) {
			subject := schema.SchemaFor(99)

			if !contains(subject.Type, schema.Integer) {
				t.Error("type was not integer")
			}
		})

		t.Run("Decimals", func(t *testing.T) {
			subject := schema.SchemaFor(0.99)

			if !contains(subject.Type, schema.Number) {
				t.Error("type was not number")
			}
		})

		t.Run("Boolean", func(t *testing.T) {
			subject := schema.SchemaFor(false)

			if !contains(subject.Type, schema.Boolean) {
				t.Error("type was not boolean")
			}
		})

		t.Run("Strings", func(t *testing.T) {
			subject := schema.SchemaFor("asdf")

			if !contains(subject.Type, schema.String) {
				t.Error("type was not string")
			}
		})

		t.Run("Arrays", func(t *testing.T) {
			subject := schema.SchemaFor(make([]int, 0))

			if !contains(subject.Type, schema.Array) {
				t.Error("type was not array")
			}

			if subject.Items == nil {
				t.Error("items was not set")
			}

			if !contains(subject.Items.Type, schema.Integer) {
				t.Error("items.type was not integer")
			}
		})
	})
}

func contains[T any](schema []T, kind T) bool {
	return slice.Contains(schema, kind)
}

func (Person) Enchance(field string, schemah *schema.Schema) {
	if field == "Age" {
		builder := schema.NewNumberBuilder(schemah)
		builder.Minimum(18)
		builder.Maximum(100)
	}
}
