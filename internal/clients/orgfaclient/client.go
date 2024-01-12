package orgfaclient

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
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
	path := "auth?login=yes"
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("AUTH_FORM", "Y")
	writer.WriteField("TYPE", "AUTH")
	writer.WriteField("USER_LOGIN", input.Login)
	writer.WriteField("USER_PASSWORD", input.Password)

	req := c.reqBuilder.SetMethod("POST").SetPath(path).SetBody(body.Bytes()).AddHeader("Content-Type", writer.FormDataContentType()).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return "", clients.ErrRequest
	}

	rawBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return "", clients.ErrInvalidEntity
	}

	if strings.Contains(string(rawBody), "errortext") {
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

func (c *Client) GetMyGroup(input *LoginInput) (string, error) {
	path := "app/interaction/myGroup"
	req := c.reqBuilder.SetMethod("GET").SetPath(path).Build()

	res, err := req.Execute(c.httpClient)
	if err != nil {
		return "", clients.ErrRequest
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", clients.ErrInvalidEntity
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