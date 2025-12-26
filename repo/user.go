package repo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
	Role        string `json:"role"`
}

type UserRepo interface {
	Create(p User) (*User, error)
	// Get(userId int) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) (*User, error)
	// Update(p User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(p User) (*User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner, role)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.FirstName,
		p.LastName,
		p.Email,
		p.Password,
		p.IsShopOwner,
		p.Role,
	).Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner, role
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		// If no rows found, return nil without error
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// func (r *userRepo) Get(userId int) (*User, error) {
// 	for _, user := range r.userList {
// 		if user.ID == userId {
// 			return user, nil
// 		}
// 	}
// 	return nil, nil
// }
// func (r *userRepo) List() ([]*User, error) {
// 	return r.userList, nil
// }
// func (r *userRepo) Update(user User) (*User, error) {
// 	for i, p := range r.userList {
// 		if p.ID == user.ID {
// 			r.userList[i] = &user
// 		}
// 	}
// 	return &user, nil
// }
// func (r *userRepo) Delete(userId int) (*User, error) {
// 	var tempList []*User

// 	for _, p := range r.userList {
// 		if p.ID != userId {
// 			tempList = append(tempList, p)
// 		}
// 	}
// 	r.userList = tempList
// 	return nil, nil
// }
