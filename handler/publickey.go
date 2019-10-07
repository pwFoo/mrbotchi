package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/mrbotchi-team/mrbotchi/securityvocabulary"
	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	PublickeyHandler struct {
		HTTPHandler
	}
)

func (h PublickeyHandler) Get(w http.ResponseWriter, r *http.Request) error {
	name := chi.URLParam(r, "name")
	if h.App.Config.User.Name != name {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	id := fmt.Sprintf("https://%s/%s", h.App.Config.Host, name)
	endpoint := strings.Join([]string{id, "/publickey"}, "")
	publickey := securityvocabulary.NewKey(endpoint, id, h.App.Config.User.PublicKey)

	body, err := json.Marshal(publickey)
	if nil != err {
		return err
	}

	return utils.WriteBody(w, body, http.StatusOK, "application/activity+json; charset=utf8")
}