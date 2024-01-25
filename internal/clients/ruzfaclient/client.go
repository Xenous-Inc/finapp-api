package ruzfaclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient/models"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
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
		log.Error(err, "BadRequest", "ruzfaclient group")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient group")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	groups := new([]models.Group)
	err = json.Unmarshal(body, groups)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient group")
		return nil, clients.ErrInvalidEntity
	}

	return *groups, nil
}

func (c *Client) GetGroupSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/group/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "ruzfaclient group schedule")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient group schedule")
		return nil, clients.ErrInvalidResponse
	}
	defer res.Body.Close()

	schedule := new([]models.Schedule)
	err = json.Unmarshal(body, schedule)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient group schedule")
		return nil, clients.ErrInvalidEntity
	}

	return *schedule, nil
}

func (c *Client) GetTeacher(input *GetEntitiesInput) ([]models.Teacher, error) {
	path := fmt.Sprintf("search?type=person&term=%s", input.Term)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "ruzfaclient teacher")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient teacher")
		return nil, clients.ErrInvalidResponse
	}
	defer res.Body.Close()

	teachers := new([]models.Teacher)
	err = json.Unmarshal(body, teachers)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient teacher")
		return nil, clients.ErrInvalidEntity
	}

	return *teachers, nil
}

func (c *Client) GetTeacherSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/person/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "ruzfaclient teacher schedule")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient teacher schedule")
		return nil, clients.ErrInvalidResponse
	}
	defer res.Body.Close()

	scheduleTeacher := new([]models.Schedule)
	err = json.Unmarshal(body, scheduleTeacher)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient teacher schedule")
		return nil, clients.ErrInvalidEntity
	}
	return *scheduleTeacher, nil
}

func (c *Client) GetAuditorium(input *GetEntitiesInput) ([]models.Auditorium, error) {
	path := fmt.Sprintf("search?&term=%s&type=auditorium", input.Term)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "ruzfaclient classroom")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient classroom")
		return nil, clients.ErrInvalidResponse
	}
	defer res.Body.Close()

	auditoriums := new([]models.Auditorium)
	err = json.Unmarshal(body, auditoriums)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient classroom")
		return nil, clients.ErrInvalidEntity
	}

	return *auditoriums, nil
}

func (c *Client) GetAuditoriumSchedule(input *GetScheduleInput) ([]models.Schedule, error) {
	path := fmt.Sprintf("schedule/auditorium/%s?start=%s&finish=%s&lng=1", input.Id, input.StartDate, input.EndDate)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()
	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "ruzfaclient classroom schedule")
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "ruzfaclient classroom schedule")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	scheduleAuditorium := new([]models.Schedule)
	err = json.Unmarshal(body, scheduleAuditorium)
	if err != nil {
		log.Error(err, "InvalidEntity", "ruzfaclient classroom schedule")
		return nil, clients.ErrInvalidEntity
	}

	return *scheduleAuditorium, nil
}
