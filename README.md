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

## Usage

```
curl -L -o openapi-go-server.zip https://github.com/wndhydrnt/openapi-go-server/archive/master.zip
unzip openapi-go-server.zip
rm openapi-go-server.zip
mv openapi-go-server-master/go-server .
rm -rf openapi-go-server-master/
openapi-generator generate -i https://petstore.swagger.io/v2/swagger.json -o . --generator-name go-server --additional-properties=sourceFolder=api,packageName=api -t ./go-server
```
