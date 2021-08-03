package apis

import (
	"encoding/json"
	"io/ioutil"
	"khaibaoyte/control"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

// GetHealthDeclaration godoc
// @Summary Get details health declaration of person
// @Description Get details health declaration of person
// @Tags health-declarations
// @Accept  json
// @Produce  json
// @Param id path string true "Get health declaration with ID"
// @Success 200 {object} entities.HealthDeclaration
// @Router /api/v1/health-declarations/{id} [get]
func GetHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	hd := control.GetHealthDeclaration(idStr)
	json.NewEncoder(w).Encode(hd)
}

// AddHealthDeclaration godoc
// @Summary Add a new health declarations
// @Description Add a new health declaration with the input payload
// @Tags health-declarations
// @Accept  json
// @Produce  json
// @Param healthDeclaration body entities.HealthDeclarationWithoutId true "Add health declaration"
// @Success 200 {object} entities.HealthDeclaration
// @Router /api/v1/health-declarations [post]
func AddHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	add := control.AddHealthDeclaration(hd)
	json.NewEncoder(w).Encode(add)
}

// UpdateHealthDeclaration godoc
// @Summary Update health declaration with id
// @Description Update health declaration with the input payload
// @Tags health-declarations
// @Accept  json
// @Produce  json
// @Param id path string true "ID health declaration need Update"
// @Param healthDeclaration body entities.HealthDeclarationWithoutId true "Update Person"
// @Success 200 {object} entities.HealthDeclaration
// @Router /api/v1/health-declarations/{id} [put]
func UpdateHealthDeclaration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	update := control.UpdateHealthDeclaration(idStr, hd)
	json.NewEncoder(w).Encode(update)
}

// AddQuestion godoc
// @Summary Add a new question
// @Description Add a new question with the input payload
// @Tags questions
// @Accept  json
// @Produce  json
// @Param question body entities.QuestionWithoutId true "Add question"
// @Success 200 {object} entities.Question
// @Router /api/v1/questions [post]
func AddQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var q bson.M
	json.Unmarshal(body, &q)
	add := control.AddQuestion(q)
	json.NewEncoder(w).Encode(add)
}

// GetQuestion godoc
// @Summary Get details question with id
// @Description Get details question with id
// @Tags questions
// @Accept  json
// @Produce  json
// @Param id path string true "Get question with ID"
// @Success 200 {object} entities.Question
// @Router /api/v1/questions/{id} [get]
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	hd := control.GetQuestion(idStr)
	json.NewEncoder(w).Encode(hd)
}

// UpdateQuestion godoc
// @Summary Update question with id
// @Description Update question with the input payload
// @Tags questions
// @Accept  json
// @Produce  json
// @Param id path string true "ID question need Update"
// @Param question body entities.QuestionWithoutId true "Update question"
// @Success 200 {object} entities.Question
// @Router /api/v1/questions/{id} [put]
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	body, _ := ioutil.ReadAll(r.Body)
	var hd bson.M
	json.Unmarshal(body, &hd)
	update := control.UpdateQuestion(idStr, hd)
	json.NewEncoder(w).Encode(update)
}
