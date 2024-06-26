package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-list/internal/models"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var person models.User

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id := h.services.Authorization.SignUp(person)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("person was created, id - %d", id)))
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var signInPerson models.SignInInput

	err := json.NewDecoder(r.Body).Decode(&signInPerson)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(signInPerson.Username, signInPerson.Password)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("token: %s", token)))
}
