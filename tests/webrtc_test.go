package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oyedeletemitope/simli-go-sdk/client"
	"github.com/oyedeletemitope/simli-go-sdk/models"
)

func TestStartWebRTCSession(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"sdp": "mock-sdp",
			"type": "answer"
		}`))
	}))
	defer server.Close()

	c := &client.Client{
		BaseURL:    server.URL,
		APIKey:     "n91d0biopn3vb4lszdts1",
		HTTPClient: &http.Client{},
	}

	req := models.WebRTCSessionRequest{
		SDP:  "mock-sdp",
		Type: "offer",
	}

	resp, err := c.StartWebRTCSession(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.SDP != "mock-sdp" {
		t.Errorf("Expected SDP to be 'mock-sdp', got %v", resp.SDP)
	}
	if resp.Type != "answer" {
		t.Errorf("Expected Type to be 'answer', got %v", resp.Type)
	}
}
