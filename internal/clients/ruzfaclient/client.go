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
	GroupTerm string `json:"group"`
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

type GetGroupScheduleInput struct {
	GroupId   string `json:"groupId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (c *Client) GetGroupSchedule(input *GetGroupScheduleInput) ([]dto.Schedule, error) {
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

type GetTeacherInput struct {
	TeacherTerm string `json:"teacher"`
}

func (c *Client) GetTeacher(input *GetTeacherInput) ([]dto.Teacher, error) {
	path := fmt.Sprintf("search?type=person&term=%s", input.TeacherTerm)
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

	teachers := new([]dto.Teacher)
	err = json.Unmarshal(body, teachers)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *teachers, nil
}

type GetTeacherScheduleInput struct {
	Id        string `json:"teacherId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (c *Client) GetTeacherSchedule(input *GetTeacherScheduleInput) ([]dto.Schedule, error) {
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

	fmt.Println(res.StatusCode)
	defer res.Body.Close()

	scheduleTeacher := new([]dto.Schedule)
	err = json.Unmarshal(body, scheduleTeacher)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}
	return *scheduleTeacher, nil
}

type GetAuditoriumInput struct {
	AuditoriumTerm string `json:"auditorium"`
}

func (c *Client) GetAuditorium(input *GetAuditoriumInput) ([]dto.Auditorium, error) {
	path := fmt.Sprintf("search?&term=%s&type=auditorium", input.AuditoriumTerm)
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

	auditoriums := new([]dto.Auditorium)
	err = json.Unmarshal(body, auditoriums)
	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return *auditoriums, nil
}

type GetAuditoriumScheduleInput struct {
	Id        string `json:"auditoriumId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (c *Client) GetAuditoriumSchedule(input *GetAuditoriumScheduleInput) ([]dto.Schedule, error) {
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

	scheduleAuditorium := new([]dto.Schedule)
	err = json.Unmarshal(body, scheduleAuditorium)
	if err != nil {
		fmt.Print(err)
		return nil, clients.ErrValidation
	}

	return *scheduleAuditorium, nil
}
