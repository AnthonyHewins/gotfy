package gotfy

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestWithAuth(t *testing.T) {
	username := "testuser"
	password := "testpass"
	expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	// Create a test server that checks for the Authorization header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != expectedAuth {
			t.Errorf("Expected Authorization header %q, got %q", expectedAuth, auth)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check Content-Type header is also present
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type header %q, got %q", "application/json", contentType)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"test123","time":1234567890,"expires":1234567890,"event":"message"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	// Create publisher with authentication
	publisher, err := NewPublisher(serverURL, WithAuth(username, password))
	if err != nil {
		t.Fatalf("Failed to create publisher: %v", err)
	}

	// Send a test message
	msg := &Message{
		Topic:   "test-topic",
		Message: "test message",
	}

	resp, err := publisher.SendMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if resp.ID != "test123" {
		t.Errorf("Expected response ID %q, got %q", "test123", resp.ID)
	}
}

func TestWithAuthAndHTTPClient(t *testing.T) {
	username := "testuser"
	password := "testpass"
	expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != expectedAuth {
			t.Errorf("Expected Authorization header %q, got %q", expectedAuth, auth)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"test456","time":1234567890,"expires":1234567890,"event":"message"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	// Create custom HTTP client
	customClient := &http.Client{}

	// Create publisher with both custom HTTP client and authentication
	publisher, err := NewPublisher(
		serverURL,
		WithHTTPClient(customClient),
		WithAuth(username, password),
	)
	if err != nil {
		t.Fatalf("Failed to create publisher: %v", err)
	}

	// Verify the custom client was set
	if publisher.httpClient != customClient {
		t.Error("Expected custom HTTP client to be set")
	}

	// Send a test message
	msg := &Message{
		Topic:   "test-topic",
		Message: "test message",
	}

	resp, err := publisher.SendMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if resp.ID != "test456" {
		t.Errorf("Expected response ID %q, got %q", "test456", resp.ID)
	}
}

func TestWithoutAuth(t *testing.T) {
	// Create a test server that should NOT receive an Authorization header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "" {
			t.Errorf("Expected no Authorization header, got %q", auth)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"test789","time":1234567890,"expires":1234567890,"event":"message"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	// Create publisher without authentication
	publisher, err := NewPublisher(serverURL)
	if err != nil {
		t.Fatalf("Failed to create publisher: %v", err)
	}

	// Send a test message
	msg := &Message{
		Topic:   "test-topic",
		Message: "test message",
	}

	resp, err := publisher.SendMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if resp.ID != "test789" {
		t.Errorf("Expected response ID %q, got %q", "test789", resp.ID)
	}
}

func TestNewPublisher_NilServer(t *testing.T) {
	_, err := NewPublisher(nil)
	if err != ErrNoServer {
		t.Errorf("Expected ErrNoServer, got %v", err)
	}
}

func TestWithHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"test999","time":1234567890,"expires":1234567890,"event":"message"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	customClient := &http.Client{}

	publisher, err := NewPublisher(serverURL, WithHTTPClient(customClient))
	if err != nil {
		t.Fatalf("Failed to create publisher: %v", err)
	}

	if publisher.httpClient != customClient {
		t.Error("Expected custom HTTP client to be set")
	}

	// Verify it works
	msg := &Message{
		Topic:   "test-topic",
		Message: "test message",
	}

	resp, err := publisher.SendMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	if resp.ID != "test999" {
		t.Errorf("Expected response ID %q, got %q", "test999", resp.ID)
	}
}

func TestDefaultHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"default","time":1234567890,"expires":1234567890,"event":"message"}`))
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	publisher, err := NewPublisher(serverURL)
	if err != nil {
		t.Fatalf("Failed to create publisher: %v", err)
	}

	if publisher.httpClient != http.DefaultClient {
		t.Error("Expected default HTTP client to be set")
	}
}
