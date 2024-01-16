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
	Login string
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

func (c *Client) GetMyGroup(input *AuthSession) ([]dto.Student, error) {
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
