package middesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	APIKey string
}

type Client struct {
	config *Config
	client *http.Client
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("https://api.middesk.com/v1/%s", path)
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: http.DefaultClient,
	}
}

type WebhookRequest struct {
	URL    string  `json:"url"`
	Secret *string `json:"secret,omitempty"`
}

type WebhookResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) CreateWebhook(request *WebhookRequest) (*WebhookResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url("webhooks"), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.APIKey, "")
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error making request: %s", body)
	}

	webhook := &WebhookResponse{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}

func (c *Client) GetWebhook(id string) (*WebhookResponse, error) {
	req, err := http.NewRequest("GET", c.url(fmt.Sprintf("webhooks/%s", id)), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.APIKey, "")
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error making request: %s", body)
	}

	webhook := &WebhookResponse{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}

func (c *Client) DeleteWebhook(id string) error {
	req, err := http.NewRequest("DELETE", c.url(fmt.Sprintf("webhooks/%s", id)), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.config.APIKey, "")
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Delete unsuccessful, expected 200 response, got %d", res.StatusCode)
	}

	return nil
}

func (c *Client) UpdateWebhook(id string, request *WebhookRequest) (*WebhookResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", c.url(fmt.Sprintf("webhooks/%s", id)), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.config.APIKey, "")
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error making request: %s", body)
	}

	webhook := &WebhookResponse{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}

	return webhook, nil
}
