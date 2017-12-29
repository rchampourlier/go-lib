package golib

// Install AWS Go SDK with `go get -u github.com/aws/aws-sdk-go`

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3 is a wrapper around AWS S3 SDK.
type S3 struct {
	Bucket string
}

// NewS3 returns a valid S3 struct. Please use it to
// create a `S3 struct` and not create it by yourself.
func NewS3(bucket string) S3 {
	return S3{
		Bucket: bucket,
	}
}

// ListObjects list objects stored in the client's S3 bucket with
// the specified `prefix` and returns their keys.
func (s3 S3) ListObjects(prefix string) ([]string, error) {
	objectKeys := make([]string, 0)
	sess := session.Must(session.NewSession())
	awsS3Client := awsS3.New(sess)

	params := &awsS3.ListObjectsInput{
		Bucket: aws.String(s3.Bucket),
		Prefix: aws.String(prefix),
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

// FindLatestInTimestampPrefixedObjects will return the key of the latest
// object by searching the greatest date. For this method to work, objects
// must be prefixed with a timestamp (e.g. ISO-8601-formatted date strings).
//
// If a delimiter is specified (use the `nil` pointer otherwise), the dates
// will be navigated using the groups defined through the delimiter as specified
// by AWS  `GET Bucket (List Objects)` API
// (see https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html).
//
// ### Return values
//
//   - `*string`: a pointer to the found key (`nil` if not found)
//   - `error`
//
// ### NB: limitations
//
//   - This method does not manage pagination yet, so it will not work for
//     list of objects with more than 1000 objects (the default AWS page limit).
//     Using delimiter will help supporting a larger total number of objects, as
//     each delimited group may contain up to 1000 objects.
//
func (s3 S3) FindLatestInTimestampPrefixedObjects(delimiter string) (*string, error) {
	sess := session.Must(session.NewSession())
	awsS3Client := awsS3.New(sess)

	params := &awsS3.ListObjectsInput{
		Bucket:    aws.String(s3.Bucket),
		Delimiter: aws.String(delimiter),
	}

	var findGreatestPrefix func(string) (string, error)
	findGreatestPrefix = func(currentPrefix string) (string, error) {
		commonPrefixes := make([]string, 0)

		params.Prefix = aws.String(currentPrefix)
		err := awsS3Client.ListObjectsPages(params,
			func(page *awsS3.ListObjectsOutput, lastPage bool) bool {
				for _, item := range page.CommonPrefixes {
					commonPrefixes = append(commonPrefixes, *item.Prefix)
				}
				return !lastPage
			},
		)
		if err != nil {
			return "", err
		}

		sort.Strings(commonPrefixes)
		if len(commonPrefixes) > 0 {
			return findGreatestPrefix(commonPrefixes[len(commonPrefixes)-1])
		}
		return currentPrefix, nil
	}
	greatestPrefix, err := findGreatestPrefix("")
	if err != nil {
		return nil, err
	}
	fmt.Printf("greatestPrefix: %s\n", greatestPrefix)

	objectKeys, err := s3.ListObjects(greatestPrefix)
	if err != nil {
		return nil, err
	}

	sort.Strings(objectKeys)
	foundKey := objectKeys[len(objectKeys)-1]
	return &foundKey, nil
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
