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
	"net/http"
	"os"
)

// Error indicates an error.
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// NewError returns an API error.
func NewError(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

// PetApiRouter defines the required methods for binding the api requests to a responses for the PetApi
// The PetApiRouter implementation should parse necessary information from the http request,
// pass the data to a PetApiServicer to perform the required actions, then write the service results to the http response.
type PetApiRouter interface {
	AddPet(http.ResponseWriter, *http.Request)
	DeletePet(http.ResponseWriter, *http.Request)
	FindPetsByStatus(http.ResponseWriter, *http.Request)
	FindPetsByTags(http.ResponseWriter, *http.Request)
	GetPetById(http.ResponseWriter, *http.Request)
	UpdatePet(http.ResponseWriter, *http.Request)
	UpdatePetWithForm(http.ResponseWriter, *http.Request)
	UploadFile(http.ResponseWriter, *http.Request)
}

// StoreApiRouter defines the required methods for binding the api requests to a responses for the StoreApi
// The StoreApiRouter implementation should parse necessary information from the http request,
// pass the data to a StoreApiServicer to perform the required actions, then write the service results to the http response.
type StoreApiRouter interface {
	DeleteOrder(http.ResponseWriter, *http.Request)
	GetInventory(http.ResponseWriter, *http.Request)
	GetOrderById(http.ResponseWriter, *http.Request)
	PlaceOrder(http.ResponseWriter, *http.Request)
}

// UserApiRouter defines the required methods for binding the api requests to a responses for the UserApi
// The UserApiRouter implementation should parse necessary information from the http request,
// pass the data to a UserApiServicer to perform the required actions, then write the service results to the http response.
type UserApiRouter interface {
	CreateUser(http.ResponseWriter, *http.Request)
	CreateUsersWithArrayInput(http.ResponseWriter, *http.Request)
	CreateUsersWithListInput(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUserByName(http.ResponseWriter, *http.Request)
	LoginUser(http.ResponseWriter, *http.Request)
	LogoutUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
}

// PetApiServicer defines the api actions for the PetApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PetApiServicer interface {
	AddPet(context.Context, Pet) error
	DeletePet(context.Context, int64, string) error
	FindPetsByStatus(context.Context, []string) ([]Pet, error)
	FindPetsByTags(context.Context, []string) ([]Pet, error)
	GetPetById(context.Context, int64) (Pet, error)
	UpdatePet(context.Context, Pet) error
	UpdatePetWithForm(context.Context, int64, string, string) error
	UploadFile(context.Context, int64, string, *os.File) (ApiResponse, error)
}

// StoreApiServicer defines the api actions for the StoreApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type StoreApiServicer interface {
	DeleteOrder(context.Context, int64) error
	GetInventory(context.Context) (map[string]int32, error)
	GetOrderById(context.Context, int64) (Order, error)
	PlaceOrder(context.Context, Order) (Order, error)
}

// UserApiServicer defines the api actions for the UserApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserApiServicer interface {
	CreateUser(context.Context, User) error
	CreateUsersWithArrayInput(context.Context, []User) error
	CreateUsersWithListInput(context.Context, []User) error
	DeleteUser(context.Context, string) error
	GetUserByName(context.Context, string) (User, error)
	LoginUser(context.Context, string, string) (string, error)
	LogoutUser(context.Context) error
	UpdateUser(context.Context, string, User) error
}
