package voyage

import (
	"time"

	"greenlight.alexedwards.net/location"
)

type Number string

type Schedule struct {
	CarrierMovements []CarrierMovement
}

type CarrierMovement struct {
	DepartureLocation location.UNLocode
	ArrivalLocation   location.UNLocode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

type Voyage struct {
	Number   Number
	Schedule Schedule
}

func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}
