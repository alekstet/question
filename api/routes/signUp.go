package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
)

func (s *Store) signUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data models.SignUp
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	err = s.Querier.SignUp(data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	helpers.Render(w, r, http.StatusCreated, nil)
}
