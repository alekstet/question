package routes_test

import (
	"testing"
	"time"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/testutils"
	"github.com/stretchr/testify/suite"
)

type CreateAnswerSuite struct {
	testutils.Store
}

func (c *CreateAnswerSuite) SetupSuiteCreateAnswer() {
	c.Db = testutils.LoadDatabase()
}

func (c *CreateAnswerSuite) SignUp() {
	var r models.SignUp
	r.Init("aleks34", "aleks", "4567", "4567")
	w := testutils.SendForm(c.T(), c.Db, "POST", "/signup", r)
	c.Assertions.Equal(201, w.Result().StatusCode)
}

func (c *CreateAnswerSuite) AddQuestion() {
	var q models.Question
	date := time.Now().Format("02.01.2006")
	q.Init(date, "How")
	w := testutils.SendForm(c.T(), c.Db, "POST", "/questions", q)
	c.Assertions.Equal(201, w.Result().StatusCode)
}

func (c *CreateAnswerSuite) CreateAnswer() {
	var u models.UserQuestion
	date := time.Now().Format("02.01.2006")
	timeNow := time.Now().Format("02.01.2006 15:04:05")
	u.Init(date, "aleks34", "Good", timeNow, timeNow)
	w := testutils.SendForm(c.T(), c.Db, "POST", "/new", u)
	c.Assertions.Equal(201, w.Result().StatusCode)
}

func (c *CreateAnswerSuite) Test() {
	c.SetupSuiteCreateAnswer()
	c.SignUp()
	c.AddQuestion()
	c.CreateAnswer()
	c.ClearDatabaseCreateAnswer()
}

func (c *CreateAnswerSuite) ClearDatabaseCreateAnswer() {
	testutils.ClearDatabase(c.Db)
}

func Test_Create_Answer(t *testing.T) {
	suite.Run(t, new(CreateAnswerSuite))
}
