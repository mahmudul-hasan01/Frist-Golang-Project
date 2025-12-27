package product


type Service interface {

	prdHandler.Service

}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(productId int) error
}