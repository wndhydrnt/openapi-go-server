/*
 * Swagger Petstore
 *
 * This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.3
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"errors"
)

// UserApiService is a service that implents the logic for the UserApiServicer
// This service should implement the business logic for every endpoint for the UserApi API.
// Include any external packages or services that will be required by this service.
type UserApiService struct {
}

// NewUserApiService creates a default api service
func NewUserApiService() UserApiServicer {
	return &UserApiService{}
}

// CreateUser - Create user
func (s *UserApiService) CreateUser(ctx context.Context, body User) (err error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'CreateUser' not implemented")
}

// CreateUsersWithArrayInput - Creates list of users with given input array
func (s *UserApiService) CreateUsersWithArrayInput(ctx context.Context, body []User) (err error) {
	// TODO - update CreateUsersWithArrayInput with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'CreateUsersWithArrayInput' not implemented")
}

// CreateUsersWithListInput - Creates list of users with given input array
func (s *UserApiService) CreateUsersWithListInput(ctx context.Context, body []User) (err error) {
	// TODO - update CreateUsersWithListInput with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'CreateUsersWithListInput' not implemented")
}

// DeleteUser - Delete user
func (s *UserApiService) DeleteUser(ctx context.Context, username string) (err error) {
	// TODO - update DeleteUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'DeleteUser' not implemented")
}

// GetUserByName - Get user by user name
func (s *UserApiService) GetUserByName(ctx context.Context, username string) (result User, err error) {
	// TODO - update GetUserByName with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return result, errors.New("service method 'GetUserByName' not implemented")
}

// LoginUser - Logs user into the system
func (s *UserApiService) LoginUser(ctx context.Context, username string, password string) (result string, err error) {
	// TODO - update LoginUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return result, errors.New("service method 'LoginUser' not implemented")
}

// LogoutUser - Logs out current logged in user session
func (s *UserApiService) LogoutUser(ctx context.Context) (err error) {
	// TODO - update LogoutUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'LogoutUser' not implemented")
}

// UpdateUser - Updated user
func (s *UserApiService) UpdateUser(ctx context.Context, username string, body User) (err error) {
	// TODO - update UpdateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return errors.New("service method 'UpdateUser' not implemented")
}
