package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	host   string
	port   string
	client *http.Client
}

func New(host, port string) (*Client, error) {
	if host == "" || port == "" {
		return nil, errors.New("host and port can not be blank")
	}
	return &Client{host: host, port: port, client: &http.Client{Timeout: time.Second * 5}}, nil
}

func (c *Client) HandleInsert(collection string, key string, data map[string]any) (map[string]string, error) {
	url := fmt.Sprintf("http://%v:%v/api/v1/insert/%v/%v", c.host, c.port, collection, key)
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var response map[string]string
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (c *Client) HandleSelect(collection string, key string) (map[string]any, error) {
	url := fmt.Sprintf("http://%v:%v/api/v1/select/%v/%v", c.host, c.port, collection, key)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var response map[string]any
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
