// Package conclusion defines the data structure for the abstract concept for a basic genealogical data item.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#310-the-conclusion-data-type
package conclusion

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	"github.com/yosefda/gedcomx/conceptual/component/note"
	"github.com/yosefda/gedcomx/conceptual/component/sourcereference"
	enum "github.com/yosefda/gedcomx/enum/confidence"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the Conclusion data type.
const URI = "http://gedcomx.org/v1/Conclusion"

// Conclusion data type.
type Conclusion struct {
	ID          string
	Lang        typedef.LocaleTag
	Sources     *[]sourcereference.SourceReference
	analysis    typedef.URI
	notes       *[]note.Note
	confidence  enum.Confidence
	attribution *attribution.Attribution
}
