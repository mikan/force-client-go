package force

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
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

// Delete deletes specified record.
func (c *Client) Delete(ctx context.Context, sObjectName, id string) error {
	if len(sObjectName) == 0 {
		return errors.New("missing sObjectName")
	}
	if len(id) == 0 {
		return errors.New("missing sObject ID")
	}

	// Build request
	req, err := c.newRequest(ctx, http.MethodDelete, sObjectResource+sObjectName+"/"+id, nil)
	if err != nil {
		c.Logger.Printf("failed to create the request: %v", err)
		return err
	}

	// Execute DELETE
	c.Logger.Printf("%s %v\n", http.MethodDelete, req.URL)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Printf("failed to execute the request: %v", err)
		return err
	}

	// Check status
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	default:
		return errors.New(fmt.Sprintf("failed to delete: %d %s", res.StatusCode, res.Status))
	}
}
