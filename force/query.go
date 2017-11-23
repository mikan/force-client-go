package force

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"reflect"
)

const queryResource = "/services/data/v%s/query"

// Query executes SOQL and store result to specified struct.
func (c *Client) Query(ctx context.Context, soql string, out interface{}) (string, error) {
	if out == nil {
		return "", errors.New("missing out argument")
	}
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return "", errors.New("out argument isn't a pointer")
	}

	// Build request
	req, err := c.newRequest(ctx, http.MethodGet, queryResource, nil)
	if err != nil {
		c.Logger.Printf("failed to create the request: %v", err)
		return "", err
	}
	values := url.Values{}
	values.Add("q", soql)
	req.URL.RawQuery = values.Encode()

	// Execute GET
	c.Logger.Printf("%s %v\n", http.MethodGet, req.URL)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return "", err
	}

	// Decode response
	if err := decodeBody(res, out); err != nil {
		c.Logger.Printf("failed to decode the response: %v", err)
		return "", err
	}
	return res.Header.Get("nextRecordsUrl"), nil
}
