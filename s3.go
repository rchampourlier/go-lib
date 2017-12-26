package lib

// Install AWS Go SDK with `go get -u github.com/aws/aws-sdk-go`

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3 is a wrapper around AWS S3 SDK.
type S3 struct {
	Bucket string
	Prefix string
}

// NewS3 returns a valid S3 struct. Please use it to
// create a Client and not create it by yourself.
func NewS3(bucket string, prefix string) S3 {
	return S3{
		Bucket: bucket,
		Prefix: prefix,
	}
}

// ListObjects list objects stored in the client's S3 bucket with
// the specified prefix and returns their keys.
func (s3 S3) ListObjects() ([]string, error) {
	objectKeys := make([]string, 0)
	sess := session.Must(session.NewSession())
	awsS3Client := awsS3.New(sess)

	params := &awsS3.ListObjectsInput{
		Bucket: aws.String(s3.Bucket),
		Prefix: nil,
	}

	var contents []*awsS3.Object
	err := awsS3Client.ListObjectsPages(params,
		func(page *awsS3.ListObjectsOutput, lastPage bool) bool {
			for i := range page.Contents {
				item := page.Contents[i]
				contents = append(contents, item)
			}
			return !lastPage
		},
	)
	if err != nil {
		return objectKeys, err
	}

	// Retrieving keys from fetched objects
	for i := range contents {
		item := contents[i]
		objectKeys = append(objectKeys, *item.Key)
	}
	return objectKeys, nil
}

// FetchObject fetches the content of the object specified by its key.
func (s3 S3) FetchObject(key string) ([]byte, error) {
	var content []byte
	sess := session.Must(session.NewSession())
	input := &awsS3.GetObjectInput{
		Bucket: aws.String(s3.Bucket),
		Key:    aws.String(key),
	}
	downloader := s3manager.NewDownloader(sess)

	// Write the contents of S3 Object to a buffer
	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, input)
	if err != nil {
		return content, fmt.Errorf("Failed to download object, %v", err)
	}
	return buff.Bytes(), nil
}

// CreateObject creates a new object on S3 with the specified key and content.
func (s3 S3) CreateObject(key string, content []byte) error {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	r := bytes.NewReader(content)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.Bucket),
		Key:    aws.String(key),
		Body:   r,
	})
	if err != nil {
		return fmt.Errorf("failed to upload object, %v", err)
	}
	return nil
}

// DeleteObject deletes the object with the specified key.
func (s3 S3) DeleteObject(key string) error {
	sess := session.Must(session.NewSession())
	awsS3Client := awsS3.New(sess)

	input := &awsS3.DeleteObjectInput{
		Bucket: aws.String(s3.Bucket),
		Key:    aws.String(key),
	}
	_, err := awsS3Client.DeleteObject(input)
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}
	return nil
}
