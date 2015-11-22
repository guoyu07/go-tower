package tower

import (
	"time"
)

type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"    bson:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	GUID         string `json:"guid"`
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	Teams        []Team `json:"teams"`
}

type Team struct {
	TeamName    string    `json:"team_name"`
	TeamGUID    string    `json:"team_guid"`
	TeamCreated time.Time `json:"team_created"`
	MemberGUID  string    `json:"member_guid"`
	Avatar      string    `json:"avatar"`
}

type TeamInfo struct {
	GUID     string   `json:"guid"`
	Name     string   `json:"name"`
	IsLocked bool     `json:"is_locked"`
	Members  []Member `json:"members"`
}

type Member struct {
	GUID     string `json:"guid"`
	Nickname string `json:"nickname"`
	IsActive bool   `json:"is_active,string"`
	Avatar   string `json:"avatar"`
}

type Project struct {
	GUID           string `json:"guid"`
	Name           string `json:"name"`
	IsArchived     bool   `json:"is_archived"`
	IconID         string `json:"icon_id"    bson:"icon_id"`
	ColorID        string `json:"color_id"    bson:"color_id"`
	Desc           string `json:"desc"    bson:"desc"`
	TodolistsCount struct {
		Completed   int `json:"completed"`
		Uncompleted int `json:"uncompleted"`
	} `json:"todolists_count`
}

type TodoList struct {
	GUID      string `json:"guid"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active,string"`
	Stick     bool   `json:"stick"`
	Completed bool   `json:"completed"`
}

type Todo struct {
	GUID        string      `json:"guid"`
	Content     string      `json:"content"`
	IsCompleted bool        `json:"is_completed"`
	Status      int         `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	DueAt       interface{} `json:"due_at"`
	IsActive    bool        `json:"is_active"`
	ClosedAt    time.Time   `json:"closed_at"`
	Creator     struct {
		GUID     string `json:"guid"`
		Nickname string `json:"nickname"`
		IsActive bool   `json:"is_active"`
		Avatar   string `json:"avatar"`
	} `json:"creator"`
	Assignee struct {
		GUID     string `json:"guid"`
		Nickname string `json:"nickname"`
		IsActive bool   `json:"is_active"`
		Avatar   string `json:"avatar"`
	} `json:"assignee"`
	Project struct {
		GUID       string `json:"guid"`
		Name       string `json:"name"`
		IsArchived bool   `json:"is_archived"`
	} `json:"project"`
	Todolist struct {
		GUID      string `json:"guid"`
		Name      string `json:"name"`
		IsActive  bool   `json:"is_active"`
		Completed bool   `json:"completed"`
	} `json:"todolist"`
	Closer struct {
		GUID     string `json:"guid"`
		Nickname string `json:"nickname"`
		IsActive bool   `json:"is_active"`
		Avatar   string `json:"avatar"`
	} `json:"closer"`
}
