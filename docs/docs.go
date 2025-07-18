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
        "/api/book": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "根据请求体中的书籍信息创建新书籍",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍管理"
                ],
                "summary": "创建新书籍",
                "parameters": [
                    {
                        "description": "书籍信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/book/list": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "分页获取书籍信息列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍管理"
                ],
                "summary": "获取书籍列表",
                "parameters": [
                    {
                        "description": "查询参数（如分页）",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BookQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/utils.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.Book"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/book/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "根据书籍 ID 获取详细信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍管理"
                ],
                "summary": "获取书籍详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "书籍ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Book"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "根据书籍 ID 更新书籍数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍管理"
                ],
                "summary": "更新书籍信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "书籍ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新的书籍信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "根据书籍 ID 删除对应书籍",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "书籍管理"
                ],
                "summary": "删除书籍",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "书籍ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "使用用户名和密码进行登录，返回 JWT Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/user/admin": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "仅角色为 admin 的用户可访问",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "仅管理员可访问",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "将当前 JWT 加入黑名单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/user/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "返回当前登录用户的用户名与角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/api/user/write": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "需要拥有 write 权限的用户才可以访问",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "写数据",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code 表示响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "Data 表示响应携带的数据"
                },
                "message": {
                    "description": "Message 表示响应消息",
                    "type": "string"
                }
            }
        },
        "controller.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Book": {
            "type": "object",
            "required": [
                "author",
                "price",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "maxLength": 100
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 1000
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string",
                    "maxLength": 200
                },
                "updated_at": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.BookQuery": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者模糊查询",
                    "type": "string",
                    "maxLength": 100
                },
                "page": {
                    "description": "当前页码，默认1",
                    "type": "integer",
                    "minimum": 1
                },
                "page_size": {
                    "description": "每页条数，默认10，最大100",
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "title": {
                    "description": "书名模糊查询",
                    "type": "string",
                    "maxLength": 200
                }
            }
        },
        "utils.PageResult": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据列表，通常是切片，如 []Book"
                },
                "page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "page_size": {
                    "description": "每页条数",
                    "type": "integer"
                },
                "total": {
                    "description": "总记录数",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Gin JWT RBAC Demo API",
	Description:      "示例项目：JWT + RBAC + Redis 黑名单 + Swagger 接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
