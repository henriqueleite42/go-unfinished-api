package s3

import (
	"errors"
	"net/http"
	"unfinished-api/src/models/files"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type s3ImageRepository struct {
	Uploader *s3manager.Uploader
}

func NewFilesImageRepository(sess *session.Session) files.ImageRepository {
	return &s3ImageRepository{
		Uploader: s3manager.NewUploader(sess),
	}
}

func (conn *s3ImageRepository) CreateFromURL(URL string) (string, error) {
	response, err := http.Get(URL)
	if err != nil || response.StatusCode != http.StatusOK {
		return "", errors.New("fail to get image")
	}
	defer response.Body.Close()

	result, err := conn.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(""),
		Key:    aws.String(uuid.New().String() + ".jpg"),
		Body:   response.Body,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
