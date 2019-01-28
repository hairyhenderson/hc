package hc

import (
	"github.com/brutella/hc/util"
)

// Transport provides accessories over a network.
type Transport interface {
	// Start starts the transport
	Start()

	// Stop stops the transport
	// Use the returned channel to wait until the transport is fully stopped.
	Stop() <-chan struct{}

	Storage() util.Storage
}
