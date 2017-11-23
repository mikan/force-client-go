package sobject

import (
	"net/http/httptest"

	"context"

	"github.com/mikan/force-client-go/force"
)

// base_test provides test helper functions.

func makeTestCredential() *force.Credential {
	return &force.Credential{ClientID: "123", ClientSecret: "123", Username: "123", Password: "123", APIToken: "123"}
}

func makeTestContext() *context.Context {
	ctx := context.Background()
	return &ctx
}

func makeTestClient(ts *httptest.Server) *force.Client {
	client, err := force.NewClient(force.UnitTest, force.DefaultVersion, nil)
	if err != nil {
		panic("failed to setup client!")
	}
	client.Session(&force.SessionID{AccessToken: "000000000000000000", TokenType: "Bearer", InstanceURL: ts.URL})
	return client
}
