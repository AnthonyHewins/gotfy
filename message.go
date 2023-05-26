package gotfy

import (
	"net/url"
	"time"
)

// Message is a struct you can create from TopicPublisher that
// will publish a message to the specified topic. This method does not allow
// for attaching files to the notification, but it can post a link to an attachment
type Message struct {
	Topic    string         `json:"topic"`              // Target topic name
	Message  string         `json:"message,omitempty"`  // Message body; set to triggered if empty or not passed
	Title    string         `json:"title,omitempty"`    // Message title
	Tags     []string       `json:"tags,omitempty"`     // List of tags that may or not map to emojis
	Priority Priority       `json:"priority,omitempty"` // Message priority with 1=min, 3=default and 5=max
	Actions  []ActionButton `json:"actions,omitempty"`  // Custom user action buttons for notifications
	ClickURL *url.URL       `json:"click,omitempty"`    // Website opened when notification is clicked
	IconURL  *url.URL       `json:"icon,omitempty"`     // URL to use as notification icon
	Delay    time.Duration  `json:"delay,omitempty"`    // Duration to delay delivery
	Email    string         `json:"email,omitempty"`    // E-mail address for e-mail notifications
	Call     string         `json:"call,omitempty"`     // Phone number to use for voice call

	AttachURLFilename string   `json:"filename,omitempty"`  // File name of the attachment
	AttachURL         *url.URL `json:"attachurl,omitempty"` // URL of an attachment
}
