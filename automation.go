package gochimp

import (
	"fmt"
	"io/ioutil"
)

// EnqueueEmail will enqueue the given email address to the emailID trigger of the configured workflow.
func (c *Client) EnqueueEmail(workflowID string, emailID string, email string) error {
	params := struct {
		Email string `json:"email_address"`
	}{email}

	resp, err := c.do("POST", fmt.Sprintf("/automations/%s/emails/%s/queue", workflowID, emailID), &params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		return nil
	}

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return err
	}

	// Dang. We should be able to match this based on the "Type", but …
	// {"type":"http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/","title":"Bad Request","status":400,"detail":"You\u2019ve already sent this email to the subscriber.","instance":""}
	if errorResponse.Status == 400 && errorResponse.Detail == ErrThisMailAlreadySentToThisSubscriber.Error() {
		return ErrThisMailAlreadySentToThisSubscriber
	}

	return errorResponse
}
