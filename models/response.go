package models

type VideoResponse struct {
	HLSURL string `json:"hls_url"`
	MP4URL string `json:"mp4_url"`
}

type SessionResponse struct {
	SessionToken string `json:"session_token"`
}

type WebRTCSessionResponse struct {
	SDP  string `json:"sdp"`
	Type string `json:"type"`
}

type StartWebRTCSessionResponse struct {
	SessionToken string      `json:"sessionToken"`
	IceServers   []IceServer `json:"iceServers"`
}
