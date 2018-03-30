package time_test

import (
	"testing"
	"time"

	timeutils "github.com/rchampourlier/golib/time"
)

func TestMsToTime(t *testing.T) {
	result, err := timeutils.MsToTime(1522445543669)
	if err != nil {
		t.Fatal(err)
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		t.Fatal(err)
	}
	expectedTime := time.Date(2018, time.March, 30, 23, 32, 23, 669000000, loc)
	if expectedTime != result {
		t.Errorf("expected %s, got %s", expectedTime, result)
	}

}
