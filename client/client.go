package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/nsupc/eurogo/models"
)

type Client struct {
	username     string
	password     string
	token        string
	base_url     string
	last_refresh time.Time
	client       http.Client
}

// creates a new eurocore client.
func New(username string, password string, base_url string) *Client {
	base_url = strings.Trim(base_url, "/")

	client := Client{
		username: username,
		password: password,
		base_url: base_url,
		client:   http.Client{},
	}

	return &client
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

	url := fmt.Sprintf("%s/login", c.base_url)
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
	c.last_refresh = time.Now()

	return nil
}

func (c *Client) SendTelegram(t models.Telegram) error {
	telegrams := []models.Telegram{t}

	return c.SendTelegrams(telegrams)
}

func (c *Client) SendTelegrams(telegrams []models.Telegram) error {
	if time.Since(c.last_refresh) > time.Hour {
		err := c.refreshToken()
		if err != nil {
			return err
		}
	}

	url := fmt.Sprintf("%s/telegrams", c.base_url)

	data, err := json.Marshal(telegrams)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		errorText, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, errorText)
	}

	return nil
}
