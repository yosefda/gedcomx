// Package address defines data structure for a street or postal address of a person or organization.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#39-the-address-data-type
package address

// URI is the identifier for the Address data type.
const URI = "http://gedcomx.org/v1/Address"

// Address data type.
type Address struct {
	Value           string
	City            string
	Country         string
	PostalCode      string
	StateOrProvince string
	Street          string
	Street2         string
	Street3         string
	Street4         string
	Street5         string
	Street6         string
}
