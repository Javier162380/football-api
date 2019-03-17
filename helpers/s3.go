package helpers

import (
	"bytes"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DownloadS3Data(s3Client *s3.S3, bucket string, key string) ([]byte, error) {
	results, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadS3Data(bucket string, key string) (bytedata []byte) {

	session := session.New()
	s3_client := s3.New(session, aws.NewConfig().WithRegion("us-east-1"))
	bufferdata, err := DownloadS3Data(s3_client,
		bucket,
		key)
	if err != nil {
		log.Fatal(err)
	}
	return bufferdata

}
