package apis

import (
	"encoding/json"
	"khaibaoyte/control"
	"net/http"
)

// GetProvinces godoc
// @Summary Get details of all provinces
// @Description Get details of all provinces
// @Tags provinces
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Province
// @Router /api/v1/provinces [get]
func GetProvinces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetProvinces()
	json.NewEncoder(w).Encode(list)
}

// GetTowns godoc
// @Summary Get details of all towns
// @Description Get details of all towns
// @Tags towns
// @Accept  json
// @Produce  json
// @Param province query string true "Get towns in province"
// @Success 200 {array} entities.Town
// @Router /api/v1/towns [get]
func GetTowns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	province := r.FormValue("province")
	list := control.GetTowns(province)
	json.NewEncoder(w).Encode(list)
}

// GetVillages godoc
// @Summary Get details of all villages
// @Description Get details of all villages
// @Tags villages
// @Accept  json
// @Produce  json
// @Param town query string true "Get villages in town"
// @Success 200 {array} entities.Village
// @Router /api/v1/villages [get]
func GetVillages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	town := r.FormValue("town")
	list := control.GetVillages(town)
	json.NewEncoder(w).Encode(list)
}

// GetGenders godoc
// @Summary Get details of all genders
// @Description Get details of all genders
// @Tags genders
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Gender
// @Router /api/v1/genders [get]
func GetGenders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetGenders()
	json.NewEncoder(w).Encode(list)
}

// GetNationalitys godoc
// @Summary Get details of all nationalitys
// @Description Get details of all nationalitys
// @Tags nationalitys
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Nationality
// @Router /api/v1/nationalitys [get]
func GetNationalitys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetNationalitys()
	json.NewEncoder(w).Encode(list)
}

// GetPersons godoc
// @Summary Get details information of all persons
// @Description Get details information of all persons
// @Tags persons
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Person
// @Router /api/v1/persons [get]
func GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetPersons()
	json.NewEncoder(w).Encode(list)
}

// GetQuestions godoc
// @Summary Get details of all questions
// @Description Get details of all questions
// @Tags questions
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.Question
// @Router /api/v1/questions [get]
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetQuestions()
	json.NewEncoder(w).Encode(list)
}

// GetHealthDeclaration godoc
// @Summary Get details of all health declarations
// @Description Get details of all health declarations
// @Tags health-declarations
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.HealthDeclaration
// @Router /api/v1/health-declarations [get]
func GetHealthDeclarations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetHealthDeclarations()
	json.NewEncoder(w).Encode(list)
}
