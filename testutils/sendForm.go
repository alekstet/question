package testutils

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"unsafe"

	"github.com/alekstet/question/api/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Store struct {
	suite.Suite
	Db *sql.DB
}

func NewStore() *routes.Store {
	return &routes.Store{
		Log:    logrus.New(),
		Routes: httprouter.New(),
	}
}

func SendForm(t *testing.T, db *sql.DB, method string, target string, body interface{}) *httptest.ResponseRecorder {
	var Body io.Reader
	store := NewStore()
	store.Db = db
	mux := routes.Routes(*store)

	switch v := body.(type) {
	case url.Values:
		body = strings.NewReader(v.Encode())
	case string:
		body = strings.NewReader(v)
	case nil:
	default:
		res, err := json.Marshal(v)
		assert.Nil(t, err)
		body = strings.NewReader(*(*string)(unsafe.Pointer(&res)))
	}

	r := httptest.NewRequest(method, target, Body)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	return w
}
