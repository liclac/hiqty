package media

import (
	"net/url"
)

// Global registry of available services.
var Services = make(map[string]Service)

// Registers a service with the registry.
func Register(svc Service) {
	Services[svc.ID()] = svc
}

// A Service facilitates communication with a streaming service of some kind.
type Service interface {
	// An arbitrary ID for the service, used for track serialization.
	ID() string

	// Return true if the URL looks interesting.
	Sniff(u *url.URL) bool

	// Resolve is called if Sniff() returns true, and resolves a URL into one or more tracks.
	Resolve(u *url.URL) ([]Track, error)
}