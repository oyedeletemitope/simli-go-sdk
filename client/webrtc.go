package client

import "github.com/oyedeletemitope/simli-go-sdk/models"

func (c *Client) StartWebRTCSession(request models.WebRTCSessionRequest) (*models.WebRTCSessionResponse, error) {
	var response models.WebRTCSessionResponse
	err := c.post("/StartWebRTCSession", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
