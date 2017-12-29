package golib_test

// This test does E2E testing and will connect and perform actions
// on AWS S3. AWS credentials must be loaded in the environment.
// You can use `.env.example` as a template of a way to load the
// credentials.

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rchampourlier/golib"
)

var bucket = os.Getenv("AWS_BUCKET")
var s3 = golib.NewS3(bucket)

func countObjects(prefix string) (int, error) {
	objectKeys, err := s3.ListObjects(prefix)
	if err != nil {
		return -1, err
	}
	return len(objectKeys), nil
}

func createObject() (string, error) {
	key := "test_object"
	err := s3.CreateObject(key, []byte("test_object content"))
	return key, err
}

func createTimestampPrefixedObject(year int, month time.Month, day int) (string, error) {
	key := fmt.Sprintf("%d/%d/%d", year, month, day)
	content := fmt.Sprintf("test_object %s", key)
	err := s3.CreateObject(key, []byte(content))
	return key, err
}

func deleteObject(key string) error {
	err := s3.DeleteObject(key)
	return err
}

func handleError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
}

func TestListObjectsWithEmptyBucket(t *testing.T) {
	count, err := countObjects("")
	handleError(err, t)
	if count != 0 {
		t.Errorf("expected bucket to be empty!")
	}
}

func TestListObjectsWithPrefix(t *testing.T) {
	_, err := createObject()
	handleError(err, t)
	count, err := countObjects("test_")
	handleError(err, t)
	if count != 1 {
		t.Errorf("expected to find 1 object matching the `test_` prefix, got %d", count)
	}
	count, err = countObjects("nope")
	handleError(err, t)
	if count != 0 {
		t.Errorf("expected to find no objects matching the `none` prefix, got %d", count)
	}
}

func TestCreateListAndDelete(t *testing.T) {
	key, err := createObject()
	handleError(err, t)
	count, err := countObjects("")
	handleError(err, t)
	if count != 1 {
		t.Errorf("expected bucket to contain 1 object, got %d", count)
	}
	deleteObject(key)
}

func TestFindLatestInTimestampPrefixedObjects(t *testing.T) {
	testObjectKeys := make([]string, 0)
	type loopParam struct {
		year  int
		month time.Month
		day   int
	}
	loopParams := []loopParam{
		loopParam{2016, time.January, 1},
		loopParam{2017, time.January, 1},
		loopParam{2017, time.February, 1},
		loopParam{2017, time.February, 2},
	}
	for _, loopParam := range loopParams {
		key, err := createTimestampPrefixedObject(loopParam.year, loopParam.month, loopParam.day)
		handleError(err, t)
		testObjectKeys = append(testObjectKeys, key)
	}

	foundKey, err := s3.FindLatestInTimestampPrefixedObjects("/")
	handleError(err, t)

	expectedKey := "2017/2/2"
	if foundKey == nil {
		t.Errorf("expected to find object with key `%s` but did not find anything", expectedKey)
	} else {
		if *foundKey != expectedKey {
			t.Errorf("expected to find object with key `%s`, got `%s`", expectedKey, *foundKey)
		}
	}

	for _, key := range testObjectKeys {
		deleteObject(key)
	}
}
