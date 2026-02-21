package schema

type (
	Builder interface {
		Schema() *Schema
	}

	SchemaBuilder interface {
		Builder
		Object(func(ObjectBuilder))
		Number(func(NumberBuilder))
		String(func(StringBuilder))
		Boolean()
		Array(func(ArrayBuilder))

		Default(any)
	}

	ObjectBuilder interface {
		Builder
		Id(string)
		Property(name string, required bool, nullable bool, cb func(SchemaBuilder))
	}

	NumberBuilder interface {
		Builder
		Integer() // make this number an integer
		Minimum(int64)
		Maximum(int64)
		MultipleOf(int64)
	}

	StringBuilder interface {
		Builder
		Pattern(string)
		Enum(...string)
		Minimum(int64)
		Maximum(int64)
	}

	ArrayBuilder interface {
		Builder
		Minimum(int64)
		Maximum(int64)
		Items(func(SchemaBuilder))
	}
)
