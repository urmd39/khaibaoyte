package apis

import (
	"encoding/json"
	"io/ioutil"
	"khaibaoyte/control"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

func GetHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	hd := control.GetHealthDeclaration(idStr)
	json.NewEncoder(w).Encode(hd)
}

func AddHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	add := control.AddHealthDeclaration(hd)
	json.NewEncoder(w).Encode(add)
}

func UpdateHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	update := control.UpdateHealthDeclaration(idStr, hd)
	json.NewEncoder(w).Encode(update)
}

func AddQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var q bson.M
	json.Unmarshal(body, &q)
	add := control.AddQuestion(q)
	json.NewEncoder(w).Encode(add)
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	hd := control.GetQuestion(idStr)
	json.NewEncoder(w).Encode(hd)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	update := control.UpdateQuestion(idStr, hd)
	json.NewEncoder(w).Encode(update)
}
