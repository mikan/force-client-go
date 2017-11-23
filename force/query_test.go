package force

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
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
			w.Write([]byte("{\"totalSize\":0,\"done\":\"true\"}"))
		},
	))
	defer ts.Close()
	client, _ := NewClient(ts.URL, UnitTest, sampleAPIVer, nil)
	ctx := context.WithValue(context.Background(), "location", ts.URL+testAuthEndpoint)
	err := client.Login(ctx, &Credential{"xxx", "xxx", "xxx", "xxx", "xxx"})
	if err != nil {
		t.Fatalf("failed to execute Login(): %v", err)
	}
	var out interface{}
	_, err = client.Query(ctx, "SELECT Id FROM Contact", &out)
	if err != nil {
		t.Fatalf("unexpected error cought: %v", err)
	}
}
