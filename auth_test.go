package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gbbr/gopherblog/models"
	"github.com/gbbr/gopherblog/models/testdb"
)

const (
	REMOTE_ADDR_TEST  = "127.0.0.1:123"
	REMOTE_USER_AGENT = "User-Agent-123"
	LOGIN_EMAIL       = "jeremy@email.com"
	LOGIN_ID          = 1
)

// Returns a test request. If valid is true, the request
// will be a valid authentication.
func getTestRequest(valid bool) (*http.Request, error) {
	req, err := http.NewRequest("GET", "/manage", nil)
	if err != nil {
		return nil, err
	}

	hash := "1:GOOD_ID_BUT_BAD_HASH"
	if valid {
		ip := strings.Split(REMOTE_ADDR_TEST, ":")[0]
		origin := []byte(LOGIN_EMAIL + ip + REMOTE_USER_AGENT)
		hash = fmt.Sprintf("%d:%x", LOGIN_ID, sha256.Sum256(origin))
	}

	req.Header.Add("User-Agent", REMOTE_USER_AGENT)
	req.RemoteAddr = REMOTE_ADDR_TEST
	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: hash,
	})

	return req, nil
}

func TestIsValidRequest(t *testing.T) {
	testdb.Config.SchemaFile = "models/testdb/schema.sql"
	testdb.SetUp()

	models.ConnectDb(testdb.Config.DbString)
	defer models.CloseDb()

	want := models.User{
		Id:       1,
		Name:     "Jeremy",
		Email:    "jeremy@email.com",
		Password: "",
	}

	// Valid request
	req, err := getTestRequest(true)
	if err != nil {
		t.Log("Error constructing request")
		t.Fail()
	}

	user, valid := isValidRequest(req)
	if !valid || !reflect.DeepEqual(want, *user) {
		t.Logf("Failed to validate good request")
		t.Fail()
	}

	// Invalid request
	req, err = getTestRequest(false)
	if err != nil {
		t.Log("Error constructing request")
		t.Fail()
	}

	user, valid = isValidRequest(req)
	if valid || user != nil {
		t.Logf("Validated bad request")
		t.Fail()
	}
}

func TestAuthenticateRoute(t *testing.T) {
	testdb.Config.SchemaFile = "models/testdb/schema.sql"
	testdb.SetUp()

	models.ConnectDb(testdb.Config.DbString)
	defer models.CloseDb()

	handlerCalled := false
	returnedHandler := authenticate(func(w http.ResponseWriter, r *http.Request, u *models.User) {
		handlerCalled = true
	})

	// Valid request
	req, err := getTestRequest(true)
	if err != nil {
		t.Log("Error constructing request")
		t.Fail()
	}

	returnedHandler(httptest.NewRecorder(), req)
	if !handlerCalled {
		t.Log("Handler was not called on valid request")
	}

	// Invalid request
	req, err = getTestRequest(false)
	if err != nil {
		t.Log("Error constructing request")
		t.Fail()
	}

	handlerCalled = false
	returnedHandler(httptest.NewRecorder(), req)
	if handlerCalled {
		t.Log("Handler was called on invalid request")
	}
}
