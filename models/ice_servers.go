package models

type ICERequest struct {
	APIKey string `json:"apiKey"`
}

type ICEResponse struct {
	URLs []string `json:"urls"`
}
