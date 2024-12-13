package models

type AudioToVideoRequest struct {
	SimliAPIKey       string `json:"simliAPIKey"`
	FaceID            string `json:"faceId"`
	AudioBase64       string `json:"audioBase64"`
	AudioFormat       string `json:"audioFormat"`
	AudioSampleRate   int    `json:"audioSampleRate"`
	AudioChannelCount int    `json:"audioChannelCount"`
}

type AudioToVideoSessionRequest struct {
	FaceID           string `json:"faceId"`
	IsJPG            bool   `json:"isJPG"`
	APIKey           string `json:"apiKey"`
	SyncAudio        bool   `json:"syncAudio"`
	MaxSessionLength int    `json:"maxSessionLength"`
	MaxIdleTime      int    `json:"maxIdleTime"`
}

type WebRTCSessionRequest struct {
	SDP  string `json:"sdp"`
	Type string `json:"type"`
}
