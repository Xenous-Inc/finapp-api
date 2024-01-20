// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/finapp/api/auditorium": {
            "post": {
                "description": "get auditorium",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetAuditorium",
                "parameters": [
                    {
                        "description": "auditorium info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetAuditoriumInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/auditorium/schedule": {
            "post": {
                "description": "get auditorium schedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetAuditoriumSchedule",
                "parameters": [
                    {
                        "description": "auditorium schedule info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetAuditoriumScheduleInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/auth": {
            "post": {
                "description": "auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrgFaRu"
                ],
                "summary": "GetGroup",
                "parameters": [
                    {
                        "description": "auth",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orgfaclient.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/auth/mygroup": {
            "get": {
                "description": "get myGroup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrgFaRu"
                ],
                "summary": "GetMyGroup",
                "parameters": [
                    {
                        "description": "myGroup info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orgfaclient.AuthSession"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/group": {
            "post": {
                "description": "get group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetGroup",
                "parameters": [
                    {
                        "description": "group info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetGroupsInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/group/schedule": {
            "post": {
                "description": "get groupSchedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetGroupSchedule",
                "parameters": [
                    {
                        "description": "group schedule info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetGroupScheduleInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/teacher": {
            "post": {
                "description": "get teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetTeacher",
                "parameters": [
                    {
                        "description": "teacher info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetTeacherInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/finapp/api/teacher/schedule": {
            "post": {
                "description": "get teacherSchedule",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TimeTable"
                ],
                "summary": "GetTeacherSchedule",
                "parameters": [
                    {
                        "description": "teacher schedule info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ruzfaclient.GetTeacherScheduleInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "orgfaclient.AuthSession": {
            "type": "object",
            "properties": {
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "orgfaclient.LoginInput": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetAuditoriumInput": {
            "type": "object",
            "properties": {
                "auditorium": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetAuditoriumScheduleInput": {
            "type": "object",
            "properties": {
                "auditoriumId": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetGroupScheduleInput": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "groupId": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetGroupsInput": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetTeacherInput": {
            "type": "object",
            "properties": {
                "teacher": {
                    "type": "string"
                }
            }
        },
        "ruzfaclient.GetTeacherScheduleInput": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "teacherId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:5555",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Finapp-api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}