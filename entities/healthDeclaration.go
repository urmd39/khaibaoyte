package entities

// Province represents the model for province
type Province struct {
	ProvinceID string `json:"_id"`
	Name       string `json:"name"`
}

// Town represents the model for town
type Town struct {
	TownID   string `json:"_id"`
	Name     string `json:"name"`
	Province string `json:"province"`
}

// Village represents the model for village
type Village struct {
	VillageID string `json:"_id"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	Town      string `json:"town"`
}

// Gender represents the model for gender
type Gender struct {
	GenderID string `json:"_id"`
	Name     string `json:"name"`
}

// Nationality represents the model for nationality
type Nationality struct {
	NationalityID string `json:"_id"`
	Name          string `json:"name"`
}

// Address represents the model for full address information
type AddressFull struct {
	Province string `json:"province"`
	Town     string `json:"town"`
	Village  string `json:"village"`
	Street   string `json:"street"`
}

type Address struct {
	Province string `json:"province"`
	Town     string `json:"town"`
	Village  string `json:"village"`
}

// Person represents the model for person information
type Person struct {
	PersonID     string `json:"_id"`
	Name         string `json:"name"`
	Gender       string `json:"sex"`
	Birthday     string `json:"birthday"`
	Nationality  string `json:"nationality"`
	Address      AddressFull
	IdentityCard string `json:"identityCard"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
}

// Person represents the model for person information without id
type PersonWithoutId struct {
	Name         string `json:"name"`
	Gender       string `json:"sex"`
	Birthday     string `json:"birthday"`
	Nationality  string `json:"nationality"`
	Address      AddressFull
	IdentityCard string `json:"identityCard"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
}

// Question represents the model for question
type Question struct {
	QuestionID string   `json:"_id"`
	Name       string   `json:"question"`
	Answer     []string `json:"answer"`
}

// Question represents the model for question without id
type QuestionWithoutId struct {
	Name   string   `json:"question"`
	Answer []string `json:"answer"`
}

// HealthDeclaration represents the model for health declaration
type HealthDeclaration struct {
	HealthDeclarationID string            `json:"_id"`
	PersonID            string            `json:"personId"`
	Note                string            `json:"note"`
	CreatedOn           string            `json:"createdOn"`
	HealthDeclarations  []Question        `json:"healthDeclaration"`
	TravelItinerarys    []TravelItinerary `json:"travelItinerary"`
}

// HealthDeclaration represents the model for health declaration without Id
type HealthDeclarationWithoutId struct {
	PersonID           string            `json:"personId"`
	Note               string            `json:"note"`
	CreatedOn          string            `json:"createdOn"`
	HealthDeclarations []Question        `json:"healthDeclaration"`
	TravelItinerarys   []TravelItinerary `json:"travelItinerary"`
}

// TravelItinerary represents the model for travel itinerary
type TravelItinerary struct {
	DayStart string `json:"dayStart"`
	DayEnd   string `json:"dayEnd"`
	Address  Address
}
