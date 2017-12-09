package force

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type createResponse struct {
	Id      string
	Errors  []string
	Success bool
}

// Create creates record. Results are SObject ID and error.
func (c *Client) Create(ctx context.Context, sObjectName string, v interface{}) (string, error) {
	if len(sObjectName) == 0 {
		return "", errors.New("missing sObjectName")
	}
	if v == nil {
		return "", errors.New("nothing data to create")
	}

	// Marshal struct
	var body []byte
	switch v.(type) {
	case string:
		body = []byte((v).(string))
	default:
		jsonBody, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		body = jsonBody
	}

	// Build request
	req, err := c.newRequest(ctx, http.MethodPost, sObjectResource+sObjectName+"/", bytes.NewReader(body))
	if err != nil {
		c.Logger.Printf("failed to create the request: %v", err)
		return "", err
	}

	// Execute POST
	c.Logger.Printf("%s %v\n", http.MethodPost, req.URL)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return "", err
	}

	// Check status
	if res.StatusCode != http.StatusCreated {
		return "", errors.New(fmt.Sprintf("failed to create: %d %s", res.StatusCode, res.Status))
	}

	// Decode response
	var out createResponse
	if err := decodeBody(res, &out); err != nil {
		c.Logger.Printf("failed to decode the response: %v", err)
		return "", err
	}
	if out.Success {
		return out.Id, nil
	} else {
		return "", errors.New(strings.Join(out.Errors, ","))
	}
}
