// domain/user.go
package domain

type User struct {
	// ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type AccessToken struct {
	Token string `json:"token"`
}

type UserInformation struct {
	Email            string                 `json:"email"`
	FirstName        string                 `json:"firstname"`
	LastName         string                 `json:"lastname"`
	PhoneNumber      string                 `json:"phonenumber"`
	DOB              string                 `json:"dob"`
	Address          string                 `json:"address"`
	ProfilePicture   string                 `json:"profilepicture"`
	Country          string                 `json:"country"`
	State            string                 `json:"state"`
	PushNotification string                 `json:"pushnotification"`
	Notificationis   map[string]interface{} `json:"notificationis"`
}
