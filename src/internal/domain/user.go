// domain/user.go
package domain

type User struct {
    ID   int `json:"id"`
    Name string `json:"name"`
    Age  int `json:"age"`
}

type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
    FindAll() ([]*User, error)
}
