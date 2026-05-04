package gotfy

import (
	"encoding/json"
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
	ClickURL string         `json:"click,omitempty"`    // Website opened when notification is clicked
	IconURL  string         `json:"icon,omitempty"`     // URL to use as notification icon
	Delay    time.Duration  `json:"delay,omitempty"`    // Duration to delay delivery
	Email    string         `json:"email,omitempty"`    // E-mail address for e-mail notifications
	Call     string         `json:"call,omitempty"`     // Phone number to use for voice call

	AttachURLFilename string `json:"filename,omitempty"`  // File name of the attachment
	AttachURL         string `json:"attachurl,omitempty"` // URL of an attachment
}

func (m *Message) MarshalJSON() ([]byte, error) {
	type wire struct {
		Topic    string         `json:"topic"`
		Message  string         `json:"message,omitempty"`
		Title    string         `json:"title,omitempty"`
		Tags     []string       `json:"tags,omitempty"`
		Priority Priority       `json:"priority,omitempty"`
		Actions  []ActionButton `json:"actions,omitempty"`
		ClickURL string         `json:"click,omitempty"`
		IconURL  string         `json:"icon,omitempty"`
		Delay    string         `json:"delay,omitempty"`
		Email    string         `json:"email,omitempty"`
		Call     string         `json:"call,omitempty"`

		AttachURLFilename string `json:"filename,omitempty"`
		AttachURL         string `json:"attachurl,omitempty"`
	}

	priority := m.Priority
	if priority < 0 {
		priority = 0
	}

	var delay string
	if m.Delay > 0 {
		delay = m.Delay.String()
	}

	return json.Marshal(wire{
		Topic:    m.Topic,
		Message:  m.Message,
		Title:    m.Title,
		Tags:     m.Tags,
		Priority: priority,
		Actions:  m.Actions,
		ClickURL: m.ClickURL,
		IconURL:  m.IconURL,
		Delay:    delay,
		Email:    m.Email,
		Call:     m.Call,

		AttachURLFilename: m.AttachURLFilename,
		AttachURL:         m.AttachURL,
	})
}
