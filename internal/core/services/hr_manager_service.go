package service

import (
	"fmt"
	"strings"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type hrService struct {
	repo ports.HRRepository
}

func InitializeHRService(r ports.HRRepository) *hrService {
	return &hrService{
		repo: r,
	}
}

func (s *hrService) Signup(hr domain.HrManager) (string, error) {

	// Call the sign up repo to insert the data of the user.
	message, err := s.repo.Create(hr, "hrmanager")

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}

func (s *hrService) Login(hr domain.HrManager) (domain.AccessToken, error) {

	// Call the login repo to insert the data of the user.
	token, err := helper.CreateToken(hr.Username)

	if err != nil {
		panic(err)
	}

	// Return the success message.

	accessToken := domain.AccessToken{
		Token: token,
	}
	return accessToken, nil
}

func (s *hrService) SetHrProfileInformation(username string, hrInformation domain.HrProfileInformation) (string, error) {

	hrInformation.Username = username

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(hrInformation, "hrprofilenformation")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil
	}

	// Return the success message.
	return "Hr information stored successfully", nil
}

func (s *hrService) GetProfileInformation(username string) (domain.HrProfileInformation, error) {

	// Call the login repo to insert the data of the user.
	hrInformation, err := s.repo.GetByField("username", username, "hrprofilenformation")

	if err != nil {
		panic(err)
	}

	// Convert the HrInformation to a map[string]interface{}.
	hrMap, ok := hrInformation.(map[string]interface{})

	if !ok {
		return domain.HrProfileInformation{}, fmt.Errorf("userInformation is not a map[string]interface{}")
	}

	// Convert the map to a domain.HrProfileInformation struct.
	hr := domain.HrProfileInformation{
		ID:               hrMap["_id"].(primitive.ObjectID),
		Username:         hrMap["username"].(string),
		Email:            hrMap["email"].(string),
		FirstName:        hrMap["firstname"].(string),
		LastName:         hrMap["lastname"].(string),
		PhoneNumber:      hrMap["phonenumber"].(string),
		DOB:              hrMap["dob"].(string),
		Address:          hrMap["address"].(string),
		ProfilePicture:   hrMap["profilepicture"].(string),
		Country:          hrMap["country"].(string),
		State:            hrMap["state"].(string),
		PushNotification: hrMap["pushnotification"].(string),
		Notificationis:   hrMap["notificationis"].(map[string]interface{}),
	}

	// Return the success message.
	return hr, nil
}

func (s *hrService) JobPosting(jobPosting domain.JobPosting) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(jobPosting, "jobposting")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Failed to insert data", nil
	}

	// Return the success message.
	return "Job posted successfully", nil
}

func (s *hrService) GetJobPosting(username string) (domain.JobPosting, error) {

	// Call the login repo to insert the data of the user.
	jobPosting, err := s.repo.GetByField("username", username, "jobposting")

	if err != nil {
		panic(err)
	}

	// Convert the HrInformation to a map[string]interface{}.
	jobMap, ok := jobPosting.(map[string]interface{})

	if !ok {
		return domain.JobPosting{}, fmt.Errorf("userInformation is not a map[string]interface{}")
	}

	// Convert the map to a domain.HrProfileInformation struct.
	// Convert the map to a domain.JobPosting struct.
	job := domain.JobPosting{
		ID:               jobMap["_id"].(primitive.ObjectID),
		UserID:           jobMap["userID"].(string),
		Title:            jobMap["jobTitle"].(string),
		Description:      jobMap["jobDesc"].(string),
		Responsibilities: strings.Split(jobMap["jobResponsibilities"].(string), ","),
		Requirements:     strings.Split(jobMap["jobRequirements"].(string), ","),
		Skills:           strings.Split(jobMap["jobSkills"].(string), ","),
		Benefits:         strings.Split(jobMap["jobBenefits"].(string), ","),
		Location:         jobMap["jobLocation"].(string),
		EmploymentType:   jobMap["jobEmploymentType"].(string),
		Industry:         jobMap["jobIndustry"].(string),
		Company: struct {
			Name        string `json:"name" bson:"name"`
			Description string `json:"description" bson:"description"`
			Size        string `json:"size" bson:"size"`
			Type        string `json:"type" bson:"type"`
			Industry    string `json:"industry" bson:"industry"`
			Website     string `json:"website" bson:"website"`
		}{
			Name:        jobMap["companyName"].(string),
			Description: jobMap["companyDescription"].(string),
			Size:        jobMap["companySize"].(string),
			Type:        jobMap["companyType"].(string),
			Industry:    jobMap["companyIndustry"].(string),
			Website:     jobMap["companyWebsite"].(string),
		},
		Contact: struct {
			Name  string `json:"name" bson:"name"`
			Email string `json:"email" bson:"email"`
			Phone string `json:"phone" bson:"phone"`
		}{
			Name:  jobMap["contactName"].(string),
			Email: jobMap["contactEmail"].(string),
			Phone: jobMap["contactPhone"].(string),
		},
		Salary: struct {
			Min      float64 `json:"min" bson:"min"`
			Max      float64 `json:"max" bson:"max"`
			Currency string  `json:"currency" bson:"currency"`
			Period   string  `json:"period" bson:"period"`
		}{
			Min:      jobMap["salaryMin"].(float64),
			Max:      jobMap["salaryMax"].(float64),
			Currency: jobMap["salaryCurrency"].(string),
			Period:   jobMap["salaryPeriod"].(string),
		},
		Remote:      jobMap["remote"].(bool),
		PublishedAt: jobMap["publishedAt"].(string),
		ExpiresAt:   jobMap["expiresAt"].(string),
	}

	// Return the success message.
	return job, nil
}



