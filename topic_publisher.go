package gotfy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	"log/slog"
)

var (
	ErrNoServer = errors.New("server is nil")
	ErrNoTopic  = errors.New("topic is nil")
)

// TopicPublisher creates messages for topics
type TopicPublisher struct {
	logger *slog.Logger

	server     *url.URL
	httpClient *http.Client
}

// NewTopicPublisher creates a topic publisher for the specified server URL,
// and uses the supplied HTTP client to resolve the request. Uses the golang
// slog package to log to; if you want to skip all logs supply slog.Logger{}
// with a blank handler, and the publisher will do a no-op
func NewTopicPublisher(slogger *slog.Logger, server *url.URL, httpClient *http.Client) (*TopicPublisher, error) {
	if slogger == nil {
		// if no logger is passed, ignore absolutely everything
		slogger = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.Level(math.MaxInt)}))
	}

	if server == nil {
		return nil, ErrNoServer
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &TopicPublisher{
		server:     server,
		httpClient: httpClient,
		logger:     slogger,
	}, nil
}

func (t *TopicPublisher) SendMessage(ctx context.Context, m *Message) (*PublishResp, error) {
	l := t.logger.With("message", m)

	l.DebugContext(ctx, "marshaling NTFY message")
	buf, err := json.Marshal(m)
	if err != nil {
		l.ErrorContext(ctx, "failed marshal", "err", err)
		return nil, err
	}

	l.DebugContext(ctx, "finished marshal, creating request struct", "server", t.server)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.server.String(), bytes.NewReader(buf))
	if err != nil {
		l.ErrorContext(ctx, "failed creating HTTP request", "err", err)
		return nil, err
	}

	l.DebugContext(ctx, "finished creation of request struct, prepping HTTP call", "req", req)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		l.ErrorContext(ctx, "failed HTTP call", "http client", t.httpClient, "req", req, "err", err)
		return nil, err
	}

	code := resp.StatusCode
	l.DebugContext(ctx, "finished HTTP call, reading response body", "status code", code)
	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		l.ErrorContext(ctx, "failed reading response body", "status code", code, "err", err)
		return nil, err
	}

	if s := resp.StatusCode; s < 200 || s >= 300 {
		l.ErrorContext(ctx, "bad HTTP response code from server", "response body", string(buf), "status code", code)
		return nil, fmt.Errorf("bad http response from server: %d", code)
	}

	l.DebugContext(ctx, "unmarshaling response body")
	var pubResp PublishResp
	if err = json.Unmarshal(buf, &pubResp); err != nil {
		l.ErrorContext(ctx, "failed unmarshaling response body", "response body", string(buf), "status code", code)
		return nil, err
	}

	l.DebugContext(ctx, "finished unmarshal")
	return &pubResp, nil
}
