package adding

// Event defines possible outcomes from the "adding actor"
type Event int

// Service provides subscription adding opperations.
type Service interface {
	SubscriptionService
	ProductService
}

// Repository provides access to persistant storage capable of processing
// subscriptions and products.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddSubscription(Subscription) (int, error)
	AddProduct(Product) (int, error)
}

// SubscriptionService is an interface representing the ability to add
// subscriptions.
type SubscriptionService interface {
	AddSubscription(...Subscription) error
}

// ProductService is an interface representing the ability to add
// subscriptions.
type ProductService interface {
	AddProduct(...Product) error
}

type service struct {
	r Repository

	// other deps, e.g. logger.
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
