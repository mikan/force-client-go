package sobject

import (
	"net/http/httptest"

	"context"
	"time"

	"github.com/mikan/force-client-go/force"
)

// base_test provides test helper functions.

func makeTestCredential() *force.Credential {
	return &force.Credential{ClientID: "123", ClientSecret: "123", Username: "123", Password: "123", APIToken: "123"}
}

func makeTestContext() *context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return &ctx
}

func makeTestClient(ts *httptest.Server) *force.Client {
	client, err := force.NewClient(ts.URL, force.UnitTest, force.DefaultVersion, nil)
	if err != nil {
		panic("failed to setup client!")
	}
	client.Session(&force.SessionID{AccessToken: "xxx", TokenType: "Bearer"})
	return client
}
