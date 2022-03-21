package auth_test

import (
	"question/api/models"
	"question/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignInSuite struct {
	testutils.Store
}

func (s *SignInSuite) SetupSuiteSignIn() {
	s.Db = testutils.LoadDatabase()
}

func (s *SignInSuite) ValidSignIn() {
	var r models.SignIn
	r.Init("aleks", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signin", r)
	s.Assertions.Equal(200, w.Result().StatusCode)
}

func (s *SignInSuite) InvalidSignIn() {
	var r models.SignIn
	r.Init("aleks", "4566")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signin", r)
	s.Assertions.Equal(400, w.Result().StatusCode)
}

func (s *SignInSuite) TestSignIn() {
	s.ValidSignIn()
	s.InvalidSignIn()
}

func (s *SignInSuite) ClearDatabaseSignIn() {
	testutils.ClearDatabase(s.Db)
}

func Test_SignIn(t *testing.T) {
	suite.Run(t, new(SignInSuite))
}
