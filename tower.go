package tower

import (
	"fmt"
	"github.com/CuriosityChina/requests"
	"net/url"
	"time"
)

type Tower struct {
	clientID     string
	clientSecret string
	username     string
	password     string

	AuthInfo struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`

		GUID     string `json:"guid"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		Teams    []Team `json:"teams"`
	}
}

func NewClient(clientID, clientSecret, username, password string) *Tower {
	requests.SetDebug(true)
	return &Tower{
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
	}
}

func (c *Tower) Auth() error {
	var (
		url     = "https://tower.im/api/v2/oauth/token"
		request = struct {
			ClientID     string `json:"client_id"    bson:"client_id"`
			ClientSecret string `json:"client_secret"    bson:"client_secret"`
			Username     string `json:"username"`
			Password     string `json:"password"`
			GrantType    string `json:"grant_type"    bson:"grant_type"`
		}{
			ClientID:     c.clientID,
			ClientSecret: c.clientSecret,
			Username:     c.username,
			Password:     c.password,
			GrantType:    "password",
		}
	)
	if _, err := requests.Post(url, request, nil, &c.AuthInfo); err != nil {
		return err
	}
	fmt.Println(">>>", c.AuthInfo.TokenType, c.AuthInfo.AccessToken)
	return nil

}
func (c *Tower) authHeader() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("%s %s", c.AuthInfo.TokenType, c.AuthInfo.AccessToken)}
}

// 获取团队信息
func (c *Tower) TeamInfo(guid string) (TeamInfo, error) {
	var (
		url      = "https://tower.im/api/v2/teams/" + guid
		response TeamInfo
	)
	if _, err := requests.Get(url, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

// 获取团队项目列表
func (c *Tower) Projects(guid string) ([]Project, error) {
	var (
		url      = "https://tower.im/api/v2/teams/" + guid + "/projects"
		response []Project
	)
	if _, err := requests.Get(url, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

// 获取项目成员列表
func (c *Tower) ProjectMembers(id string) ([]Member, error) {
	var (
		url      = "https://tower.im/api/v2/projects/" + id + "/members"
		response []Member
	)
	if _, err := requests.Get(url, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

// 获取项目任务清单列表
func (c *Tower) TodoLists(projectID string) ([]TodoList, error) {
	var (
		url      = "https://tower.im/api/v2/projects/" + projectID + "/todolists"
		response []TodoList
	)
	if _, err := requests.Get(url, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Tower) CreateTodolist(projectID, title string) (TodoList, error) {
	var (
		url      = "https://tower.im/api/v2/projects/" + projectID + "/todolists"
		response TodoList
		request  = struct {
			Title string `json:"title"`
		}{
			Title: title,
		}
	)
	if _, err := requests.Post(url, request, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Tower) Todos(id string, since, till time.Time, count, completed int) ([]Todo, error) {
	var (
		url      = "https://tower.im/api/v2/todolists/" + id + "/todos"
		response TodoList
		request  = struct {
			Since     time.Time `json:"since,omitempty"`
			Till      time.Time `json:"till,omitempty"`
			Count     int       `json:"count,,omitempty"`
			Completed int       `json:"completed"`
		}{
			Since:     since,
			Till:      till,
			Count:     count,
			Completed: completed,
		}
	)

	if _, err := requests.Get(url, c.authHeader(), &response); err != nil {
		return response, err
	}
	return response, nil
}

// func (c *Tower) CreateTodo(listGUID string) (Todo, error) {
// }

// func (c *Tower) CommentOnTodo(id, content string) error {

// }
// func (c *Tower) CommentOnTodolist(id, content string) error {

// }
