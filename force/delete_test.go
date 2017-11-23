package force

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, testAuthEndpoint) {
				loginResponse(w)
				return
			}
			if r.Method != http.MethodDelete {
				t.Fatalf("must use %s method", http.MethodDelete)
			}
			w.WriteHeader(http.StatusNoContent)
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
	err = client.Delete(ctx, "Contact", sampleSObjectID)
	if err != nil {
		t.Fatalf("unexpected error cought: %v", err)
	}
}
