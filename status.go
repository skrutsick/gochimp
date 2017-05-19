package gochimp

type Status string

const (
	DontChange = Status("")

	// Subscribed - This address is on the list and ready to receive email. You can only send campaigns to ‘subscribed’ addresses.
	Subscribed = Status("subscribed")
	// Unsubscribed - This address is on the list and ready to receive email. You can only send campaigns to ‘subscribed’ addresses.
	Unsubscribed = Status("unsubscribed")
	// Pending - This address used to be on the list but isn’t anymore.
	Pending = Status("pending")
	// Cleaned - This address requested to be added with double-opt-in but hasn’t confirmed their subscription yet.
	Cleaned = Status("cleaned")
)
