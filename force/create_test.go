package force

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, testAuthEndpoint) {
				loginResponse(w)
				return
			}
			if r.Method != http.MethodPost {
				t.Fatalf("must use %s method", http.MethodPost)
			}
			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal("failed to read body")
			}
			body := string(raw)
			if !strings.Contains(body, "FirstName") {
				t.Fatal("FirstName not found")
			}
			if !strings.Contains(body, "LastName") {
				t.Fatal("LastName not found")
			}
			w.Header().Set("content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			response, _ := json.Marshal(createResponse{sampleSObjectID, nil, true})
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
	client.Create(ctx, "Contact", "{\"FirstName\":\"Test\",\"LastName\":\"User\"}")
}
