package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/nsupc/eurogo/dispatches"
	"github.com/nsupc/eurogo/telegrams"
	"github.com/nsupc/eurogo/templates"
)

type Client struct {
	username    string
	password    string
	token       string
	baseUrl     string
	lastRefresh time.Time
	client      http.Client
}

// creates a new eurocore client.
func New(username string, password string, base_url string) *Client {
	base_url = strings.Trim(base_url, "/")

	client := Client{
		username: username,
		password: password,
		baseUrl:  base_url,
		client: http.Client{
			Timeout: 5 * time.Second,
		},
	}

	return &client
}

func (c *Client) validateToken() error {
	if time.Since(c.lastRefresh) > time.Hour {
		err := c.refreshToken()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) refreshToken() error {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type ResponseData struct {
		Token string `json:"token"`
	}

	loginData := LoginData{
		Username: c.username,
		Password: c.password,
	}

	url := fmt.Sprintf("%s/login", c.baseUrl)
	data, err := json.Marshal(loginData)
	if err != nil {
		return err
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	responseData := ResponseData{}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return err
	}

	c.token = responseData.Token
	c.lastRefresh = time.Now()

	return nil
}

func (c *Client) makeRequest(method string, endpoint string, body io.Reader) (*http.Response, error) {
	err := c.validateToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", c.baseUrl, endpoint)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetTelegrams() (telegrams.List, error) {
	t := telegrams.List{}

	err := c.validateToken()
	if err != nil {
		return t, err
	}

	resp, err := c.makeRequest("GET", "/telegrams", http.NoBody)
	if err != nil {
		return t, err
	}

	if resp.StatusCode != 200 {
		errorText, _ := io.ReadAll(resp.Body)
		return t, fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// convenience wrapper around client.SendTelegrams() with a single telegram
func (c *Client) SendTelegram(t telegrams.NewTelegram) error {
	telegrams := []telegrams.NewTelegram{t}

	return c.SendTelegrams(telegrams)
}

func (c *Client) SendTelegrams(t []telegrams.NewTelegram) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	for _, tg := range t {
		if tg.Type == telegrams.Undefined {
			return errors.New("telegram type is undefined")
		}
	}

	resp, err := c.makeRequest("POST", "/telegrams", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		errorText, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	return nil
}

func (c *Client) DeleteTelegram(t telegrams.DeleteTelegram) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	resp, err := c.makeRequest("DELETE", "/telegrams", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		errorText, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	return nil
}

func (c *Client) dispatch(method string, endpoint string, data []byte) (dispatches.Status, error) {
	s := dispatches.Status{}

	resp, err := c.makeRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return s, err
	}

	if resp.StatusCode != 201 {
		errorText, _ := io.ReadAll(resp.Body)
		return s, fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (c *Client) CreateDispatch(d dispatches.NewDispatch) (dispatches.Status, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return dispatches.Status{}, err
	}

	return c.dispatch("POST", "/dispatches", data)
}

func (c *Client) EditDispatch(d dispatches.EditDispatch) (dispatches.Status, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return dispatches.Status{}, err
	}

	endpoint := fmt.Sprintf("/dispatches/%d", d.Id)

	return c.dispatch("PUT", endpoint, data)
}

func (c *Client) DeleteDispatch(id int) (dispatches.Status, error) {
	endpoint := fmt.Sprintf("/dispatches/%d", id)

	return c.dispatch("PUT", endpoint, []byte{})
}

func (c *Client) template(method string, endpoint string, statusCode int, reqBody io.Reader) (templates.Template, error) {
	t := templates.Template{}

	resp, err := c.makeRequest(method, endpoint, reqBody)
	if err != nil {
		return t, err
	}

	if resp.StatusCode != statusCode {
		errorText, _ := io.ReadAll(resp.Body)
		return t, fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (c *Client) GetTemplate(id string) (templates.Template, error) {
	endpoint := fmt.Sprintf("/templates/%s", id)

	return c.template("GET", endpoint, 200, http.NoBody)
}

func (c *Client) CreateTemplate(t templates.NewTemplate) (templates.Template, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return templates.Template{}, err
	}

	return c.template("POST", "/templates", 201, bytes.NewBuffer(data))
}

func (c *Client) EditTemplate(t templates.EditTemplate) (templates.Template, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return templates.Template{}, err
	}

	endpoint := fmt.Sprintf("/templates/%s", t.Id)

	return c.template("PATCH", endpoint, 200, bytes.NewBuffer(data))
}
