package gotfy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var (
	ErrNoTopic = errors.New("topic is nil")
)

// TopicPublisher creates messages for topics
type TopicPublisher struct {
	server     *url.URL
	httpClient *http.Client
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

func (t *TopicPublisher) SendMessage(ctx context.Context, m *Message) error {
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.server.String(), bytes.NewReader(buf))
	if err != nil {
		return err
	}

	_, err = t.httpClient.Do(req)
	return err
}
