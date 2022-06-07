package routes_test

import (
	"testing"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/testutils"
	"github.com/stretchr/testify/suite"
)

type GetAnswersSuite struct {
	testutils.Store
}

func (g *GetAnswersSuite) SetupSuiteGetAnswers() {
	g.Db = testutils.LoadDatabase()
}

func (g *GetAnswersSuite) AddQuestion() {
	var r models.Question
	r.Init("24.03.2022", "Lol1")
	w := testutils.SendForm(g.T(), g.Db, "POST", "/questions", nil)
	g.Assertions.Equal(201, w.Result().StatusCode)
}

func (g *GetAnswersSuite) GetAnswers() {
	w := testutils.SendForm(g.T(), g.Db, "GET", "/new", nil)
	g.Assertions.Equal(200, w.Result().StatusCode)
}

func (g *GetAnswersSuite) Test() {
	g.SetupSuiteGetAnswers()
	g.AddQuestion()
	g.GetAnswers()
	g.ClearDatabaseGetAnswers()
}

func (g *GetAnswersSuite) ClearDatabaseGetAnswers() {
	testutils.ClearDatabase(g.Db)
}

func Test_Get_Answers(t *testing.T) {
	suite.Run(t, new(GetAnswersSuite))
}
