package force

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Credential struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	APIToken     string `json:"api_token"`
}

type SessionID struct {
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	ID          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
	Error       string `json:"error"`
	ErrorDesc   string `json:"error_description"`
}

const (
	productionLoginURL = "https://login.salesforce.com/services/oauth2/token"
	sandboxLoginURL    = "https://test.salesforce.com/services/oauth2/token"
)

// Login authenticates the credential.
func (c *Client) Login(ctx context.Context, cred *Credential) error {
	// Build request
	loginUrl := c.loginURL(ctx)
	values := url.Values{}
	values.Add("grant_type", "password")
	values.Add("client_id", cred.ClientID)
	values.Add("client_secret", cred.ClientSecret)
	values.Add("username", cred.Username)
	values.Add("password", cred.Password+cred.APIToken)
	c.Logger.Printf("POST %s", loginUrl)
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(values.Encode()))
	if err != nil {
		c.Logger.Printf("failed to create the http request: %v", err)
		return err
	}
	req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execute POST
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			c.Logger.Printf("failed to close body: %v", err)
		}
	}()

	// Decode response
	var session SessionID
	body, err := readBody(res)
	if err != nil {
		c.Logger.Printf("failed to read the response: %v", err)
		return err
	}
	if err := decode(body, &session); err != nil {
		c.Logger.Printf("failed to decode the response: %v", err)
		return err
	}

	// Check response
	if len(session.Error) > 0 {
		msg := fmt.Sprintf("failed to login: %s (%s)", session.Error, session.ErrorDesc)
		c.Logger.Printf(msg)
		return errors.New(msg)
	}
	c.session = &session
	c.Logger.Printf("successfully logged in: " + cred.Username)
	return nil
}

func (c *Client) loginURL(ctx context.Context) string {
	switch c.Env {
	case Production:
		return productionLoginURL
	case Sandbox:
		return sandboxLoginURL
	case UnitTest:
		value, ok := ctx.Value("location").(string)
		if !ok {
			log.Panic("must configure unit test parameter: location")
		}
		return value
	default:
		panic("unsupported environment: " + string(c.Env))
	}
}