func (s *hrService) GetAllJobPosting() ([]domain.JobPosting, error) {
	// Call the repository to get all job postings.
	jobPostings, err := s.repo.GetAll("jobposting")
	if err != nil {
		return nil, err
	}

	var postings []domain.JobPosting
	for _, jp := range jobPostings {

		// Convert the map to a domain.JobPosting struct.
		posting := domain.JobPosting{
			ID:               jp["_id"].(primitive.ObjectID),
			JobID:            jp["jobID"].(string),
			UserID:           jp["userID"].(string),
			Title:            jp["title"].(string),
			Description:      jp["description"].(string),
			Responsibilities: helper.ConvertToSlice(jp["requirements"].(primitive.A)),
			Requirements:     helper.ConvertToSlice(jp["requirements"].(primitive.A)),
			Skills:           helper.ConvertToSlice(jp["skills"].(primitive.A)),
			Benefits:         helper.ConvertToSlice(jp["benefits"].(primitive.A)),
			Location:       jp["location"].(string),
			EmploymentType: jp["employmentType"].(string),
			Industry:       jp["industry"].(string),
			Company: struct {
				Name        string `json:"name" bson:"name"`
				Description string `json:"description" bson:"description"`
				Size        string `json:"size" bson:"size"`
				Type        string `json:"type" bson:"type"`
				Industry    string `json:"industry" bson:"industry"`
				Website     string `json:"website" bson:"website"`
			}{
				Name:        jp["company"].(map[string]interface{})["name"].(string),
				Description: jp["company"].(map[string]interface{})["description"].(string),
				Size:        jp["company"].(map[string]interface{})["size"].(string),
				Type:        jp["company"].(map[string]interface{})["type"].(string),
				Industry:    jp["company"].(map[string]interface{})["industry"].(string),
				Website:     jp["company"].(map[string]interface{})["website"].(string),
			},
			Contact: struct {
				Name  string `json:"name" bson:"name"`
				Email string `json:"email" bson:"email"`
				Phone string `json:"phone" bson:"phone"`
			}{
				Name:  jp["contact"].(map[string]interface{})["name"].(string),
				Email: jp["contact"].(map[string]interface{})["email"].(string),
				Phone: jp["contact"].(map[string]interface{})["phone"].(string),
			},
			Salary: struct {
				Min      float64 `json:"min" bson:"min"`
				Max      float64 `json:"max" bson:"max"`
				Currency string  `json:"currency" bson:"currency"`
				Period   string  `json:"period" bson:"period"`
			}{
				Min:      jp["salary"].(map[string]interface{})["min"].(float64),
				Max:      jp["salary"].(map[string]interface{})["max"].(float64),
				Currency: jp["salary"].(map[string]interface{})["currency"].(string),
				Period:   jp["salary"].(map[string]interface{})["period"].(string),
			},
			Remote:      jp["remote"].(bool),
			PublishedAt: jp["publishedAt"].(string),
			ExpiresAt:   jp["expiresAt"].(string),
		}
		postings = append(postings, posting)
	}

	return postings, nil
}
