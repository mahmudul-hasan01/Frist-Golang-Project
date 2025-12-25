package repo

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
	Get(userId int) (*User, error)
	Find(email, pass string) (*User, error)
	List() ([]*User, error)
	Delete(userId int) (*User, error)
	Update(p User) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	generateUser(repo)
	return repo
}

func (r *userRepo) Create(p User) (*User, error) {
	p.ID = len(r.userList) + 1
	r.userList = append(r.userList, &p)
	return &p, nil
}
func (r *userRepo) Get(userId int) (*User, error) {
	for _, user := range r.userList {
		if user.ID == userId {
			return user, nil
		}
	}
	return nil, nil
}
func (r *userRepo) List() ([]*User, error) {
	return r.userList, nil
}
func (r *userRepo) Update(user User) (*User, error) {
	for i, p := range r.userList {
		if p.ID == user.ID {
			r.userList[i] = &user
		}
	}
	return &user, nil
}
func (r *userRepo) Delete(userId int) (*User, error) {
	var tempList []*User

	for _, p := range r.userList {
		if p.ID != userId {
			tempList = append(tempList, p)
		}
	}
	r.userList = tempList
	return nil, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	for _, user := range r.userList {
		if user.Email == email && user.Password == pass {
			return user, nil
		}
	}
	return nil, nil
}

func generateUser(r *userRepo) {

	prd1 := &User{ID: 1}
	prd2 := &User{ID: 2}
	prd3 := &User{ID: 3}

	r.userList = append(r.userList, prd1)
	r.userList = append(r.userList, prd2)
	r.userList = append(r.userList, prd3)
}
