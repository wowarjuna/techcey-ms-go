package cargo

import (
	"time"

	"greenlight.alexedwards.net/location"
	"greenlight.alexedwards.net/voyage"
)

type Leg struct {
	VoyageNumber   voyage.Number     `json:"voyage_number"`
	LoadLocation   location.UNLocode `json:"from"`
	UnloadLocation location.UNLocode `json:"to"`
	LoadTime       time.Time         `json:"load_time"`
	UnloadTime     time.Time         `json:"unload_time"`
}
type Itinerary struct {
	Legs []Leg `json:"legs"`
}
