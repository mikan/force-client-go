package force

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Env defines force.com target environments
type Env int

// Available environments
const (
	Production Env = iota
	Sandbox
	UnitTest
)

const DefaultVersion = "40.0"

type Client struct {
	Version    string
	Env        Env
	HttpClient *http.Client
	Logger     *log.Logger
	session    *SessionID
}

const sObjectResource = "/services/data/v%s/sobjects/"

var versionMatcher = regexp.MustCompile("[0-9]+\\.[0-9]+")

// NewClient creates the new force.com client.
func NewClient(env Env, version string, logger *log.Logger) (*Client, error) {
	if env < 0 {
		return nil, errors.New("invalid env: " + string(env))
	}
	if !versionMatcher.MatchString(version) {
		return nil, errors.New("invalid version number: " + version)
	}
	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	client := &Client{version, env, &http.Client{}, logger, nil}
	return client, nil
}

// Session sets already authorized session to the current client.
func (c *Client) Session(session *SessionID) {
	c.session = session
}

func (c *Client) newRequest(ctx context.Context, method, resource string, body io.Reader) (*http.Request, error) {
	location := c.session.InstanceURL + fmt.Sprintf(resource, c.Version)
	req, err := http.NewRequest(method, location, body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", c.session.TokenType+" "+c.session.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func readBody(resp *http.Response) (string, error) {
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

func decode(body string, out interface{}) error {
	return json.Unmarshal([]byte(body), out)
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out) // better than ReadAll + Unmarshal
}
