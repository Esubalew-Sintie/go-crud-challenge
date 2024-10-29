// internal/adapters/http/handler.go
package http

import (
	"encoding/json"
	"go-crud-challenge/internal/domain"
	"go-crud-challenge/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(service *service.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) GetPersons(w http.ResponseWriter, r *http.Request) {
	
	persons, err := h.service.GetAllPersons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(persons)
}

func (h *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id != "" {
		person, err := h.service.GetPersonByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(person)
		return
	}
	http.Error(w, "id is required", http.StatusBadRequest)
}


func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person domain.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	createdPerson, err := h.service.CreatePerson(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdPerson)
}

func (h *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var person domain.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	updatedPerson, err := h.service.UpdatePerson(id, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedPerson)
}

func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeletePerson(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
