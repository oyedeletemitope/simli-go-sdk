package client

import "github.com/oyedeletemitope/simli-go-sdk/models"

func (c *Client) StartAudioToVideoSession(request models.AudioToVideoSessionRequest) (*models.SessionResponse, error) {
	var response models.SessionResponse
	err := c.post("/startAudioToVideoSession", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
