# openapi-go-server

Templates for the go-server generator of the [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

The templates are a fork of the [original go-server generator](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go-server) templates.

## Features and Changes

* Enable [DisallowUnknownFields](https://golang.org/pkg/encoding/json/#Decoder.DisallowUnknownFields) on JSON decoder
* Optionally set response code and body on error
* Remove request logging
* Optionally set a logger to log errors in Controllers
* Pass context of the incoming `http.Request` to methods of a Service
* Let methods of a Service return types

## Known Caveats

* Supports only one return type on methods of a Service
* Supports JSON encoding of response body only

## Example

See [examples/petstore](examples/petstore).
