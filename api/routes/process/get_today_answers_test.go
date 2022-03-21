package process_test

import (
	"fmt"
	"question/testutils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetAnswersSuite struct {
	testutils.Store
}

func (g *GetAnswersSuite) SetupSuiteGetAnswers() {
	g.Db = testutils.LoadDatabase()
	fmt.Println("load ok")
}

func (g *GetAnswersSuite) GetAnswers() {
	w := testutils.SendForm(g.T(), g.Db, "GET", "/new", nil)
	g.Assertions.Equal(201, w.Result().StatusCode)
}

func (g *GetAnswersSuite) Test() {
	g.GetAnswers()
}

func (g *GetAnswersSuite) ClearDatabaseGetAnswers() {
	testutils.ClearDatabase(g.Db)
}

func Test_Get_Answers(t *testing.T) {
	suite.Run(t, new(GetAnswersSuite))
}
