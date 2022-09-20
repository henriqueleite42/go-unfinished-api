package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetS3() *session.Session {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	return session.Must(session.NewSession(&aws.Config{
		Endpoint: &awsEndpoint,
	}))
}
