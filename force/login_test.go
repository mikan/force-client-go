package force

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Fatal("must use POST method")
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
			w.Header().Set("content-Type", "application/json;charset=UTF-8")
			w.Write([]byte("{\"access_token\":\"xxx\",\"instance_url\":\"xxx\",\"id\":\"xxx\",\"token_type\":\"xxx\"" +
				",\"issued_at\":\"xxx\",\"signature\":\"xxx\"}"))
		},
	))
	defer ts.Close()
	client, err := NewClient("cs58.salesforce.com", UnitTest, "40.0", nil)
	if err != nil {
		t.Fatal("failed to setup client!")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = context.WithValue(ctx, "location", ts.URL)
	err = client.Login(ctx, &Credential{"xxx", "xxx", "xxx", "xxx", "xxx"})
	if err != nil {
		t.Fatalf("failed to execute Login(): %v", err)
	}
}
