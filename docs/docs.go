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
            "name": "Afin Tachtiar",
            "email": "afintchtr@gmail.com"
        },
        "license": {
            "name": "AfinT",
            "url": "https://github.com/afintchtr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments/": {
            "get": {
                "description": "Get details of all comments",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Get all comments details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "delete": {
                "description": "Just delete an existing comment",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Delete an existing comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update an existing comment with new message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Update an existing comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{id}": {
            "get": {
                "description": "Get details of one comment by its id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Get comment details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comments to be shown",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{photoId}": {
            "post": {
                "description": "Create a new comment with message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "Create a new comment",
                "parameters": [
                    {
                        "description": "create comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/photos/": {
            "get": {
                "description": "Get details of all photos",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Get all photo details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new photo with title, caption and url",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Create a new photo",
                "parameters": [
                    {
                        "description": "create photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{id}": {
            "delete": {
                "description": "Just delete an existing photo",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Delete an existing photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update an existing photo with new title, caption, and url",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Update an existing photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/social-medias/": {
            "get": {
                "description": "Get details of all social medias",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "social-media"
                ],
                "summary": "Get all social media details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new social media with name and url",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "social-media"
                ],
                "summary": "Create a new social media",
                "parameters": [
                    {
                        "description": "create social media",
                        "name": "models.SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-medias/{id}": {
            "get": {
                "description": "Get details of one social media by its id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "social-media"
                ],
                "summary": "Get social media details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the social media to be shown",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "delete": {
                "description": "Just delete an existing social media",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "social-media"
                ],
                "summary": "Delete an existing social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the social media to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update an existing social media with new name and url",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "social-media"
                ],
                "summary": "Update an existing social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the social media to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/users/": {
            "get": {
                "description": "Get details of all users",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all user details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    }
                }
            }
        },
        "/users/login/": {
            "post": {
                "description": "Authenticate existing user and generate token",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login to existing account",
                "parameters": [
                    {
                        "description": "get token",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/register/": {
            "post": {
                "description": "Create a new user with email and password",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photoID": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "social_medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "social_medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8085",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyGram API (Afin)",
	Description:      "This is what ive been practiced in Golang class Bootcamp BRI Golang for Final project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
