package schema

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
	if cb != nil {
		cb(NewNumberBuilder(s.schema))
	} else {
		s.schema.Type = Merge(s.schema.Type, Number)
	}
}

func (s *schemaBuilder) String(cb func(StringBuilder)) {
	if cb != nil {
		cb(NewStringBuilder(s.schema))
	} else {
		s.schema.Type = Merge(s.schema.Type, String)
	}
}

func (s *schemaBuilder) Boolean() {
	s.schema.Type = Merge(s.schema.Type, Boolean)
}

func (s *schemaBuilder) Array(cb func(ArrayBuilder)) {
	cb(NewArrayBuilder(s.schema))
}

func (s *schemaBuilder) Default(defaultValue any) {
	s.schema.Default = defaultValue
}

func (s *schemaBuilder) Ref(id string) {
	s.schema.Ref = id
}

func (s *schemaBuilder) Schema() *Schema {
	return s.schema
}

func NewObjectBuilder(schema *Schema) ObjectBuilder {
	if schema.Properties == nil {
		schema.Properties = make(map[string]*Schema)
	}

	schema.Type = Merge(schema.Type, Object)

	return &objectBuilder{schema}
}

func (o *objectBuilder) Property(property string, required, nullable bool, cb func(SchemaBuilder)) {
	propertySchema := &Schema{}

	if nullable {
		propertySchema.Type = Merge(propertySchema.Type, Null)
	}

	if required {
		o.schema.Required = append(o.schema.Required, property)
	}

	cb(NewSchemaBuilder(propertySchema))

	bubbleDefs(o.schema, propertySchema)
	o.schema.Properties[property] = propertySchema
}

func (o *objectBuilder) AdditionalProperties(boolOrSchema any) {
	o.schema.AdditionalProperties = boolOrSchema
}

func (o *objectBuilder) Id(id string) {
	o.schema.Id = id
}

func (o *objectBuilder) Schema() *Schema {
	return o.schema
}

func NewNumberBuilder(schema *Schema) NumberBuilder {
	schema.Type = Merge(schema.Type, Number)

	return &numberBuilder{schema}
}

func (n *numberBuilder) Integer() {
	// a little bit dirty
	n.schema.Type = Drop(n.schema.Type, Number)
	n.schema.Type = Merge(n.schema.Type, Integer)
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
	schema.Type = Merge(schema.Type, String)

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
	schema.Type = Merge(schema.Type, Array)

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

	bubbleDefs(a.schema, itemSchema)
	a.schema.Items = itemSchema
}

func (a *arrayBuilder) Ref(id string) {
	a.schema.Ref = id
}

func (a *arrayBuilder) Schema() *Schema {
	return a.schema
}

func bubbleDefs(parent, child *Schema) {
	if len(child.Defs) == 0 {
		return
	}

	if parent.Defs == nil {
		parent.Defs = make(map[string]*Schema)
	}

	for key, value := range child.Defs {
		parent.Defs[key] = value
		delete(child.Defs, key) // sketchy
	}
}
