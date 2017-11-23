package force

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, testAuthEndpoint) {
				loginResponse(w)
				return
			}
			if r.Method != http.MethodGet {
				t.Fatalf("must use %s method", http.MethodGet)
			}
			w.Header().Set("content-Type", "application/json")
			response, _ := json.Marshal([]byte("{}"))
			w.Write(response)
		},
	))
	defer ts.Close()
	client, _ := NewClient(UnitTest, sampleAPIVer, nil)
	ctx := context.WithValue(context.Background(), "location", ts.URL+testAuthEndpoint)
	err := client.Login(ctx, &Credential{"xxx", "xxx", "xxx", "xxx", "xxx"})
	if err != nil {
		t.Fatalf("failed to execute Login(): %v", err)
	}
	client.session.InstanceURL = ts.URL
	var res interface{}
	err = client.Read(ctx, "Contact", sampleSObjectID, &res)
	if err != nil {
		t.Fatalf("unexpected error cought: %v", err)
	}
}
