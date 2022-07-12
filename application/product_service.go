package application

type ProductService struct {
	Percistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Percistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

//func (s *ProductService) Create(name string, price float64) (ProductInterface, error)
//func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error)
//func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error)
