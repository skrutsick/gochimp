package gochimp

import (
	"net/url"
)

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	CheckSubscription(listID string, email string) (*MemberResponse, error)
	Subscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	Unsubscribe(listID, email string) (*MemberResponse, error)
	UpdateSubscription(listID string, email string, status Status, mergeFields map[string]interface{}) (*MemberResponse, error)
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
}
