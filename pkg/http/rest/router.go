package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjcain/example-subscription-api/pkg/adding"
)

// Handler returns a http.Handler with given configured services.
func Handler(a adding.Service) http.Handler {
	r := chi.NewRouter()

	// apply middleware.
	r.Post("/subscription", addSubscription(a))
	r.Post("/product", addProduct(a))

	return r
}

// addSubscription returns a handler for POST /subscription requests.
func addSubscription(s adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var newSub adding.Subscription
		err := decoder.Decode(&newSub)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// lazy error handling..!
		s.AddSubscription(newSub)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New subscription added.")
	}
}

// addProduct returns a handler for POST /subscription requests.
func addProduct(s adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var newProduct adding.Product
		err := decoder.Decode(&newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// lazy error handling..!
		s.AddProduct(newProduct)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New subscription added.")
	}
}
