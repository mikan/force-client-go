package sobject

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var allResponse = `
{
  "totalSize" : 2,
  "done" : true,
  "records" : [ {
    "attributes" : {
      "type" : "Knowledge__kav",
      "url" : "/services/data/v41.0/sobjects/Knowledge__kav/xxx"
    },
    "ArticleCreatedDate" : "2017-09-27T05:21:35.000+0000",
    "Title": "Test 1",
    "Summary": "Summary 1"
  }, {
    "attributes" : {
      "type" : "Knowledge__kav",
      "url" : "/services/data/v41.0/sobjects/Knowledge__kav/yyy"
    },
    "ArticleCreatedDate" : "2017-09-21T09:41:54.000+0000",
    "Title": "Test 2",
    "Summary": "Summary 2"
  } ]
}
`

func TestAllKnowledge(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Fatal("must use GET method")
			}
			if len(r.Header.Get("Authorization")) == 0 {
				t.Fatal("no authorization header")
			}
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.Write([]byte(allResponse))
		},
	))
	defer ts.Close()
	client := makeTestClient(ts)
	set, err := AllKnowledge(*makeTestContext(), client)
	if err != nil {
		t.Fatalf("error occurs: %v", err)
	}
	if len(set) != 2 {
		t.Fatalf("len(set): expected=3 actual=%d", len(set))
	}
}
