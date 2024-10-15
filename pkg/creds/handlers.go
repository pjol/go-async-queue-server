package creds

import "net/http"

// Checks to see if user has an open exchange that hasn't been filled out yet.
func (s *CredService) HandleCheckExchange(w http.ResponseWriter, r *http.Request) {

}

// Calls out to opencred server to create an exchange for a user.
func (s *CredService) HandlePostExchange(w http.ResponseWriter, r *http.Request) {

}

// Handles the opencred callback, determines the address corresponding to the exchange and enqueues the VP token and address to be sent to the prover.
func (s *CredService) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Parse callback into ExchangeResult struct

	// Check if prover available
	// If so, create async task for looping over queue from next avaibable server
	// If not, add VP and address to queue.
}
