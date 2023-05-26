package gotfy

import (
	"errors"
	"net/http"
	"net/url"
)

var (
	ErrNoTopic = errors.New("topic is nil")
)

// TopicPublisher creates messages for topics
type TopicPublisher struct {
	httpClient *http.Client
	server     *url.URL
}

// NewTopicPublisher creates a topic publisher for the specified server URL,
// and uses the supplied HTTP client to resolve the request
func NewTopicPublisher(server *url.URL, httpClient *http.Client) (*TopicPublisher, error) {
	if server == nil {
		return nil, ErrNoTopic
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &TopicPublisher{
		httpClient: httpClient,
		server:     server,
	}, nil
}
