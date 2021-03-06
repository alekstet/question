package process

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) updateAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.UserQuestion{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	err = s.Querier.UpdateAnswer(*data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	helpers.Render(w, r, http.StatusOK, nil)
}
