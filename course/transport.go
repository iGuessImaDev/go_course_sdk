package course

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/iGuessImaDev/gocourse_domain/domain"
	c "github.com/ncostamagna/go_http_client/client"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    string      `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	ClientHTTP struct {
		client c.Transport
	}
)

func NewHTTPClient(baseUrl, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &ClientHTTP{
		client: c.New(header, baseUrl, 5000*time.Millisecond, true),
	}
}

func (c *ClientHTTP) Get(id string) (*domain.Course, error) {

	dataResponse := DataResponse{Data: &domain.Course{}}

	u := url.URL{}
	u.Path += fmt.Sprintf("/courses/%s", id)
	resp := c.client.Get(u.String())
	if resp.Err != nil {
		return nil, resp.Err
	}

	if resp.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", resp)}
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("%s", resp)
	}

	if err := resp.FillUp(&dataResponse); err != nil {
		return nil, err
	}

	return dataResponse.Data.(*domain.Course), nil
}
