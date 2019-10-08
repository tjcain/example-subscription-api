package storage

import "github.com/tjcain/example-subscription-api/pkg/adding"

// SubscriptionStore is an interface representing the ability to handle
// subscriptions.
type SubscriptionStore interface {
	AddSubscription(adding.Subscription) (int, error)
	UpdateSubscription(Subscription) error
	DeleteSubscription(int) error
}

// ProductStore is an interface representing the ability to handle plans.
type ProductStore interface {
	AddProduct(adding.Product) (int, error)
	UpdateProduct(Product) error
	DeleteProduct(int) error
}

// Store is the interface to the backing store.
type Store interface {
	SubscriptionStore
	ProductStore

	// Migrate migrates the database.
	Migrate()
}
