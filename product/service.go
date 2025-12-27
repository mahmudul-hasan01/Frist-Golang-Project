package product

type service struct {
	prdctRepo ProductRepo
}

func NewServer(prdctRepo ProductRepo) Service {
	return &service{
		prdctRepo: prdctRepo
	}
}


func (svc *service) Create(p domain.Product) (*domain.Product, error){
	return svc.prdctRepo.Create(p)
}
func (svc *service) 	Get(productId int) (*domain.Product, error){
	return svc.prdctRepo.Create(productId)
}
func (svc *service) 	List() ([]*domain.Product, error){
	return svc.prdctRepo.List()
}
func (svc *service) 	Update(p domain.Product) (*domain.Product, error){
	return svc.prdctRepo.Update(p)
}
func (svc *service) 	Delete(productId int) error{
	return svc.prdctRepo.Delete(productId)
}
