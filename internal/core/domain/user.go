package domain

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserName struct {
	Username string `json:"username" bson:"username"`
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

type WorkExperience struct {
	Company          string            `json:"company" bson:"company"`
	Position         string            `json:"position" bson:"position"`
	StartDate        string            `json:"startdate" bson:"startdate"`
	EndDate          string            `json:"enddate" bson:"enddate"`
	Stipend          string            `json:"stipend" bson:"stipend"`
	Responsibilities string            `json:"responsibilities" bson:"responsibilities"`
	Skills           map[string]string `json:"skills" bson:"skills"`
	About            string            `json:"about" bson:"about"`
	WorkAccomplished string            `json:"workaccomplished" bson:"workaccomplished"`
}
type Projects struct {
	ProjectName        string            `json:"projectname" bson:"projectname"`
	ProjectDescription string            `json:"projectdescription" bson:"projectdescription"`
	StartDate          string            `json:"startdate" bson:"startdate"`
	EndDate            string            `json:"enddate" bson:"enddate"`
	ProjectLink        string            `json:"projectlink" bson:"projectlink"`
	ProjectSkills      map[string]string `json:"projectskills" bson:"projectskills"`
	ProjectAbout       string            `json:"projectabout" bson:"projectabout"`
}

type WorkInformation struct {
	Useremail        string                 `json:"useremail" bson:"useremail"`
	ExpectedStripend string                 `json:"expectedstripend" bson:"expectedstripend"`
	Skils            map[string]interface{} `json:"skils" bson:"skils"`
	About            string                 `json:"about" bson:"about"`
	Hackathons       []string               `json:"hackathons" bson:"hackathons"`
	Projects         []Projects             `json:"projects" bson:"projects"`
	WorkExperience   []WorkExperience       `json:"workexperience" bson:"workexperience"`
}
