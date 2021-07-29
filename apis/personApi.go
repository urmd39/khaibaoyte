package apis

import (
	"encoding/json"
	"io/ioutil"
	"khaibaoyte/control"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

func AddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var person bson.M
	json.Unmarshal(body, &person)
	p := control.AddPerson(person)
	json.NewEncoder(w).Encode(p)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	person := control.GetPerson(id)
	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var p bson.M
	json.Unmarshal(body, &p)
	person := control.UpdatePerson(id, p)
	json.NewEncoder(w).Encode(person)
}
