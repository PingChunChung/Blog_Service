// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "取得多個標籤",
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "頁碼",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每頁數量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增標籤",
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "建立者",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改標籤",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "標籤ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "標籤名稱",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "狀態",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "刪除標籤",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "標籤ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    },
                    "400": {
                        "description": "請求錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    },
                    "500": {
                        "description": "內部錯誤",
                        "schema": {
                            "$ref": "#/definitions/errcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errcode.Error": {
            "type": "object"
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "Modified_by": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "deleted_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "boolean"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "部落格系統",
	Description:      "Golang webApplication練習",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}