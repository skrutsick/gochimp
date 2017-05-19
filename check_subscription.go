package gochimp

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CheckSubscription ...
func (c *Client) CheckSubscription(listID string, email string) (*MemberResponse, error) {
	// Hash email
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))
	// Make request
	resp, err := c.do(
		"GET",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		nil,
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
