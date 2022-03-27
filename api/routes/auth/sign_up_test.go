package auth_test

import (
	"testing"

	"github.com/question/api/models"
	"github.com/question/testutils"

	"github.com/stretchr/testify/suite"
)

type SignUpSuite struct {
	testutils.Store
}

func (s *SignUpSuite) SetupSuiteSignUp() {
	s.Db = testutils.LoadDatabase()
}

func (s *SignUpSuite) ValidSignUp() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "4567", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(201, w.Result().StatusCode)
}

func (s *SignUpSuite) ExistsUserSignUp() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "4567", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(400, w.Result().StatusCode)
}

func (s *SignUpSuite) DiffPassSignUp() {
	var r models.SignUp
	r.Init("aleks35", "aleks", "456", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(400, w.Result().StatusCode)
}

func (s *SignUpSuite) Test() {
	s.SetupSuiteSignUp()
	s.ValidSignUp()
	s.ExistsUserSignUp()
	s.DiffPassSignUp()
	s.ClearDatabaseSignUp()
}

func (s *SignUpSuite) ClearDatabaseSignUp() {
	testutils.ClearDatabase(s.Db)
}

func Test_SignUp(t *testing.T) {
	suite.Run(t, new(SignUpSuite))
}
