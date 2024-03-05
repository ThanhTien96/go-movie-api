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
        "/genres": {
            "get": {
                "description": "Get Genres",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Get Genres",
                "operationId": "get-genres",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "sorts",
                        "name": "sorts",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create Genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Create Genre",
                "operationId": "create-genre",
                "parameters": [
                    {
                        "description": "GenreCreateRequest",
                        "name": "GenreCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateGenreRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Genres",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Delete Genres",
                "operationId": "delete-genres",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "gen_ids",
                        "name": "gen_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/genres/{gen_id}": {
            "get": {
                "description": "Get Genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Get Genre",
                "operationId": "get-genre",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "gen_id",
                        "name": "gen_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Update Genre",
                "operationId": "update-genre",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "gen_id",
                        "name": "gen_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "GenreCreateRequest",
                        "name": "GenreCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateGenreRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Genre"
                ],
                "summary": "Delete Genre",
                "operationId": "delete-genre",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "gen_id",
                        "name": "gen_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movies": {
            "get": {
                "description": "Get Movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get Movies",
                "operationId": "get-movies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "start_date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "end_date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "sorts",
                        "name": "sorts",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "filters",
                        "name": "filters",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Create Movie",
                "operationId": "create-movie",
                "parameters": [
                    {
                        "description": "MovieCreateRequest",
                        "name": "MovieCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateMovieRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Delete Movies",
                "operationId": "delete-movies",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "mov_ids",
                        "name": "mov_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movies/{mov_id}": {
            "get": {
                "description": "Get Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get Movie",
                "operationId": "get-movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "mov_id",
                        "name": "mov_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Update Movie",
                "operationId": "update-movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "mov_id",
                        "name": "mov_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "MovieCreateRequest",
                        "name": "MovieCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateMovieRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Delete Movie",
                "operationId": "delete-movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "mov_id",
                        "name": "mov_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.CreateGenreRequest": {
            "type": "object",
            "properties": {
                "gen_title": {
                    "type": "string"
                }
            }
        },
        "models.CreateMovieRequest": {
            "type": "object",
            "properties": {
                "mov_dt_rel": {
                    "type": "integer"
                },
                "mov_lang": {
                    "type": "string"
                },
                "mov_rel_country": {
                    "type": "string"
                },
                "mov_time": {
                    "type": "integer"
                },
                "mov_title": {
                    "type": "string"
                },
                "mov_year": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
