package models

import (
	"testing"

	"github.com/gbbr/gopherblog/models/testdb"
)

func TestUserFetch(t *testing.T) {
	testdb.SetUp()

	testCases := []struct {
		Id              int
		xpName, xpEmail string
	}{
		{2, "Mathias", "mathias@company.it"},
		{1, "Jeremy", "jeremy@email.com"},
	}

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	for _, test := range testCases {
		user := &User{Id: test.Id}
		if err := user.Fetch(); err != nil {
			t.Log("An error occurred while fetching the user")
			t.Fatal()
		}

		if user.Name != test.xpName || user.Email != test.xpEmail {
			t.Log("Did not retrieve user correctly")
			t.Fail()
		}
	}
}

func TestUserLoginCorrect(t *testing.T) {
	testdb.SetUp()

	testCases := []struct {
		login, pass string
	}{
		{"jeremy@email.com", "5f4dcc3b5aa765d61d8327deb882cf99"},
		{"mathias@company.it", "7c6a180b36896a0a8c02787eeafb0e4c"},
	}

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	for _, test := range testCases {
		user := &User{
			Email:    test.login,
			Password: test.pass,
		}

		if !user.LoginCorrect() {
			t.Log("Failed to log in")
			t.Logf("%#v", user)
			t.Fail()
		}
	}
}
