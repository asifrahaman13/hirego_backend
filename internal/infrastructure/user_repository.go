// package infrastructure

// import (
// 	"context"
// 	"fmt"
// 	"github.com/asifrahaman13/hirego/src/internal/domain"
// 	"github.com/asifrahaman13/hirego/src/internal/services"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// type UserRepository struct {
// 	// Any dependencies or configuration needed
// }

// // FindAll implements domain.UserRepository.
// // FindAll implements domain.UserRepository.
// func (r *UserRepository) FindAll() ([]*domain.User, error) {
// 	// Simulating a database query to find all users
// 	mongodb_client, err := services.Services()
   
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	coll := mongodb_client.Database("hirego").Collection("users")
    
// 	// Create the filters. 
// 	fillter := bson.D{}
    
// 	// Find the users.
// 	cursor, err := coll.Find(context.TODO(), fillter)

// 	fmt.Printf("cursor: %v\n", cursor)

// 	if err != nil {
// 		panic(err)
// 	}

//     // Define the cursor data type. 
// 	var results []*domain.User
	
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		panic(err)	}

// 	// Return a slice of pointers to user objects
// 	return results, nil
// }

// // Save implements domain.UserRepository.
// func (r *UserRepository) Save(user *domain.User) error {
// 	panic("unimplemented")
// }

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{}
// }

// func (r *UserRepository) FindByID(id int) (*domain.User, error) {
// 	// Simulating a database query to find user by ID
// 	// For demonstration purposes, we'll just return a hardcoded user

// 	// Dummy user data
// 	 user := &domain.User{
//         FirstName: "John",
//         LastName:  "Doe",
//         Email:     "john.doe@example.com",
//         Password:  "password",
//     }


// 	fmt.Print(*user)

// 	return user, nil
// }
