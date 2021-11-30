package booking

import (
	"errors"
	"time"

	"greenlight.alexedwards.net/cargo"
	"greenlight.alexedwards.net/location"
)

var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	BookNewCargo(origin location.UNLocode, destination location.UNLocode, deadline time.Time) (cargo.TrackingID, error)
}
