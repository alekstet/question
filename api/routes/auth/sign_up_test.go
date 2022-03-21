package auth_test

import (
	"fmt"
	"question/api/models"
	"question/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignUpSuite struct {
	testutils.Store
}

func (s *SignUpSuite) SetupSuiteSignUp() {
	s.Db = testutils.LoadDatabase()
	fmt.Println("load ok")
}

func (s *SignUpSuite) ValidSignUp() {
	var r models.SignUp
	r.Init("aleks35", "aleks", "4567", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(201, w.Result().StatusCode)
}

func (s *SignUpSuite) ExistsUserSignUp() {
	var r models.SignUp
	r.Init("aleks35", "aleks", "4567", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(400, w.Result().StatusCode)
}

func (s *SignUpSuite) DiffPassSignUp() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "456", "4567")
	w := testutils.SendForm(s.T(), s.Db, "POST", "/signup", r)
	s.Assertions.Equal(400, w.Result().StatusCode)
}

func (s *SignUpSuite) Test() {
	s.ValidSignUp()
	s.ExistsUserSignUp()
	s.DiffPassSignUp()
}

func (s *SignUpSuite) ClearDatabaseSignUp() {
	testutils.ClearDatabase(s.Db)
}

func Test_SignUp(t *testing.T) {
	suite.Run(t, new(SignUpSuite))
}
