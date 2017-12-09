package force

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Update updates specified record.
func (c *Client) Update(ctx context.Context, sObjectName, id string, v interface{}) error {
	if len(sObjectName) == 0 {
		return errors.New("missing sObjectName")
	}
	if len(id) == 0 {
		return errors.New("missing sObject ID")
	}
	if v == nil {
		return errors.New("nothing data to update")
	}

	// Marshal struct
	var body []byte
	switch v.(type) {
	case string:
		body = []byte((v).(string))
	default:
		jsonBody, err := json.Marshal(v)
		if err != nil {
			return err
		}
		body = jsonBody
	}

	// Build request
	req, err := c.newRequest(ctx, http.MethodPatch, sObjectResource+sObjectName+"/"+id, bytes.NewReader(body))
	if err != nil {
		c.Logger.Printf("failed to create the request: %v", err)
		return err
	}

	// Execute PATCH
	c.Logger.Printf("%s %v\n", http.MethodPatch, req.URL)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return err
	}

	// Check status
	switch res.StatusCode {
	case http.StatusOK:
		fallthrough
	case http.StatusNoContent:
		return nil
	default:
		return errors.New(fmt.Sprintf("failed to update: %d %s", res.StatusCode, res.Status))
	}
}
