package presentation

/*
What is a "type" and what is a hint of how to display/format it?

It's about form and form data.
text -> input[text], select, radion, checkbox
number -> input[number]
colour -> input[colour]
date -> input[date]
datatime -> input[datetime]
email -> input[email]
url -> input[url]
longtext -> textarea
uuid -> "generated"
money -> input[number]?, input[text]?
regexp -> input[text]
tag -> input[text] + special?
*/
type (
	Format string // (presentation)

	// TODO form hint field multiselt, select, radio, checkbox (presentation)
	// TODO number or rows in the textarea for longtext? (presentation)
	// TODO field to indicate that this field can be label when entity is used in dropdown (or value on reference subtype) (presentation)
	// TODO field to indicate that the value is generated elsewhere (like uuid, hashes, pk etc.)
	// TODO field to indicate number of decimals in decimal subtype? (presentation)
	// TODO format for id?
	// TODO format for number and decimal?
	// TODO format for select/radio/checkbox/select[multi]?
)

const (
	Text     = Format("text")
	Longtext = Format("longtext")
	Date     = Format("date")
	Datetime = Format("datetime")
	Tag      = Format("tag")
	Colour   = Format("colour")
	Email    = Format("email")
	Url      = Format("url")
	UUID     = Format("uuid")
	Money    = Format("money")
	Regex    = Format("regex")
)
