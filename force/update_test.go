package force

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, testAuthEndpoint) {
				loginResponse(w)
				return
			}
			if r.Method != http.MethodPatch {
				t.Fatalf("must use %s method", http.MethodPatch)
			}
		},
	))
	defer ts.Close()
	client, _ := NewClient(ts.URL, UnitTest, sampleAPIVer, nil)
	ctx := context.WithValue(context.Background(), "location", ts.URL+testAuthEndpoint)
	err := client.Login(ctx, &Credential{"xxx", "xxx", "xxx", "xxx", "xxx"})
	if err != nil {
		t.Fatalf("failed to execute Login(): %v", err)
	}
	err = client.Update(ctx, "Contact", sampleSObjectID, "{\"FirstName\":\"Test2\"}")
	if err != nil {
		t.Fatalf("unexpected error cought: %v", err)
	}
}
