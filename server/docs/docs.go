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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alga/add": {
            "post": {
                "description": "添加藻类图片数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "AddAlga",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/alga/anno": {
            "get": {
                "description": "藻类图像获取标注",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetAnnotationByAlga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "藻类名称",
                        "name": "alga",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/alga/search": {
            "get": {
                "description": "搜索藻类图像",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "SearchAlga",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/anno/add": {
            "post": {
                "description": "添加标注",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "AddAnnotation",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/anno/delete": {
            "get": {
                "description": "删除标注",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "DeleteAnnotation",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/anno/update": {
            "post": {
                "description": "修改标注信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "UpdateAnnotation",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/data": {
            "get": {
                "description": "获取所有藻类数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetData",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/river/add": {
            "post": {
                "description": "添加河流数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "AddRiver",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/river/all": {
            "get": {
                "description": "获取所有河流数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetRivers",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "获取所有用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetUsers",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/anno": {
            "get": {
                "description": "用户获取个人标注",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetAnnotationByUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "user",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "get": {
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "DeleteUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "GetUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/pwd": {
            "post": {
                "description": "修改用户密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "ChangePassword",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "description": "修改用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aldb"
                ],
                "summary": "UpdateUser",
                "responses": {
                    "200": {
                        "description": "{code, msg, data}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Aldb Example API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
