package database

var productsList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imageUrl"`
	Price       float64 `json:"price"`
}

func Store(p Product) Product {
	p.ID = len(productsList) + 1
	productsList = append(productsList, p)
	return p
}

func List() []Product {
	return productsList
}

func Get(productId int) *Product {
	for _, product := range productsList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for i, p := range productsList {
		if p.ID == product.ID {
			productsList[i] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product

	for _, p := range productsList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}
	productsList = tempList
}

func init() {

	prd1 := Product{ID: 1, Title: "Product 1", Description: "Description 1", ImgUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Orange-Fruit-Pieces.jpg/2560px-Orange-Fruit-Pieces.jpg", Price: 100}
	prd2 := Product{ID: 2, Title: "Product 2", Description: "Description 2", ImgUrl: "https://images.pexels.com/photos/206959/pexels-photo-206959.jpeg?cs=srgb&dl=pexels-pixabay-206959.jpg&fm=jpg", Price: 200}
	prd3 := Product{ID: 3, Title: "Product 3", Description: "Description 3", ImgUrl: "https://media.istockphoto.com/id/1224636159/photo/closeup-of-a-red-guava-cut-in-half-in-the-background-several-guavas-and-green-leaf.jpg?s=612x612&w=0&k=20&c=KJ9YilkRRuFh0bnw64Ol0IZDfoQF7UIxyC6dRVIjaoA=", Price: 300}

	productsList = append(productsList, prd1)
	productsList = append(productsList, prd2)
	productsList = append(productsList, prd3)
}
