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

	data := url.Values{}
	data.Set("AUTH_FORM", "Y")
	data.Set("TYPE", "AUTH")
	data.Set("USER_LOGIN", input.Login)
	data.Set("USER_PASSWORD", input.Password)

	req := c.reqBuilder.SetMethod("POST").SetPath(path).SetBody([]byte(data.Encode())).SetContentURLEncoded().Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return "", clients.ErrRequest
	}

	rawBody, err := io.ReadAll(res.Body)
	fmt.Println(string(rawBody))
	defer res.Body.Close()

	if err != nil {
		return "", clients.ErrInvalidEntity
	}

	if strings.Contains(string(rawBody), "errortext") {
		fmt.Println(string(rawBody))
		return "", clients.ErrUnauthorized
	}

	if strings.Contains(string(rawBody), "<title>Авторизация</title>") {
		fmt.Println(string(rawBody))
		return "",clients.ErrUnauthorized
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()
	
	profileId := ""

	cookies := res.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "BX_ORG_FA_RU_PROFILE_ID" {
			profileId = cookie.Value
			break
		}
	}

	if profileId == "" {
		log.Println("Unable to find to profileId")
		return nil, clients.ErrUnauthorized
	}
	fmt.Println(profileId)

	student := &dto.Data{}
	err = json.Unmarshal(body, student)

	if err != nil {
		return nil, clients.ErrValidation
	}

	if student.Error != 0 {
		return nil, clients.ErrRequest
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	recordBookList := new([]dto.RecordBookItem)
	err = json.Unmarshal(body, &recordBookList)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
	}

	return *recordBookList, nil
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	miniProfile := new([]dto.MiniProfile)
	err = json.Unmarshal(body, &miniProfile)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
	}

	return *miniProfile, nil
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	profile := dto.AllDataProfile{}
	err = json.Unmarshal(body, &profile)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
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
		//fmt.Errorf(err.Error())
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		//fmt.Errorf(err.Error())
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	//order := new([]dto.Order)
	var order []dto.Order
	err = json.Unmarshal(body, &order)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
	}

	return order, nil
}

type GetStudentCardInput struct {
	*AuthSession
	ProfileId string
}

func (c *Client) GetStudentCard(input *GetStudentCardInput) (*dto.StudentCard, error) {
	phpSessionId := fmt.Sprintf("PHPSESSID=%s", input.SessionId)
	// path := fmt.Sprintf("bitrix/vuz/api/profiles/studentCard/%s", input.ProfileId) 
	path := "bitrix/vuz/api/profiles/studentCard/93491" 
	req := c.reqBuilder.SetMethod("GET").SetPath(path).AddHeader("Cookie", phpSessionId).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	//order := new([]dto.StudentCard)
	var order *dto.StudentCard
	err = json.Unmarshal(body, &order)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
	}

	return order, nil
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, clients.ErrRequest
	}
	fmt.Println(string(body))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			return nil, clients.ErrUnauthorized
		}

		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	studyPlan := new([]dto.StudyPlan)
	err = json.Unmarshal(body, &studyPlan)

	if err != nil {
		fmt.Println(err.Error())
		return nil, clients.ErrValidation
	}

	return *studyPlan, nil
}