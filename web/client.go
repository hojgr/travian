package web

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

// GetVillage1HTML Returns response from Village1.php
func (c *Client) GetVillage2HTML() (*http.Response, error) {
	resp, err := c.GoClient.Get(c.BaseURL + "/village2.php")
	return resp, err
}

// GetComposeMessageHTML Returns response from Village1.php
func (c *Client) GetComposeMessageHTML() (*http.Response, error) {
	resp, err := c.GoClient.Get(c.BaseURL + "/msg.php?t=1")
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

// UpgradeField upgrades field to a higher level
func (c *Client) BuildBuilding(parcelId int, buildingId int, key string) {
	parcelStr := strconv.Itoa(parcelId)
	buildingStr := strconv.Itoa(buildingId)
	c.GoClient.Get(c.BaseURL + "/village2.php?id=" + parcelStr + "&b=" + buildingStr + "&k=" + key)
}

// GetActionKey returns a key for actions (building, upgrading, ...)
func (c *Client) GetActionKey(resp *http.Response) (key string, found bool) {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	keyInput := doc.Find("input[type=hidden][name=k]").First()

	return keyInput.Attr("value")
}

type BuildingQueue struct {
	Tasks []Task
}

type Task struct {
	Name     string
	OldLevel int
	NewLevel int
	TimeLeft int
}

func (c *Client) GetBuildingQueue(resp *http.Response) BuildingQueue {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	contract := doc.Find("table#building_contract")

	queue := BuildingQueue{}

	contract.Find("tbody tr").Each(func(index int, s *goquery.Selection) {
		name, newLevel, oldLevel := parseNameAndLevel(s.Find("td:nth-child(2)").Text())
		time := parseStringTimeToSeconds(s.Find("td:nth-child(3) span").Text())

		queue.Tasks = append(queue.Tasks, Task{
			Name:     name,
			OldLevel: oldLevel,
			NewLevel: newLevel,
			TimeLeft: time,
		})
	})

	return queue
}

func parseNameAndLevel(nameAndLevel string) (name string, newLevel, oldLevel int) {
	nameAndLevelRegexp := regexp.MustCompile("(.*) \\(level ([0-9]+)")

	matches := nameAndLevelRegexp.FindAllStringSubmatch(nameAndLevel, 1)
	name = matches[0][1]
	newLevelString := matches[0][2]

	newLevel, _ = strconv.Atoi(newLevelString)
	oldLevel = newLevel - 1

	return name, newLevel, oldLevel
}

func parseStringTimeToSeconds(time string) int {
	timeArr := strings.Split(time, ":")

	hour, _ := strconv.Atoi(timeArr[0])
	minute, _ := strconv.Atoi(timeArr[1])
	second, _ := strconv.Atoi(timeArr[2])

	return hour*60*60 + minute*60 + second
}

func (c *Client) Raid(villageId string, key string) {
	village := villageId

	troopId := "11"
	troopCount := "5"

	v := url.Values{}
	v.Set("id", village)
	v.Set("c", "4")
	v.Set("t["+troopId+"]", troopCount)
	v.Set("k", key)

	c.POST(c.BaseURL+"/v2v.php", v)

}

func (c *Client) CanAttack(x, y string, key string) bool {
	troopId := "11"
	troopCount := "1"

	v := url.Values{}
	v.Set("x", x)
	v.Set("y", y)

	v.Set("c", "4")

	v.Set("t["+troopId+"]", troopCount)
	v.Set("k", key)

	res, _ := c.POST(c.BaseURL+"/v2v.php", v)

	doc, _ := goquery.NewDocumentFromResponse(res)

	f := doc.Find("div#content p.error")
	contents, _ := f.Html()

	return contents == "" // if empty, no error occurred, it can go forward
}
