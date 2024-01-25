package ruzfaclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient/models"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	requestbuidler "github.com/dr3dnought/request_builder"
)

type Client struct {
	httpClient *http.Client
	reqBuilder *requestbuidler.RequestBuilder
	Cfg        *config.Config
}

func NewClient(url string, cfg *config.Config) *Client {
	return &Client{
		Cfg:        cfg,
		httpClient: &http.Client{},
		reqBuilder: requestbuidler.New(url),
	}
}

type GetScheduleInput struct {
	Id        string `json:"teacherId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type GetEntitiesInput struct {
	Term string `json:"group"`
}

func (c *Client) GetGroups(input *GetEntitiesInput) ([]models.Group, error) {
	path := fmt.Sprintf("search?type=group&term=%s", input.Term)
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

	groups := new([]models.Group)
	err = json.Unmarshal(body, groups)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *groups, nil
}

func (c *Client) GetGroupSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/group/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
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

	schedule := new([]models.Schedule)
	err = json.Unmarshal(body, schedule)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}

	return *schedule, nil
}

func (c *Client) GetTeacher(input *GetEntitiesInput) ([]models.Teacher, error) {
	path := fmt.Sprintf("search?type=person&term=%s", input.Term)
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

	teachers := new([]models.Teacher)
	err = json.Unmarshal(body, teachers)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *teachers, nil
}

func (c *Client) GetTeacherSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/person/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
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

	scheduleTeacher := new([]models.Schedule)
	err = json.Unmarshal(body, scheduleTeacher)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}
	return *scheduleTeacher, nil
}

func (c *Client) GetAuditorium(input *GetEntitiesInput) ([]models.Auditorium, error) {
	path := fmt.Sprintf("search?&term=%s&type=auditorium", input.Term)
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

	auditoriums := new([]models.Auditorium)
	err = json.Unmarshal(body, auditoriums)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *auditoriums, nil
}

func (c *Client) GetAuditoriumSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/auditorium/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	fmt.Println(res.StatusCode)
	defer res.Body.Close()

	scheduleAuditorium := new([]models.Schedule)
	err = json.Unmarshal(body, scheduleAuditorium)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}

	return *scheduleAuditorium, nil
}
