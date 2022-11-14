// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "add new course to the course list",
                "operationId": "create-course",
                "parameters": [
                    {
                        "description": "course",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/course/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "get a course by ID",
                "operationId": "get-course_by_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "update a course by ID",
                "operationId": "update-course-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "delete a course by ID",
                "operationId": "delete-course-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/courses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courses"
                ],
                "summary": "get all items in the course list",
                "operationId": "get-all-courses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CourseList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/person": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "get all person from database",
                "operationId": "get-persons",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PersonListDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "add new person to database",
                "operationId": "add-new-person",
                "parameters": [
                    {
                        "description": "person",
                        "name": "course",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Person"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/person/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "get person by id",
                "operationId": "get-person-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Person"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "update person by id",
                "operationId": "update-person-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person"
                ],
                "summary": "delete person by id",
                "operationId": "delete-person-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Course": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CourseList": {
            "type": "object",
            "properties": {
                "_metadata": {
                    "$ref": "#/definitions/model.PaginationResponse"
                },
                "courses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Course"
                    }
                }
            }
        },
        "model.PaginationResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                }
            }
        },
        "model.Person": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "login",
                "passwd",
                "role_id"
            ],
            "properties": {
                "deleted": {
                    "type": "boolean"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "login": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "passwd": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                },
                "patronymic": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "person_id": {
                    "type": "integer"
                },
                "role_id": {
                    "type": "integer",
                    "maximum": 6,
                    "minimum": 1
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.PersonListDto": {
            "type": "object",
            "properties": {
                "_metadata": {
                    "$ref": "#/definitions/model.PaginationResponse"
                },
                "persons_dto": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Person"
                    }
                }
            }
        },
        "model.ResponseMessage": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
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
	BasePath:    "/api/v1",
	Schemes:     []string{"http"},
	Title:       "GolangUnitedSchool",
	Description: "# _API for student info collection service._",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
