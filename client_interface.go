package gochimp

import (
	"errors"
	"net/url"
)

var ErrThisMailAlreadySentToThisSubscriber = errors.New("You’ve already sent this email to the subscriber.")

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	CheckSubscription(listID string, email string) (*MemberResponse, error)
	Subscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	Unsubscribe(listID, email string) (*MemberResponse, error)
	UpdateSubscription(listID string, email string, status Status, mergeFields map[string]interface{}) (*MemberResponse, error)
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
	EnqueueEmail(workflowID string, emailID string, email string) error
}
