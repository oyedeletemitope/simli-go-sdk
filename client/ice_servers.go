package client

import "github.com/oyedeletemitope/simli-go-sdk/models"

func (c *Client) GetICEServers(request models.ICERequest) (*models.ICEResponse, error) {
	var response models.ICEResponse
	err := c.post("/getIceServers", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
