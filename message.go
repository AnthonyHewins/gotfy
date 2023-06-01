package gotfy

import (
	"encoding/json"
	"fmt"
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

func (m *Message) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(m.Topic)
	if err != nil {
		return nil, err
	}
	buf = []byte(fmt.Sprintf(`{"topic":%s`, buf))

	if x := m.Message; x != "" {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"message":%s`, mm)...)
	}

	if x := m.Title; x != "" {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"title":%s`, mm)...)
	}

	if x := m.Tags; len(x) > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"tags":%s`, mm)...)
	}

	if x := m.Priority; x > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"priority":%s`, mm)...)
	}

	if x := m.Actions; len(x) > 0 {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"actions":%s`, mm)...)
	}

	if x := m.ClickURL; x != nil {
		mm, err := json.Marshal(x.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"click":%s`, mm)...)
	}

	if x := m.AttachURL; x != nil {
		mm, err := json.Marshal(x.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"attachurl":%s`, mm)...)
	}

	if x := m.IconURL; x != nil {
		mm, err := json.Marshal(x.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"icon":%s`, mm)...)
	}

	if x := m.Delay; x > 0 {
		mm, err := json.Marshal(x.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"delay":%s`, mm)...)
	}

	if x := m.Email; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"email":%s`, mm)...)
	}

	if x := m.Call; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"call":%s`, mm)...)
	}

	if x := m.AttachURLFilename; x != `` {
		mm, err := json.Marshal(x)
		if err != nil {
			return nil, err
		}
		buf = append(buf, fmt.Sprintf(`,"filename":%s`, mm)...)
	}

	return append(buf, '}'), nil
}
