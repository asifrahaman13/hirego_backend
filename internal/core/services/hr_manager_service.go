package service

import (
	"fmt"
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
