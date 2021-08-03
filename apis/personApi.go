package apis

import (
	"encoding/json"
	"io/ioutil"
	"khaibaoyte/control"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

// AddPerson godoc
// @Summary Add a new person information
// @Description Create a new person infor with the input payload
// @Tags persons
// @Accept  json
// @Produce  json
// @Param person body entities.PersonWithoutId true "Add Person"
// @Success 200 {object} entities.Person
// @Router /api/v1/persons [post]
func AddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var person bson.M
	json.Unmarshal(body, &person)
	p := control.AddPerson(person)
	json.NewEncoder(w).Encode(p)
}

// GetPerson godoc
// @Summary Get details information of person
// @Description Get details information of person
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path string true "Get person with ID"
// @Success 200 {object} entities.Person
// @Router /api/v1/persons/{id} [get]
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	log.Println(id)
	person := control.GetPerson(id)
	json.NewEncoder(w).Encode(person)
}

// UpdatePerson godoc
// @Summary Update person information
// @Description Update person infor with the input payload
// @Tags persons
// @Accept  json
// @Produce  json
// @Param id path string true "ID Person need Update"
// @Param person body entities.PersonWithoutId true "Update Person"
// @Success 200 {object} entities.Person
// @Router /api/v1/persons/{id} [put]
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var p bson.M
	json.Unmarshal(body, &p)
	person := control.UpdatePerson(id, p)
	json.NewEncoder(w).Encode(person)
}
