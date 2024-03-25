package domain

type AuthRepository interface {
	Signup(user *User) (interface{}, error)
	// Login(user *User) (interface{}, error)
}