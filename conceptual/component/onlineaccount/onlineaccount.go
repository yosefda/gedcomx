// Package onlineaccount defines data structure to represent a description of an account for an online service provider.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#38-the-onlineaccount-data-type
package onlineaccount

import "github.com/yosefda/gedcomx/typedef"

// Idenfitier is the identifier for the OnlineAccount data type.
const Idenfitier = "http://gedcomx.org/v1/OnlineAccount"

// Properties of the OnlineAccount data type.
type Properties struct {
	ServiceHomepage typedef.URI
	AccountName     string
}
