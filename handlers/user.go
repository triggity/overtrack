package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"

	"github.com/gorilla/mux"
	"github.com/triggity/overtrack/models"
)

type UserHandler struct {
	dao *models.UserDao
}

func NewUserHandler(client *elastic.Client) *UserHandler {
	return &UserHandler{
		models.NewUserDao(client),
	}
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.dao.List(r.Context())
	if err != nil {
		log.Info("Ohhh no!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(struct { Users []models.User `json:"users"`}{users})
	w.Write(s)
}

func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	users, err := u.dao.GetByID(r.Context(), id)
	if err != nil {
		log.Info("Ohhh no!", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, err := json.Marshal(struct{ User models.User `json:"user"` }{users})
	if err != nil {
		log.Info("EEEEEEEE", err)
	}
	w.Write(s)
}
