// domain/user.go
package domain


type User struct {
    // ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    FirstName string            `json:"firstname" bson:"firstname"`
    LastName  string            `json:"lastname" bson:"lastname"`
    Email     string            `json:"email" bson:"email"`
    Password  string            `json:"password" bson:"password"`
}



type AccessToken struct {
    Token string `json:"token"`
}