package golib_test

// This test does E2E testing and will connect and perform actions
// on AWS S3. AWS credentials must be loaded in the environment.
// You can use `.env.example` as a template of a way to load the
// credentials.

import (
	"os"
	"testing"

	"github.com/rchampourlier/golib"
)

func buildS3() golib.S3 {
	bucket := os.Getenv("AWS_BUCKET")
	return golib.NewS3(bucket, "")
}

func countObjects(s3 golib.S3) (int, error) {
	objectKeys, err := s3.ListObjects()
	if err != nil {
		return -1, err
	}
	return len(objectKeys), nil
}

func createObject(s3 golib.S3) error {
	err := s3.CreateObject("test_object", []byte("test_object content"))
	return err
}

func deleteObject(s3 golib.S3) error {
	err := s3.DeleteObject("test_object")
	return err
}

func handleError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
}

func TestListObjectsWithEmptyBucket(t *testing.T) {
	s3 := buildS3()
	count, err := countObjects(s3)
	handleError(err, t)
	if count != 0 {
		t.Errorf("expected bucket to be empty!")
	}
}

func TestCreateListAndDelete(t *testing.T) {
	s3 := buildS3()
	err := createObject(s3)
	handleError(err, t)
	count, err := countObjects(s3)
	handleError(err, t)
	if count != 1 {
		t.Errorf("expected bucket to contain 1 object, got %d", count)
	}
	deleteObject(s3)
}
