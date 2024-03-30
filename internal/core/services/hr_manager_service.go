package service

import (
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/bson"
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

	// Marshal the MongoDB document into a byte slice.
	bsonBytes, err := bson.Marshal(hrInformation)
	if err != nil {
		return domain.HrProfileInformation{}, err
	}

	// Unmarshal the byte slice into a domain.HrProfileInformation struct.
	var hr domain.HrProfileInformation
	err = bson.Unmarshal(bsonBytes, &hr)
	if err != nil {
		return domain.HrProfileInformation{}, err
	}

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
	jobPosting, err := s.repo.GetByField("jobID", username, "jobposting")
	if err != nil {
		panic(err)
	}

	// Marshal the MongoDB document into a byte slice.
	bsonBytes, err := bson.Marshal(jobPosting)
	if err != nil {
		return domain.JobPosting{}, err
	}

	// Unmarshal the byte slice into a domain.JobPosting struct.
	var job domain.JobPosting
	err = bson.Unmarshal(bsonBytes, &job)
	if err != nil {
		return domain.JobPosting{}, err
	}

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
		// Marshal the MongoDB document into a byte slice.
		bsonBytes, err := bson.Marshal(jp)
		if err != nil {
			return nil, err
		}

		// Unmarshal the byte slice into a domain.JobPosting struct.
		var posting domain.JobPosting
		err = bson.Unmarshal(bsonBytes, &posting)
		if err != nil {
			return nil, err
		}

		postings = append(postings, posting)
	}

	return postings, nil
}
