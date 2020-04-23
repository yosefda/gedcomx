package conceptual

// URI is fundamental to the GEDCOM X conceptual mode.
// It is used to identify data types, data instances and elements of the controlled vocabularies.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#134-the-uri-reference
type URI string

// Timestamp refers to an instance of time, including values for year, month, date, hour, minute, second and timezone (ISO 8601).
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#133-basic-data-types
type Timestamp string
