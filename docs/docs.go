// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "In success case returns Access JWT Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Try to sign in user",
                "parameters": [
                    {
                        "description": "Credentials input",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/LoginResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        },
        "/classrooms/": {
            "get": {
                "description": "Return list of classrooms which found by provided query term",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classrooms"
                ],
                "summary": "Return List of classrooms",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Classroom search mask",
                        "name": "term",
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
                                "$ref": "#/definitions/Clasroom"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        },
        "/groups/": {
            "get": {
                "description": "Return list of student groups which found by provided query term",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Return List of groups",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Group search mask",
                        "name": "term",
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
                                "$ref": "#/definitions/Group"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        },
        "/teachers/": {
            "get": {
                "description": "Return list of teachers which found by provided query term",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teachers"
                ],
                "summary": "Return List of teachers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Teacher search mask",
                        "name": "term",
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
                                "$ref": "#/definitions/Teacher"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ApiError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Validation error"
                }
            }
        },
        "Clasroom": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string",
                    "example": "25 | Филиалы Челябинский филиал | Лекционная"
                },
                "id": {
                    "type": "string",
                    "example": "3587"
                },
                "title": {
                    "type": "string",
                    "example": "24"
                }
            }
        },
        "Group": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string",
                    "example": "Финансовый факультет | Очная"
                },
                "id": {
                    "type": "string",
                    "example": "110694"
                },
                "title": {
                    "type": "string",
                    "example": "ФФ21-1"
                }
            }
        },
        "LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "Bibt8877"
                },
                "username": {
                    "type": "string",
                    "example": "227789"
                }
            }
        },
        "LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9"
                }
            }
        },
        "Teacher": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string",
                    "example": "Департамент гуманитарных наук"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0001-2345-6789-000000005451"
                },
                "title": {
                    "type": "string",
                    "example": "Махаматов Таир"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5051/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
