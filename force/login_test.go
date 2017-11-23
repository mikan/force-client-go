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

const testAuthEndpoint = "/auth"

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Fatalf("must use %s method", http.MethodPost)
			}
			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal("failed to read body")
			}
			body := string(raw)
			if !strings.Contains(body, "client_id") {
				t.Fatal("client_id not found")
			}
			if !strings.Contains(body, "client_secret") {
				t.Fatal("client_secret not found")
			}
			if !strings.Contains(body, "username") {
				t.Fatal("username not found")
			}
			if !strings.Contains(body, "password") {
				t.Fatal("password not found")
			}
			loginResponse(w)
		},
	))
	defer ts.Close()
	client, _ := NewClient(ts.URL, UnitTest, sampleAPIVer, nil)
	ctx := context.WithValue(context.Background(), "location", ts.URL+testAuthEndpoint)
	err := client.Login(ctx, &Credential{"xxx", "xxx", "xxx", "xxx", "xxx"})
	if err != nil {
		t.Fatalf("failed to execute Login(): %v", err)
	}
}

func loginResponse(w http.ResponseWriter) {
	w.Header().Set("content-Type", "application/json")
	response, _ := json.Marshal(SessionID{"xxx", "xxx", "xxx", "xxx", "xxx", "xxx", "", ""})
	w.Write(response)
}
