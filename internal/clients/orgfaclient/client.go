package orgfaclient

import (
	"encoding/json"
	"fmt"
	"strings"

	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
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

type LoginInput struct {
	Login    string
	Password string
}

func (c *Client) Login(input *LoginInput) (string, error) {
	path := "app/interaction/?login=yes"
	//package constants
	data := url.Values{}
	data.Set(dto.AUTH_FORM, dto.Y)
	data.Set(dto.TYPE, dto.AUTH)
	data.Set(dto.USER_LOGIN, input.Login)
	data.Set(dto.USER_PASSWORD, input.Password)

	req := c.reqBuilder.SetMethod("POST").SetPath(path).SetBody([]byte(data.Encode())).SetContentURLEncoded().Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return "", clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return "", clients.ErrUnauthorized
		}

		return "", fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	if strings.Contains(string(rawBody), "errortext") {
		return "", clients.ErrUnauthorized
	}

	if strings.Contains(string(rawBody), "<title>Авторизация</title>") {
		return "", clients.ErrUnauthorized
	}

	phpSessionId := ""

	cookies := res.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "PHPSESSID" {
			phpSessionId = cookie.Value
			break
		}
	}

	if phpSessionId == "" {
		log.Println("Unable to find to sessionId")
		return "", clients.ErrUnauthorized
	}

	return phpSessionId, nil
}

type AuthSession struct {
	SessionId string
}

type GetMyGroupInput struct {
	*AuthSession
}

func (c *Client) GetMyGroup(input *GetMyGroupInput) ([]dto.Student, error) {
	path := "bitrix/vuz/api/interaction/myGroup"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	student := &dto.Data{}
	err = json.Unmarshal(body, student)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	if student.Error != 0 {
		return nil, fmt.Errorf("Unknown error got from ORG.FA.RU: ErrorCode: %d", student.Error)
	}

	return student.Student, nil
}

type GetRecordBookInput struct {
	*AuthSession
}

func (c *Client) GetRecordBook(input *AuthSession) ([]dto.RecordBookItem, error) {
	path := "bitrix/vuz/api/marks2/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	recordBookList := make([]dto.RecordBookItem, 0)
	err = json.Unmarshal(body, &recordBookList)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return recordBookList, nil
}

type GetMiniProfileInput struct {
	*AuthSession
}

func (c *Client) GetMiniProfile(input *GetMiniProfileInput) ([]dto.MiniProfile, error) {
	path := "bitrix/vuz/api/profile/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	miniProfile := make([]dto.MiniProfile, 0)
	err = json.Unmarshal(body, &miniProfile)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return miniProfile, nil
}

type GetProfileInput struct {
	*AuthSession
}

func (c *Client) GetProfile(input *GetProfileInput) ([]dto.Profile, error) {
	path := "bitrix/vuz/api/profile/current"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	profile := dto.AllDataProfile{}
	err = json.Unmarshal(body, &profile)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return profile.Profile, nil
}

type GetOrderInput struct {
	*AuthSession
}

func (c *Client) GetOrder(input *GetOrderInput) ([]dto.Order, error) {
	path := "bitrix/vuz/api/orders/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	order := make([]dto.Order, 0)
	err = json.Unmarshal(body, &order)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return order, nil
}

type GetStudentCardInput struct {
	*AuthSession
	ProfileId string
}

func (c *Client) GetStudentCard(input *GetStudentCardInput) (*dto.StudentCard, error) {
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	path := fmt.Sprintf("bitrix/vuz/api/profiles/studentCard/%s", input.ProfileId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	var studentCard *dto.StudentCard
	err = json.Unmarshal(body, &studentCard)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return studentCard, nil
}

type GetStudyPlanInput struct {
	*AuthSession
}

func (c *Client) GetStudyPlan(input *GetStudyPlanInput) ([]dto.StudyPlan, error) {
	path := "bitrix/vuz/api/rups/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	studyPlan := make([]dto.StudyPlan, 0)
	err = json.Unmarshal(body, &studyPlan)

	if err != nil {
		return nil, clients.ErrInvalidEntity
	}

	return studyPlan, nil
}
