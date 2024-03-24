package domain



type AuthRepository interface {
	Signup(user *User) (interface{}, error)
}