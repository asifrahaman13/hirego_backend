// domain/user.go
package domain

// import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    // ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    FirstName string            `json:"firstname" bson:"firstname"`
    LastName  string            `json:"lastname" bson:"lastname"`
    Email     string            `json:"email" bson:"email"`
    Password  string            `json:"password" bson:"password"`
}



type UserRepository interface {
    FindByID(id int) (*User, error)
    Save(user *User) error
    FindAll() ([]*User, error)
}
