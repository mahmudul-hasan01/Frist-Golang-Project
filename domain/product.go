package domian


type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	ImgUrl      string  `json:"imageUrl" db:"image_url"`
	Price       float64 `json:"price" db:"price"`
}