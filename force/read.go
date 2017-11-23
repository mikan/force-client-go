package force

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

// Read reads specified record.
func (c *Client) Read(ctx context.Context, sObjectName, id string, out interface{}) error {
	if len(sObjectName) == 0 {
		return errors.New("missing sObjectName")
	}
	if len(id) == 0 {
		return errors.New("missing sObject ID")
	}
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return errors.New("out argument isn't a pointer")
	}

	// Build request
	req, err := c.newRequest(ctx, http.MethodGet, sObjectResource+sObjectName+"/"+id, nil)
	if err != nil {
		c.Logger.Printf("failed to create the request: %v", err)
		return err
	}

	// Execute GET
	c.Logger.Printf("%s %v\n", http.MethodGet, req.URL)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return err
	}

	// Check status
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("failed to get: %d %s", res.StatusCode, res.Status))
	}

	// Decode response
	if err := decodeBody(res, &out); err != nil {
		c.Logger.Printf("failed to decode the response: %v", err)
		return err
	}
	return nil
}
