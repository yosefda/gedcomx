// Package typedef defines all types used in GEDCOM X.
package typedef

// URI is used to identify data types, data instances and elements of the controlled vocabularies.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#134-the-uri-reference
type URI string

// Timestamp is an instance of time, including values for year, month, date, hour, minute, second and timezone (ISO 8601).
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#133-basic-data-types
type Timestamp struct{}

// LocaleTag is an IETF BCP 47 locale tag for identifying language.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#il8n
// https://tools.ietf.org/html/bcp47
type LocaleTag struct{}

// Date is the date representation for exchanging dates associated with genealogical data.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/date-format-specification.md
type Date struct{}
