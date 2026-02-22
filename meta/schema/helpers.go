package schema

import (
	"reflect"
	"strings"
)

func SchemaFor(it any) *Schema {
	s := &Schema{}

	v := reflect.TypeOf(it)

	schemaFor(v, s)

	return s
}

func schemaFor(v reflect.Type, schema *Schema) {
	switch v.Kind() {
	case reflect.Pointer:
		schemaFor(v.Elem(), schema) // drop pointer
	case reflect.Struct:
		schemaForStruct(v, schema)
	case reflect.Slice, reflect.Array:
		schemaForArray(v, schema)
	case reflect.Bool:
		builder := NewSchemaBuilder(schema)
		builder.Boolean()
	case reflect.String:
		builder := NewSchemaBuilder(schema)
		builder.String(func(sb StringBuilder) {})
	case reflect.Int, reflect.Int16, reflect.Int64, reflect.Int8:
		builder := NewSchemaBuilder(schema)
		builder.Number(func(nb NumberBuilder) {
			nb.Integer()
		})
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		builder := NewSchemaBuilder(schema)
		builder.Number(func(nb NumberBuilder) {
			nb.Integer()
			nb.Minimum(0)
		})
	case reflect.Float32, reflect.Float64:
		builder := NewSchemaBuilder(schema)
		builder.Number(func(nb NumberBuilder) {})
	}
	// TODO maps? but unles key is string...?
}

func schemaForStruct(v reflect.Type, schema *Schema) {
	builder := NewObjectBuilder(schema)
	builder.Id(v.Name())

	fieldCount := v.NumField()

	for i := 0; i < fieldCount; i++ {
		rf := v.Field(i)

		if rf.IsExported() {
			name := rf.Name
			nullable := false
			required := true

			jsonTag, ok := rf.Tag.Lookup("json")

			if ok {
				jsonTag = strings.Replace(jsonTag, "string", "", -1)

				// ignore json ignored fields
				if jsonTag == "-" {
					continue
				}

				tagContent := strings.Split(jsonTag, ",")

				for _, it := range tagContent {
					if it != "" && it != "-" {
						if it == "omitempty" {
							nullable = true
							required = false
						} else {
							// ~sketchy~ working 🎉
							name = it
						}
					}
				}
			}

			builder.Property(name, required, nullable, func(sb SchemaBuilder) {
				schemaFor(rf.Type, sb.Schema())

				// give structs that implements SchemaHook, speciul treatment.
				if v.AssignableTo(reflect.TypeOf((*SchemaHook)(nil)).Elem()) {
					vValue := reflect.New(v).Interface()
					hook, ok := vValue.(SchemaHook)

					if ok {
						hook.Enchance(name, sb.Schema())
					}
				}
			})
		}
	}
}

func schemaForArray(v reflect.Type, schema *Schema) {
	builder := NewArrayBuilder(schema)
	builder.Items(func(sb SchemaBuilder) {
		schemaFor(v.Elem(), sb.Schema())
	})
}
