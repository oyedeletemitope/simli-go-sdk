package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oyedeletemitope/simli-go-sdk/client"
	"github.com/oyedeletemitope/simli-go-sdk/models"
)

func TestStartAudioToVideoSession(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"session_token": "mock-session-token"
		}`))
	}))
	defer server.Close()

	c := &client.Client{
		BaseURL:    server.URL,
		APIKey:     "n91d0biopn3vb4lszdts1",
		HTTPClient: &http.Client{},
	}

	req := models.AudioToVideoSessionRequest{
		FaceID:           "6ebf0aa7-6fed-443d-a4c6-fd1e3080b215",
		IsJPG:            true,
		APIKey:           "n91d0biopn3vb4lszdts1",
		SyncAudio:        true,
		MaxSessionLength: 3600,
		MaxIdleTime:      600,
	}

	resp, err := c.StartAudioToVideoSession(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.SessionToken != "mock-session-token" {
		t.Errorf("Expected SessionToken to be 'mock-session-token', got %v", resp.SessionToken)
	}
}
