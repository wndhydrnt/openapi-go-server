/*
 * Swagger Petstore
 *
 * This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.3
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/api"
	//
	"github.com/wndhydrnt/openapi-go-server/examples/petstore/api"
)

func main() {
	log.Printf("Server started")

	PetApiService := api.NewPetApiService()
	PetApiController := api.NewPetApiController(PetApiService)

	StoreApiService := api.NewStoreApiService()
	StoreApiController := api.NewStoreApiController(StoreApiService)

	UserApiService := api.NewUserApiService()
	UserApiController := api.NewUserApiController(UserApiService)

	router := api.NewRouter(PetApiController, StoreApiController, UserApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
