package apis

import (
	"encoding/json"
	"khaibaoyte/control"
	"net/http"
)

func ListProvince(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListProvince()
	json.NewEncoder(w).Encode(list)
}

func ListTown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	province := r.FormValue("province")
	list := control.GetListTown(province)
	json.NewEncoder(w).Encode(list)
}

func ListVillage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	town := r.FormValue("town")
	list := control.GetListVillage(town)
	json.NewEncoder(w).Encode(list)
}

func ListSex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListSex()
	json.NewEncoder(w).Encode(list)
}

func ListNationality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListNationality()
	json.NewEncoder(w).Encode(list)
}

func ListPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListPerson()
	json.NewEncoder(w).Encode(list)
}

func ListQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListQuestion()
	json.NewEncoder(w).Encode(list)
}

func ListHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetListHealthDeclaration()
	json.NewEncoder(w).Encode(list)
}
