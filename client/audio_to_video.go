package client

import "github.com/oyedeletemitope/simli-go-sdk/models"

func (c *Client) AudioToVideo(request models.AudioToVideoRequest) (*models.VideoResponse, error) {
	var response models.VideoResponse
	err := c.post("/audioToVideoStream", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
