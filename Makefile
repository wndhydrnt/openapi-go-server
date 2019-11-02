generate_petstore:
	docker run --rm -v $(PWD):/openapi-go-server openapitools/openapi-generator-cli:v4.2.0 generate -i https://petstore.swagger.io/v2/swagger.json -o /openapi-go-server/examples/petstore --generator-name go-server --additional-properties=sourceFolder=api,packageName=api -t /openapi-go-server/go-server
	docker run --rm -v $(PWD):/openapi-go-server -w "/openapi-go-server" golang:1 gofmt -w ./examples/petstore

test: generate_petstore
	docker run --rm -v $(PWD):/openapi-go-server -w "/openapi-go-server/examples/petstore" golang:1 go build
