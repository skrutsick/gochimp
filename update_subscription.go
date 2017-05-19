package gochimp

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c *Client) Unsubscribe(listID, email string) (*MemberResponse, error) {
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))
	params := map[string]interface{}{
		"status": Unsubscribed,
	}
	resp, err := c.do(
		"PATCH",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		&params,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		// Unmarshal response into MemberResponse struct
		memberResponse := new(MemberResponse)
		if err := json.Unmarshal(data, memberResponse); err != nil {
			return nil, err
		}
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}

// UpdateSubscription ...
func (c *Client) UpdateSubscription(listID, email string, status Status, mergeFields map[string]interface{}) (*MemberResponse, error) {
	// Hash email
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))
	// Make request
	params := map[string]interface{}{
		"email_address": email,
		"merge_fields":  mergeFields,
	}
	if status != DontChange {
		params["status"] = status
	}
	resp, err := c.do(
		"PUT",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		&params,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		// Unmarshal response into MemberResponse struct
		memberResponse := new(MemberResponse)
		if err := json.Unmarshal(data, memberResponse); err != nil {
			return nil, err
		}
		return memberResponse, nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
