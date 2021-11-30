package cargo

import (
	"time"

	"greenlight.alexedwards.net/location"
	"greenlight.alexedwards.net/voyage"
)

type Delivery struct {
	Itinerary               Itinerary
	RouteSpecification      RouteSpecification
	RoutingStatus           RoutingStatus
	TransportStatus         TransportStatus
	NextExpectedActivity    HandlingActivity
	LastEvent               HandlingEvent
	LastKnownLocation       location.UNLocode
	Currentvoyage           voyage.Number
	ETA                     time.Time
	IsMisdirected           bool
	IsUnloadedAtDestination bool
}
