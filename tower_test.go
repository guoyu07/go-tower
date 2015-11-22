package tower

import (
	"testing"
)

func TestPassword(t *testing.T) {

	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	t.Logf("%v", clt.AuthInfo.Teams[0].TeamGUID)
}

func TestTeam(t *testing.T) {
	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	if result, err := clt.TeamInfo(""); err != nil {
		t.Fatalf("%v", err.Error())
	} else {
		t.Logf("%v", result)
	}

}

func TestProjects(t *testing.T) {
	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	if result, err := clt.Projects(""); err != nil {
		t.Fatalf("%v", err.Error())
	} else {
		t.Logf("%v", len(result))
		t.Logf("%v", result)
	}
}

func TestProjectMembers(t *testing.T) {
	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	if result, err := clt.ProjectMembers(""); err != nil {
		t.Fatalf("%v", err.Error())
	} else {
		t.Logf("%v", len(result))
		t.Logf("%v", result)
	}
}

func TestTodoLists(t *testing.T) {
	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	if result, err := clt.TodoLists(""); err != nil {
		t.Fatalf("%v", err.Error())
	} else {
		t.Logf("%v", len(result))
		t.Logf("%v", result)
	}
}

func TestCreateTodolist(t *testing.T) {
	clientID, clientSecret := "", ""
	clt := NewClient(clientID, clientSecret, "", "")
	if err := clt.Auth(); err != nil {
		t.Fatalf("%v", err.Error())
	}
	if result, err := clt.CreateTodolist("", "test from tower"); err != nil {
		t.Fatalf("%v", err.Error())
	} else {
		t.Logf("%v", result)
	}
}
