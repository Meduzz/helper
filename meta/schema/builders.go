package schema

import (
	"github.com/Meduzz/helper/fp/slice"
)

type (
	schemaBuilder struct {
		schema *Schema
	}

	objectBuilder struct {
		schema *Schema
	}

	numberBuilder struct {
		schema *Schema
	}

	stringBuilder struct {
		schema *Schema
	}

	arrayBuilder struct {
		schema *Schema
	}
)

func NewSchemaBuilder(schema *Schema) SchemaBuilder {
	return &schemaBuilder{schema}
}

func (s *schemaBuilder) Object(cb func(ObjectBuilder)) {
	cb(NewObjectBuilder(s.schema))
}

func (s *schemaBuilder) Number(cb func(NumberBuilder)) {
	cb(NewNumberBuilder(s.schema))
}

func (s *schemaBuilder) String(cb func(StringBuilder)) {
	cb(NewStringBuilder(s.schema))
}

func (s *schemaBuilder) Boolean() {
	s.schema.Type = append(s.schema.Type, Boolean)
}

func (s *schemaBuilder) Array(cb func(ArrayBuilder)) {
	cb(NewArrayBuilder(s.schema))
}

func (s *schemaBuilder) Default(defaultValue any) {
	s.schema.Default = defaultValue
}

func (s *schemaBuilder) Schema() *Schema {
	return s.schema
}

func NewObjectBuilder(schema *Schema) ObjectBuilder {
	if schema.Properties == nil {
		schema.Properties = make(map[string]*Schema)
	}

	schema.Type = append(schema.Type, Object)

	return &objectBuilder{schema}
}

func (o *objectBuilder) Property(property string, required, nullable bool, cb func(SchemaBuilder)) {
	propertySchema := &Schema{}

	if nullable {
		propertySchema.Type = append(propertySchema.Type, Null)
	}

	if required {
		o.schema.Required = append(o.schema.Required, property)
	}

	cb(NewSchemaBuilder(propertySchema))

	o.schema.Properties[property] = propertySchema
}

func (o *objectBuilder) Id(id string) {
	o.schema.Id = id
}

func (o *objectBuilder) Schema() *Schema {
	return o.schema
}

func NewNumberBuilder(schema *Schema) NumberBuilder {
	schema.Type = append(schema.Type, Number)

	return &numberBuilder{schema}
}

func (n *numberBuilder) Integer() {
	// a little bit dirty
	n.schema.Type = slice.Filter(n.schema.Type, func(k Kind) bool {
		return k != Number
	})
	n.schema.Type = append(n.schema.Type, Integer)
}

func (n *numberBuilder) Minimum(min int64) {
	n.schema.Minimum = min
}

func (n *numberBuilder) Maximum(max int64) {
	n.schema.Maximum = max
}

func (n *numberBuilder) MultipleOf(step int64) {
	n.schema.MultipleOf = step
}

func (n *numberBuilder) Schema() *Schema {
	return n.schema
}

func NewStringBuilder(schema *Schema) StringBuilder {
	schema.Type = append(schema.Type, String)

	return &stringBuilder{schema}
}

func (s *stringBuilder) Pattern(pattern string) {
	s.schema.Pattern = pattern
}

func (s *stringBuilder) Enum(enum ...string) {
	s.schema.Enum = append(s.schema.Enum, enum...)
}

func (s *stringBuilder) Minimum(minLength int64) {
	s.schema.MinLength = minLength
}

func (s *stringBuilder) Maximum(maxLength int64) {
	s.schema.MaxLength = maxLength
}

func (s *stringBuilder) Schema() *Schema {
	return s.schema
}

func NewArrayBuilder(schema *Schema) ArrayBuilder {
	schema.Type = append(schema.Type, Array)

	return &arrayBuilder{schema}
}

func (a *arrayBuilder) Minimum(minItems int64) {
	a.schema.MinItems = minItems
}

func (a *arrayBuilder) Maximum(maxItems int64) {
	a.schema.MaxItems = maxItems
}

func (a *arrayBuilder) Items(cb func(SchemaBuilder)) {
	itemSchema := &Schema{}
	cb(NewSchemaBuilder(itemSchema))
	a.schema.Items = itemSchema
}

func (a *arrayBuilder) Schema() *Schema {
	return a.schema
}
