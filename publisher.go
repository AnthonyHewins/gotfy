package gotfy

import (
	"bytes"
	"context"
	"encoding/base64"
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

// Option is a functional option for configuring a Publisher
type Option func(*Publisher)

// WithHTTPClient sets a custom HTTP client for the publisher
func WithHTTPClient(client *http.Client) Option {
	return func(p *Publisher) {
		p.httpClient = client
	}
}

// WithAuth adds basic authentication to the publisher's headers
func WithAuth(username, password string) Option {
	return func(p *Publisher) {
		auth := username + ":" + password
		encoded := base64.StdEncoding.EncodeToString([]byte(auth))
		p.Headers.Set("Authorization", "Basic "+encoded)
	}
}

// NewPublisher creates a topic publisher for the specified server URL.
// Options can be provided to customize the publisher behavior.
func NewPublisher(server *url.URL, opts ...Option) (*Publisher, error) {
	if server == nil {
		return nil, ErrNoServer
	}

	p := &Publisher{
		server:     server,
		httpClient: http.DefaultClient,
		Headers:    http.Header{"Content-Type": []string{"application/json"}},
	}

	for _, opt := range opts {
		opt(p)
	}

	return p, nil
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

	for name, headers := range t.Headers {
		for _, h := range headers {
			req.Header.Set(name, h)
		}
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
