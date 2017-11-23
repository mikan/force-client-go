package force

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

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
