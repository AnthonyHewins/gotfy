package gotfy

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
)

var (
	ErrNoServer = errors.New("server is nil")
	ErrNoTopic  = errors.New("topic is nil")
)

// Publisher creates messages for topics
type Publisher struct {
	server     *url.URL
	httpClient *http.Client

	Headers http.Header
}

// NewPublisher creates a topic publisher for the specified server URL,
// and uses the supplied HTTP client to resolve the request
func NewPublisher(server *url.URL, httpClient *http.Client) (*Publisher, error) {
	if server == nil {
		return nil, ErrNoServer
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Publisher{
		server:     server,
		httpClient: httpClient,
		Headers:    http.Header{"Content-Type": []string{"application/json"}},
	}, nil
}

func (t *Publisher) SendMessage(ctx context.Context, m *Message) (*PublishResp, error) {
	buf, err := json.MarshalContext(ctx, m)
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
	defer resp.Body.Close()

	if s := resp.StatusCode; s < 200 || s >= 300 {
		return nil, fmt.Errorf("non-200 http response code from server: %d", s)
	}

	var pubResp PublishResp
	if err = json.NewDecoder(resp.Body).DecodeContext(ctx, &pubResp); err != nil {
		return nil, err
	}

	return &pubResp, nil
}
