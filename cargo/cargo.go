package cargo

import (
	"time"

	"greenlight.alexedwards.net/location"
)

type TrackingID string

type Cargo struct {
	TrackingID         TrackingID
	Origin             location.UNLocode
	RouteSpecification RouteSpecification
	Itinerary          Itinerary
	Delivery           Delivery
}

type RouteSpecification struct {
	Origin          location.UNLocode
	Destination     location.UNLocode
	ArrivalDeadline time.Time
}

type RoutingStatus int

type TransportStatus int
