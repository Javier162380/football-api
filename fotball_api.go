package main

import (
	"fmt"
	"encoding/json"
	"football_api/helpers"
	"football_api/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	session := session.New()
	s3_client := s3.New(session, aws.NewConfig().WithRegion("us-east-1"))
	GamesData, err := helpers.DownloadS3Data(s3_client, helpers.Bucket_name,
		helpers.Games_key)

	if err != nil {
		log.Fatal(err)
	}

	GamesStats := models.Games{}

	err = json.Unmarshal(GamesData, &GamesStats)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GamesStats Already downloaded ", GamesStats)
}
