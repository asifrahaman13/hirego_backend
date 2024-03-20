// domain/user.go
package domain

type User struct {
    ID   int
    Name string
    Age  int
}

type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
}
