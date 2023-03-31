package aggregate

import "github.com/samverrall/sitesmiths-api/pkg/events"

// Root is used to represent aggregate roots in a domain.
// It also contains an event que based on the observer
// behavorial design pattern.
//
// See: https://refactoring.guru/design-patterns/observer/go/example#example-0
type Root struct {
	Events []events.Observer
}
