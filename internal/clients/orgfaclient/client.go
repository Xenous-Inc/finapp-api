package orgfaclient

import (
	"encoding/json"
	"fmt"
	"strings"

	"io"
	"net/http"
	"net/url"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/models"
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

type LoginInput struct {
	Login    string
	Password string
}

func (c *Client) Login(input *LoginInput) (string, error) {
	path := "app/interaction/?login=yes"

	data := url.Values{}
	data.Set(models.AUTH_TYPE, models.Y)
	data.Set(models.TYPE, models.AUTH)
	data.Set(models.USER_LOGIN, input.Login)
	data.Set(models.USER_PASSWORD, input.Password)

	req := c.reqBuilder.SetMethod("POST").SetPath(path).SetBody([]byte(data.Encode())).SetContentURLEncoded().Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient login")
		return "", clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient login")
			return "", clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient login")
		return "", fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient login")
		return "", clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	if strings.Contains(string(rawBody), "errortext") {
		log.Warn("Unauthorized", "orgfaclient login")
		return "", clients.ErrUnauthorized
	}

	if strings.Contains(string(rawBody), "<title>Авторизация</title>") {
		log.Warn("Unauthorized", "orgfaclient login")
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
		log.Warn("Unauthorized, PHPSESSID is empty", "orgfaclient login")
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

func (c *Client) GetMyGroup(input *GetMyGroupInput) ([]models.Student, error) {
	path := "bitrix/vuz/api/interaction/myGroup"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient myGroup")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient myGroup")
			return nil, clients.ErrUnauthorized
		}

		log.Warn("BadRequest", "orgfaclient myGroup")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient myGroup")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	student := new(models.Data)
	err = json.Unmarshal(body, student)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient myGroup")
		return nil, clients.ErrInvalidEntity
	}

	if student.Error != 0 {
		log.Error(err, "InvalidResponse", "orgfaclient myGroup")
		return nil, fmt.Errorf("Unknown error got from ORG.FA.RU: ErrorCode: %d", student.Error)
	}

	return student.Student, nil
}

type GetRecordBookInput struct {
	*AuthSession
}

func (c *Client) GetRecordBook(input *GetRecordBookInput) ([]models.RecordBookItem, error) {
	path := "bitrix/vuz/api/marks2/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient recordBook")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient recordBook")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient recordBook")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient recordBook")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	recordBookList := make([]models.RecordBookItem, 0)
	err = json.Unmarshal(body, &recordBookList)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient recordBook")
		return nil, clients.ErrInvalidEntity
	}

	return recordBookList, nil
}

type GetMiniProfileInput struct {
	*AuthSession
}

func (c *Client) GetProfile(input *GetMiniProfileInput) (*models.MiniProfile, error) {
	path := "bitrix/vuz/api/profile/bootstrap"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient profile")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient profile")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient profile")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient profile")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	data := new(models.AllDataMiniProfile)
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient profile")
		return nil, clients.ErrInvalidEntity
	}

	return &data.MiniProfile, nil
}

type GetProfileInput struct {
	*AuthSession
}

func (c *Client) GetProfileDetails(input *GetProfileInput) (*models.ProfileDetails, error) {
	path := "bitrix/vuz/api/profile/current"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient profileDetails")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient profileDetails")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient profileDetails")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient profileDetails")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	profile := new(models.ProfileDetails)
	err = json.Unmarshal(body, &profile)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient profileDetails")
		return nil, clients.ErrInvalidEntity
	}

	return profile, nil
}

type GetOrderInput struct {
	*AuthSession
}

func (c *Client) GetOrder(input *GetOrderInput) ([]models.Order, error) {
	path := "bitrix/vuz/api/orders/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient order")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient order")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient order")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient order")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	order := make([]models.Order, 0)
	err = json.Unmarshal(body, &order)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient order")
		return nil, clients.ErrInvalidEntity
	}

	return order, nil
}

type GetStudentCardInput struct {
	*AuthSession
	ProfileId string
}

func (c *Client) GetStudentCard(input *GetStudentCardInput) (*models.StudentCard, error) {
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	path := fmt.Sprintf("bitrix/vuz/api/profiles/studentCard/%s", input.ProfileId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient studentCard")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient studentCard")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient studentCard")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient studentCard")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	studentCard := new(models.StudentCard)
	err = json.Unmarshal(body, &studentCard)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient studentCard")
		return nil, clients.ErrInvalidEntity
	}

	return studentCard, nil
}

type GetStudyPlanInput struct {
	*AuthSession
}

func (c *Client) GetStudyPlan(input *GetStudyPlanInput) ([]models.StudyPlan, error) {
	path := "bitrix/vuz/api/rups/"
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		log.Error(err, "BadRequest", "orgfaclient studyPlan")
		return nil, clients.ErrRequest
	}

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			log.Warn("Unauthorized", "orgfaclient studyPlan")
			return nil, clients.ErrUnauthorized
		}
		log.Warn("BadRequest", "orgfaclient studyPlan")
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err, "InvalidResponse", "orgfaclient studyPlan")
		return nil, clients.ErrInvalidResponse
	}

	defer res.Body.Close()

	studyPlan := make([]models.StudyPlan, 0)
	err = json.Unmarshal(body, &studyPlan)

	if err != nil {
		log.Error(err, "InvalidEntity", "orgfaclient studyPlan")
		return nil, clients.ErrInvalidEntity
	}

	return studyPlan, nil
}
