package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
	Role        string `json:"role"`
}

var users []User

func (u User) Store() User {

	if u.ID != 0 {
		return u
	}

	u.ID = len(users) + 1

	users = append(users, u)

	return u

}

func FindUserByEmail(email, pass string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == pass {
			return &user
		}
	}
	return nil
}
