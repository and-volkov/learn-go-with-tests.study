package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":       true,
		"http://facebook.com":     true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}
