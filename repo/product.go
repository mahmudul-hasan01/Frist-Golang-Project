package repo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// type Product struct {
// 	ID          int     `json:"id" db:"id"`
// 	Title       string  `json:"title" db:"title"`
// 	Description string  `json:"description" db:"description"`
// 	ImgUrl      string  `json:"imageUrl" db:"image_url"`
// 	Price       float64 `json:"price" db:"price"`
// }

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

--------------------------------------------------
-- CREATE
--------------------------------------------------

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {

	query := `
		INSERT INTO products (title, description, image_url, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImgUrl,
		p.Price,
	).Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

--------------------------------------------------
-- GET BY ID
--------------------------------------------------

func (r *productRepo) Get(productId int) (*domain.Product, error) {

	query := `
		SELECT id, title, description, image_url, price
		FROM products
		WHERE id = $1
	`

	var p domain.Product
	err := r.db.Get(&p, query, productId)

	if err != nil {  
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

--------------------------------------------------
-- LIST ALL
--------------------------------------------------

func (r *productRepo) List() ([]*domain.Product, error) {

	query := `
		SELECT id, title, description, image_url, price
		FROM products
		ORDER BY id DESC
	`

	var products []*domain.Product
	err := r.db.Select(&products, query)

	if err != nil {
		return nil, err
	}

	return products, nil
}

--------------------------------------------------
-- UPDATE
--------------------------------------------------

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {

	query := `
		UPDATE products
		SET
			title = $1,
			description = $2,
			image_url = $3,
			price = $4
		WHERE id = $5
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImgUrl,
		p.Price,
		p.ID,
	).Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

--------------------------------------------------
-- DELETE
--------------------------------------------------

func (r *productRepo) Delete(productId int) error {

	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(query, productId)
	return err
}
