package ruzfaclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient/dto"
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

type GetGroupsInput struct {
	GroupTerm string
}

func (c *Client) GetGroups(input *GetGroupsInput) ([]dto.Group, error) {
	path := fmt.Sprintf("search?type=group&term=%s", input.GroupTerm)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}
	defer res.Body.Close()

	groups := new([]dto.Group)
	err = json.Unmarshal(body, groups)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *groups, nil
}

type GetScheduleInput struct {
	//start request
	GroupId   string 
	StartDate string 
	EndDate   string 
}

func (c *Client) GetSchedule(input *GetScheduleInput) ([]dto.Schedule, error) {
	path := fmt.Sprintf("schedule/group/%s?start=%s&finish=%s&lng=1", input.GroupId, input.StartDate, input.EndDate)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}



	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	//fmt.Print(string(body))
	fmt.Println(res.StatusCode)
	defer res.Body.Close()

	schedule := new([]dto.Schedule)
	err = json.Unmarshal(body, schedule)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}

	return *schedule, nil
}
