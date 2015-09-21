package web

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	GoClient *http.Client
	BaseURL  string
}

func NewClient(baseURL string) *Client {
	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	return &Client{client, baseURL}
}

func (c *Client) POST(url string, vals url.Values) (resp *http.Response, err error) {
	return c.GoClient.PostForm(url, vals)
}

func (c *Client) GET(url string) (resp *http.Response, err error) {
	return c.GoClient.Get(url)
}

func (c *Client) GetStatisticsHTML() (map[string]string, error) {
	resp, err := c.GoClient.Get(c.BaseURL + "/statistics.php")

	if resp.StatusCode != 200 {
		return nil, errors.New("Status code is not 200, it is " + string(resp.StatusCode))
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)

	s := doc.Find("tr.hl").First()

	position := strings.TrimRight(s.Find(".ra").Text(), ".")
	username := strings.Trim(s.Find(".pla").Find("a").Text(), " ")
	alliance := s.Find("al").Text()
	pop := s.Find(".pop").Text()
	village_count := s.Find(".vil").Text()

	return map[string]string{
		"position":      position,
		"username":      username,
		"alliance":      alliance,
		"pop":           pop,
		"village_count": village_count,
	}, nil
}
