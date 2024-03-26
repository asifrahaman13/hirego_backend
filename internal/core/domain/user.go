// domain/user.go
package domain

type User struct {
	// ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type AccessToken struct {
	Token string `json:"token"`
}

type UserInformation struct {
	Email            string                 `json:"email" bson:"email"`
	FirstName        string                 `json:"firstname" bson:"firstname"`
	LastName         string                 `json:"lastname" bson:"lastname"`
	PhoneNumber      string                 `json:"phonenumber" bson:"phonenumber"`
	DOB              string                 `json:"dob" bson:"dob"`
	Address          string                 `json:"address" bson:"address"`
	ProfilePicture   string                 `json:"profilepicture" bson:"profilepicture"`
	Country          string                 `json:"country" bson:"country"`
	State            string                 `json:"state" bson:"state"`
	PushNotification string                 `json:"pushnotification" bson:"pushnotification"`
	Notificationis   map[string]interface{} `json:"notificationis" bson:"notificationis"`
}
