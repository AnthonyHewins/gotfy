package gotfy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrNoServer = errors.New("server is nil")
	ErrNoTopic  = errors.New("topic is nil")
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
		return nil, ErrNoServer
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &TopicPublisher{
		server:     server,
		httpClient: httpClient,
	}, nil
}

func (t *TopicPublisher) SendMessage(ctx context.Context, m *Message) (*PublishResp, error) {
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.server.String(), bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	code := resp.StatusCode
	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if s := resp.StatusCode; s < 200 || s >= 300 {
		return nil, fmt.Errorf("bad http response from server: %d", code)
	}

	var pubResp PublishResp
	if err = json.Unmarshal(buf, &pubResp); err != nil {
		return nil, err
	}

	return &pubResp, nil
}
