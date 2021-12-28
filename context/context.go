package context1

import (
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get the context from the request
		ctx := r.Context()

		// Our long-running request will need a way to report back
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		// One of the following things is gonna happen:
		select {
		// either our goroutine is going to report back using our data channel
		case d := <-data:
			fmt.Fprint(w, d)
		// or that context is going to be terminated, in which case we need to cancel out of it.
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
