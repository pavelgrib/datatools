package scraper

import (
	"http"
	"testing"
)

var (
	testUrl = []string{"www.google.com", "www.ebates.com", "www.ebay.com"}
)

func mockGetResponse(url, agent) *http.Response {

}

func TestSetupUser(t *testing.T) {
	u := setupClient()
	t.Logf("%v", u)
}

func TestGetResponse(t *testing.T) {
	res := getResponse(testUrl[0], "ipad")
	if res == nil {
		t.Errorf("response is nil")
	}
}
