package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oyedeletemitope/simli-go-sdk/client"
	"github.com/oyedeletemitope/simli-go-sdk/models"
)

func TestAudioToVideoBasic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"hls_url": "http://mock-hls-url.com",
			"mp4_url": "http://mock-mp4-url.com"
		}`))
	}))
	defer server.Close()

	c := &client.Client{
		BaseURL:    server.URL,
		APIKey:     "",
		HTTPClient: &http.Client{},
	}

	req := models.AudioToVideoRequest{
		SimliAPIKey:       "",
		FaceID:            "",
		AudioBase64:       "mock-audio-data",
		AudioFormat:       "pcm16",
		AudioSampleRate:   16000,
		AudioChannelCount: 1,
	}

	resp, err := c.AudioToVideo(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.HLSURL != "http://mock-hls-url.com" {
		t.Errorf("Expected HLSURL to be 'http://mock-hls-url.com', got %v", resp.HLSURL)
	}
	if resp.MP4URL != "http://mock-mp4-url.com" {
		t.Errorf("Expected MP4URL to be 'http://mock-mp4-url.com', got %v", resp.MP4URL)
	}
}
