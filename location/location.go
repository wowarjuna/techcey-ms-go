package location

import "errors"

type UNLocode string

type Location struct {
	UNLocode UNLocode
	Name     string
}

var ErrUnknown = errors.New("unknown location")
