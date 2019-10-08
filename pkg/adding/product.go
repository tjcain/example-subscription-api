package adding

// Product defines the properties of a product to be added.
type Product struct {
	ID int `json:"id"`

	// ...
}

// AddProduct adds given products to the database.
func (s *service) AddProduct(prods ...Product) error {
	// validation ....

	for _, product := range prods {
		_, err := s.r.AddProduct(product)
		if err != nil {
			return err // or error validation
		}
	}

	return nil
}
