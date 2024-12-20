package models

type ICERequest struct {
	APIKey string `json:"apiKey"`
}

type ICEResponse struct {
	URLs []string `json:"urls"`
}

type IceServer struct {
	URLs       []string `json:"urls"`
	Username   string   `json:"username,omitempty"`
	Credential string   `json:"credential,omitempty"`
}
