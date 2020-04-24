// Package enum specifies a set of enumerated values that identify common event
// types that are relevant to genealogical research.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/event-types-specification.md
package enum

import (
	"github.com/yosefda/gedcomx/typedef"
)

// Event enum type.
type Event typedef.URI

// Event values.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/event-types-specification.md#2-event-types
const (
	Adoption          Event = "http://gedcomx.org/Adoption"
	AdultChristening  Event = "http://gedcomx.org/AdultChristening"
	Annulment         Event = "http://gedcomx.org/Annulment"
	Baptism           Event = "http://gedcomx.org/Baptism"
	BarMitzvah        Event = "http://gedcomx.org/BarMitzvah"
	BatMitzvah        Event = "http://gedcomx.org/BatMitzva"
	Birth             Event = "http://gedcomx.org/Birth"
	Blessing          Event = "http://gedcomx.org/Blessing"
	Burial            Event = "http://gedcomx.org/Burial"
	Census            Event = "http://gedcomx.org/Census"
	Christening       Event = "http://gedcomx.org/Christening"
	Circumcision      Event = "http://gedcomx.org/Circumcision"
	Confirmation      Event = "http://gedcomx.org/Confirmation"
	Cremation         Event = "http://gedcomx.org/Cremation"
	Death             Event = "http://gedcomx.org/Death"
	Divorce           Event = "http://gedcomx.org/Divorce"
	DivorceFiling     Event = "http://gedcomx.org/DivorceFiling"
	Education         Event = "http://gedcomx.org/Education"
	Engagement        Event = "http://gedcomx.org/Engagement"
	Emigration        Event = "http://gedcomx.org/Emigration"
	Excommunication   Event = "http://gedcomx.org/Excommunication"
	FirstCommunion    Event = "http://gedcomx.org/FirstCommunion"
	Funeral           Event = "http://gedcomx.org/Funeral"
	Immigration       Event = "http://gedcomx.org/Immigration"
	LandTransaction   Event = "http://gedcomx.org/LandTransaction"
	Marriage          Event = "http://gedcomx.org/Marriage"
	MilitaryAward     Event = "http://gedcomx.org/MilitaryAward"
	MilitaryDischarge Event = "http://gedcomx.org/MilitaryDischarge"
	Mission           Event = "http://gedcomx.org/Mission"
	MoveFrom          Event = "http://gedcomx.org/MoveFrom"
	MoveTo            Event = "http://gedcomx.org/MoveTo"
	Naturalization    Event = "http://gedcomx.org/Naturalization"
	Ordination        Event = "http://gedcomx.org/Ordination"
	Retirement        Event = "http://gedcomx.org/Retirement"
)
