package adding

// Subscription defines the properties of a subscription to be added.
type Subscription struct {
	ID int `json:"id"`

	// .....
}

// AddSubscription adds given subscriptions to the database.
func (s *service) AddSubscription(subs ...Subscription) error {
	// validation ....

	for _, subscription := range subs {
		err := s.r.AddSubscription(subscription)
		if err != nil {
			return err // or error validation
		}
	}

	return nil
}
