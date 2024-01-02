package orgfaclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/dto"
	requestbuidler "github.com/dr3dnought/request_builder"
)

type Client struct {
	httpClient *http.Client
	reqBuilder *requestbuidler.RequestBuilder
}

func NewClient(httpClient *http.Client, url string) *Client {
	return &Client{
		httpClient: httpClient,
		reqBuilder: requestbuidler.New(url),
	}
}

type SearchGroupInput struct {
	groupTerm string
}

func (c *Client) SearchGroups(input SearchGroupInput) ([]dto.Group, error) {
	path := fmt.Sprintf("search?type=group&term=%s", input.groupTerm)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, err
	}

	var groups []dto.Group

	if err := json.NewDecoder(res.Body).Decode(&groups); err != nil {
		return nil, err
	}

	return groups, nil
}