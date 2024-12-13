package models

type TextToVideoRequest struct {
	TTSAPIKey   string      `json:"ttsAPIKey"`
	SimliAPIKey string      `json:"simliAPIKey"`
	FaceID      string      `json:"faceId"`
	RequestBody interface{} `json:"requestBody"`
}
