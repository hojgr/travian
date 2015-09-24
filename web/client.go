package web

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

// Client is HTTP interface that is customized for Travian
type Client struct {
	GoClient *http.Client
	BaseURL  string
}

// NewClient creates new instance of Client struct
func NewClient(baseURL string) *Client {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	return &Client{client, baseURL}
}

// POST Creates POST request
func (c *Client) POST(url string, vals url.Values) (resp *http.Response, err error) {
	return c.GoClient.PostForm(url, vals)
}

// GET creates GET request
func (c *Client) GET(url string) (resp *http.Response, err error) {
	return c.GoClient.Get(url)
}

// Login logs client in
func (c *Client) Login(username, password string) {
	c.POST(c.BaseURL, url.Values{
		"name": {username},
		"pass": {password},
	})
}

// GetStatisticsHTML Returns response from Statistics.php
func (c *Client) GetStatisticsHTML() (*http.Response, error) {
	resp, err := c.GoClient.Get(c.BaseURL + "/statistics.php")
	return resp, err
}

// GetVillage1HTML Returns response from Village1.php
func (c *Client) GetVillage1HTML() (*http.Response, error) {
	resp, err := c.GoClient.Get(c.BaseURL + "/village1.php")
	return resp, err
}

// GetCookie returns PHPSESSID
func (c *Client) GetCookie() string {
	url, _ := url.Parse("http://s5.zravian.com")
	cookies := c.GoClient.Jar.Cookies(url)

	for _, cookie := range cookies {
		if cookie.Name == "PHPSESSID" {
			return cookie.Value
		}
	}

	return ""
}

// UpgradeField upgrades field to a higher level
func (c *Client) UpgradeField(id int, key string) {
	c.GoClient.Get(c.BaseURL + "/village1.php?id=" + strconv.Itoa(id) + "&k=" + key)
}
