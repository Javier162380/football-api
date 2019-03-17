module github.com/Javier162380/football-api.git

require (
	football-api/helpers v0.0.0
	football-api/models v0.0.0
	github.com/aws/aws-sdk-go v1.18.3
	github.com/gorilla/mux v1.7.0
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
)

replace football-api/helpers v0.0.0 => ./helpers

replace football-api/models v0.0.0 => ./models
