package clients

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type ClientInterface interface {
	DoReq(target string) ([]byte, error)
}

type Client struct {
	client *http.Client
}

func NewClient(cli *http.Client) ClientInterface {
	return &Client{client: cli}
}

func (c *Client) DoReq(target string) ([]byte, error) {
	// check if url has scheme
	if !strings.HasPrefix("http://", target) {
		target = "http://" + target
	}
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
