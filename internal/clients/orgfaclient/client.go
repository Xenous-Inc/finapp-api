package orgfaclient

import (
	"net/http"
	"net/url"

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

type SignInPayload struct {
	Login    string
	Password string
}

func (c *Client) SignIn(payload *SignInPayload) (*dto.SignInResponse, error) {
	formData := url.Values{
		"USER_LOGIN":    []string{payload.Login},
		"USER_PASSOWRD": []string{payload.Password},
		"AUTH":          []string{"Y"},
		"back_url":      []string{"/"},
	}
	encodedData := []byte(formData.Encode())

	req := c.reqBuilder.SetMethod("POST").SetContentURLEncoded().SetPath("/login=true").SetBody(encodedData).Build()
	resp, err := req.Execute(c.httpClient)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 302 {
		return nil, ErrUnauthorized
	}

	for _, e := range resp.Cookies() {
		if e.Name == "PHPSESSID" {
			if e.Value != "" {
				return &dto.SignInResponse{
					PhpSessionId: e.Value,
				}, nil
			}
		}
	}

	return nil, ErrUnauthorized
}

type GetProfilePayload struct {
}

func (c *Client) GetProfile(payload *GetProfilePayload) {

}
