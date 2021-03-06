// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/genders": {
            "get": {
                "description": "Get details of all genders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "genders"
                ],
                "summary": "Get details of all genders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Gender"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/health-declarations": {
            "get": {
                "description": "Get details of all health declarations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health-declarations"
                ],
                "summary": "Get details of all health declarations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.HealthDeclaration"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new health declaration with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health-declarations"
                ],
                "summary": "Add a new health declarations",
                "parameters": [
                    {
                        "description": "Add health declaration",
                        "name": "healthDeclaration",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.HealthDeclarationWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.HealthDeclaration"
                        }
                    }
                }
            }
        },
        "/api/v1/health-declarations/{id}": {
            "get": {
                "description": "Get details health declaration of person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health-declarations"
                ],
                "summary": "Get details health declaration of person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get health declaration with ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.HealthDeclaration"
                        }
                    }
                }
            },
            "put": {
                "description": "Update health declaration with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health-declarations"
                ],
                "summary": "Update health declaration with id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID health declaration need Update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Person",
                        "name": "healthDeclaration",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.HealthDeclarationWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.HealthDeclaration"
                        }
                    }
                }
            }
        },
        "/api/v1/nationalitys": {
            "get": {
                "description": "Get details of all nationalitys",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nationalitys"
                ],
                "summary": "Get details of all nationalitys",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Nationality"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/persons": {
            "get": {
                "description": "Get details information of all persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Get details information of all persons",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Person"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new person infor with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Add a new person information",
                "parameters": [
                    {
                        "description": "Add Person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.PersonWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Person"
                        }
                    }
                }
            }
        },
        "/api/v1/persons/{id}": {
            "get": {
                "description": "Get details information of person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Get details information of person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get person with ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Person"
                        }
                    }
                }
            },
            "put": {
                "description": "Update person infor with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Update person information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID Person need Update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.PersonWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Person"
                        }
                    }
                }
            }
        },
        "/api/v1/provinces": {
            "get": {
                "description": "Get details of all provinces",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provinces"
                ],
                "summary": "Get details of all provinces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Province"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/questions": {
            "get": {
                "description": "Get details of all questions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Get details of all questions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Question"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new question with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Add a new question",
                "parameters": [
                    {
                        "description": "Add question",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.QuestionWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Question"
                        }
                    }
                }
            }
        },
        "/api/v1/questions/{id}": {
            "get": {
                "description": "Get details question with id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Get details question with id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get question with ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Question"
                        }
                    }
                }
            },
            "put": {
                "description": "Update question with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Update question with id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID question need Update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update question",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.QuestionWithoutId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Question"
                        }
                    }
                }
            }
        },
        "/api/v1/towns": {
            "get": {
                "description": "Get details of all towns",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "towns"
                ],
                "summary": "Get details of all towns",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get towns in province",
                        "name": "province",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Town"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/villages": {
            "get": {
                "description": "Get details of all villages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "villages"
                ],
                "summary": "Get details of all villages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get villages in town",
                        "name": "town",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Village"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Address": {
            "type": "object",
            "properties": {
                "province": {
                    "type": "string"
                },
                "town": {
                    "type": "string"
                },
                "village": {
                    "type": "string"
                }
            }
        },
        "entities.AddressFull": {
            "type": "object",
            "properties": {
                "province": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "town": {
                    "type": "string"
                },
                "village": {
                    "type": "string"
                }
            }
        },
        "entities.Gender": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.HealthDeclaration": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "createdOn": {
                    "type": "string"
                },
                "healthDeclaration": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Question"
                    }
                },
                "note": {
                    "type": "string"
                },
                "personId": {
                    "type": "string"
                },
                "travelItinerary": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.TravelItinerary"
                    }
                }
            }
        },
        "entities.HealthDeclarationWithoutId": {
            "type": "object",
            "properties": {
                "createdOn": {
                    "type": "string"
                },
                "healthDeclaration": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Question"
                    }
                },
                "note": {
                    "type": "string"
                },
                "personId": {
                    "type": "string"
                },
                "travelItinerary": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.TravelItinerary"
                    }
                }
            }
        },
        "entities.Nationality": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Person": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "address": {
                    "$ref": "#/definitions/entities.AddressFull"
                },
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "identityCard": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "entities.PersonWithoutId": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/entities.AddressFull"
                },
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "identityCard": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        },
        "entities.Province": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Question": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "answer": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "entities.QuestionWithoutId": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "entities.Town": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                }
            }
        },
        "entities.TravelItinerary": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/entities.Address"
                },
                "dayEnd": {
                    "type": "string"
                },
                "dayStart": {
                    "type": "string"
                }
            }
        },
        "entities.Village": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "town": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Health Declarations API",
	Description: "This is a service for health declaration",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
