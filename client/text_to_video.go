package client

import "github.com/oyedeletemitope/simli-go-sdk/models"

func (c *Client) TextToVideo(request models.TextToVideoRequest) (*models.VideoResponse, error) {
	var response models.VideoResponse
	err := c.post("/textToVideoStream", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
