module github.com/Javier162380/football-api.git

require (
	football_api/helpers v0.0.0
	football_api/models v0.0.0
	github.com/aws/aws-sdk-go v1.18.3
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
)

replace football_api/helpers v0.0.0 => ./helpers

replace football_api/models v0.0.0 => ./models
