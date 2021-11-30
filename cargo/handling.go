package cargo

import (
	"greenlight.alexedwards.net/location"
	"greenlight.alexedwards.net/voyage"
)

type HandlingEventType int

type HandlingActivity struct {
	Type         HandlingEventType
	Location     location.UNLocode
	VoyageNumber voyage.Number
}

type HandlingEvent struct {
	TrackingID TrackingID
	Activity   HandlingActivity
}
