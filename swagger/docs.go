// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

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
        "contact": {
            "name": "radyatama",
            "email": "mohradyatama24@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/crawl/web": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crawl"
                ],
                "summary": "CrawlWeb",
                "parameters": [
                    {
                        "type": "string",
                        "description": "lang",
                        "name": "Accept-Language",
                        "in": "header"
                    },
                    {
                        "description": "request payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CrawlRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagger.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.CrawlResponse"
                                        },
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "type": "object"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagger.BadRequestErrorValidationResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/swagger.ValidationErrors"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagger.RequestTimeoutResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "type": "object"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/swagger.InternalServerErrorResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "errors": {
                                            "type": "array",
                                            "items": {
                                                "type": "object"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CrawlRequest": {
            "type": "object",
            "required": [
                "web_url"
            ],
            "properties": {
                "web_url": {
                    "type": "string"
                }
            }
        },
        "domain.CrawlResponse": {
            "type": "object",
            "properties": {
                "h_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.HTags"
                    }
                },
                "meta_content": {
                    "$ref": "#/definitions/domain.MetaContent"
                },
                "source_code_html_url": {
                    "type": "string"
                }
            }
        },
        "domain.HTags": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tags": {
                    "type": "string"
                }
            }
        },
        "domain.MetaContent": {
            "type": "object",
            "properties": {
                "meta_description": {
                    "type": "string"
                },
                "meta_keywords": {
                    "type": "string"
                },
                "meta_title": {
                    "type": "string"
                }
            }
        },
        "swagger.BadRequestErrorValidationResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "KDMU-02-006"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string",
                    "example": "permintaan tidak valid, kesalahan muncul ketika permintaan Anda memiliki parameter yang tidak valid."
                },
                "request_id": {
                    "type": "string",
                    "example": "24fa3770-628c-49de-aa17-3a338f73d99b"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2022-04-27 23:19:56"
                }
            }
        },
        "swagger.BaseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "OK"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string",
                    "example": "operasi berhasil dieksekusi."
                },
                "request_id": {
                    "type": "string",
                    "example": "24fa3770-628c-49de-aa17-3a338f73d99b"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2022-04-27 23:19:56"
                }
            }
        },
        "swagger.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "KDMU-02-008"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string",
                    "example": "terjadi kesalahan, silakan hubungi administrator."
                },
                "request_id": {
                    "type": "string",
                    "example": "24fa3770-628c-49de-aa17-3a338f73d99b"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2022-04-27 23:19:56"
                }
            }
        },
        "swagger.RequestTimeoutResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "KDMU-02-009"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string",
                    "example": "permintaan telah melampaui batas waktu, harap request kembali."
                },
                "request_id": {
                    "type": "string",
                    "example": "24fa3770-628c-49de-aa17-3a338f73d99b"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2022-04-27 23:19:56"
                }
            }
        },
        "swagger.ValidationErrors": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "MobilePhone wajib diisi."
                },
                "message": {
                    "type": "string",
                    "example": "ActiveDate harus format yang benar yyyy-mm-dd."
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "v1",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Api Gateway V1",
	Description: "api \"API Gateway v1\"",
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
