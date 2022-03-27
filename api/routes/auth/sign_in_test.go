package auth_test

import (
	"testing"

	"github.com/question/api/models"
	"github.com/question/testutils"

	"github.com/stretchr/testify/suite"
)

type SignInSuite struct {
	testutils.Store
}

func (s *SignInSuite) SetupSuiteSignIn() {
	s.Db = testutils.LoadDatabase()
}

func (s *SignInSuite) SignUpForSignIn() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "4567", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(201, w.Result().StatusCode)
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
	s.SignUpForSignIn()
	s.SetupSuiteSignIn()
	s.ValidSignIn()
	s.InvalidSignIn()
	s.ClearDatabaseSignIn()
}

func (s *SignInSuite) ClearDatabaseSignIn() {
	testutils.ClearDatabase(s.Db)
}

func Test_SignIn(t *testing.T) {
	suite.Run(t, new(SignInSuite))
}
