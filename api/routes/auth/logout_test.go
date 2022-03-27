package auth_test

import (
	"fmt"
	"testing"

	"github.com/question/api/models"
	"github.com/question/testutils"

	"github.com/stretchr/testify/suite"
)

type LogoutSuite struct {
	testutils.Store
}

func (l *LogoutSuite) SetupSuiteLogout() {
	l.Db = testutils.LoadDatabase()
}

func (l *LogoutSuite) SignIn() {
	var r models.SignIn
	r.Init("sasha2010", "21narufu")
	fmt.Println(r)
	w := testutils.SendForm(l.T(), l.Db, "POST", "/signin", r)
	l.Assertions.Equal(200, w.Result().StatusCode)
}

func (l *LogoutSuite) InvalidSignIn() {
	var r models.SignIn
	r.Init("aleks", "4566")
	w := testutils.SendForm(l.T(), l.Db, "POST", "/logout", r)
	l.Assertions.Equal(400, w.Result().StatusCode)
}

func (l *LogoutSuite) TestLogout() {
	l.SetupSuiteLogout()
	l.SignIn()
	l.InvalidSignIn()
	l.ClearDatabaseLogout()
}

func (l *LogoutSuite) ClearDatabaseLogout() {
	testutils.ClearDatabase(l.Db)
}

func Test_Logout(t *testing.T) {
	suite.Run(t, new(LogoutSuite))
}
