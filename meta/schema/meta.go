package schema

var meta = WithBuider(func(sb SchemaBuilder) {
	sb.Object(func(ob ObjectBuilder) {
		ob.Id("Schema")
		ob.Property("kind", true, false, func(sb SchemaBuilder) {
			sb.Array(func(ab ArrayBuilder) {
				ab.Items(func(sb SchemaBuilder) {
					sb.String(func(sb StringBuilder) {})
				})
			})
		})
		ob.Property("properties", false, true, func(sb SchemaBuilder) {
			sb.Object(func(ob ObjectBuilder) {
				ob.AdditionalProperties(&Schema{
					Ref: "#/$defs/schema",
				})
			})
		})
		ob.Property("pattern", false, true, func(sb SchemaBuilder) {
			sb.String(nil)
		})
		ob.Property("items", false, true, func(sb SchemaBuilder) {
			sb.Array(func(ab ArrayBuilder) {
				ab.Ref("#/$defs/schema")
			})
		})
		ob.Property("enum", false, true, func(sb SchemaBuilder) {
			sb.Array(func(ab ArrayBuilder) {
				ab.Items(func(sb SchemaBuilder) {
					sb.String(nil)
				})
			})
		})
		ob.Property("required", false, true, func(sb SchemaBuilder) {
			sb.Array(func(ab ArrayBuilder) {
				ab.Items(func(sb SchemaBuilder) {
					sb.String(nil)
				})
			})
		})
		ob.Property("default", false, true, func(sb SchemaBuilder) {
			sb.Schema().Type = append(sb.Schema().Type, Object, Array, Number, Integer, Boolean, String)
		})
		ob.Property("minimum", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("maxiumum", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("multipleOf", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("minLength", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("maxLength", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("minItems", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("maxItems", false, true, func(sb SchemaBuilder) {
			sb.Number(func(nb NumberBuilder) {
				nb.Integer()
			})
		})
		ob.Property("$id", false, true, func(sb SchemaBuilder) {
			sb.String(nil)
		})
		ob.Property("$defs", false, true, func(sb SchemaBuilder) {
			sb.Object(func(ob ObjectBuilder) {})
		})
		ob.Property("$ref", false, true, func(sb SchemaBuilder) {
			sb.String(nil)
		})
	})
})
