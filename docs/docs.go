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
        "/api/v1/auth/login": {
            "post": {
                "description": "Login Handles user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "User credentials for login",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.UserSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create session",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "description": "Handles user logout.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "Logout successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signup": {
            "post": {
                "description": "Handles user signup.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "New user details for signup",
                        "name": "newUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.UserSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Signup successful",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to add user",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/email/delete/{id}": {
            "delete": {
                "description": "Delete an email message based on its identifier",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Delete an email message",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the email message",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deletion success status",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad id",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to delete email message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/email/send": {
            "post": {
                "description": "Send a new email message to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Send a new email message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Email message in JSON format",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.EmailSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID of the send email message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad JSON in request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to add email message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/email/update/{id}": {
            "put": {
                "description": "Update an existing email message based on its identifier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Update an email message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID of the email message",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Email message in JSON format",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.EmailSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update success status",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad id or Bad JSON",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to update email message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/email/{id}": {
            "get": {
                "description": "Get an email message by its unique identifier",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Get an email message by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the email message",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email message data",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad id in request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "Email not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/emails/draft": {
            "get": {
                "description": "Get a list of all email messages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Display the list of email messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of all email messages",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "DB error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "JSON encoding error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/emails/incoming": {
            "get": {
                "description": "Get a list of all email messages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Display the list of email messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of all email messages",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "DB error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "JSON encoding error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/emails/sent": {
            "get": {
                "description": "Get a list of all email messages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Display the list of email messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of all email messages",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "DB error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "JSON encoding error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/emails/spam": {
            "get": {
                "description": "Get a list of all email messages",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emails"
                ],
                "summary": "Display the list of email messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of all email messages",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "DB error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "JSON encoding error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/folder/add": {
            "post": {
                "description": "Add a new folder message to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "folders"
                ],
                "summary": "Add a new folder message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Folder message in JSON format",
                        "name": "folder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.FolderSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID of the send folder message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad JSON in request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to add folder message",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/avatar/delete": {
            "delete": {
                "description": "Handles requests to delete user avatar.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user avatar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User avatar deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/avatar/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Handles requests to upload user avatar.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Upload user avatar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Avatar file to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "File uploaded and saved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Error processing file or failed to get file",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/delete/{id}": {
            "delete": {
                "description": "Handles requests to delete user data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User data deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/get": {
            "get": {
                "description": "Retrieve the user associated with the current session",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/update": {
            "put": {
                "description": "Handles requests to update user data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "CSRF Token",
                        "name": "X-Csrf-Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Updated user data",
                        "name": "updatedUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.UserSwag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User data updated successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Not authorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/verify-auth": {
            "get": {
                "description": "Verify user authentication using sessions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Verify user authentication",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.EmailSwag": {
            "type": "object",
            "properties": {
                "dateOfDispatch": {
                    "description": "DateOfDispatch is the date when the email was sent.",
                    "type": "string"
                },
                "deleted": {
                    "description": "Deleted indicates whether the email has been deleted.",
                    "type": "boolean"
                },
                "draftStatus": {
                    "description": "DraftStatus indicates whether the email is a draft.",
                    "type": "boolean"
                },
                "id": {
                    "description": "ID is the unique identifier of the email in the database.",
                    "type": "integer"
                },
                "mark": {
                    "description": "Mark is a flag, such as marking the email as a favorite.",
                    "type": "boolean"
                },
                "photoId": {
                    "description": "PhotoID is the link to the photo attached to the email, if any.",
                    "type": "string"
                },
                "readStatus": {
                    "description": "ReadStatus indicates whether the email has been read.",
                    "type": "boolean"
                },
                "recipientEmail": {
                    "description": "RecipientEmail is the Email of the recipient user",
                    "type": "string"
                },
                "replyToEmailId": {
                    "description": "ReplyToEmailID is the ID of the email to which a reply can be sent.",
                    "type": "integer"
                },
                "senderEmail": {
                    "description": "SenderEmail is the Email of the sender user",
                    "type": "string"
                },
                "spamStatus": {
                    "description": "SpamStatus indicates whether the email is a spam",
                    "type": "boolean"
                },
                "text": {
                    "description": "Text is the body of the email.",
                    "type": "string"
                },
                "topic": {
                    "description": "Topic is the subject of the email.",
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.FolderSwag": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID he unique ID of the folder in the database.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name the name of the folder.",
                    "type": "string"
                },
                "profileId": {
                    "description": "ProfileId the unique identifier of the user who owns the folder.",
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "body": {},
                "status": {
                    "type": "integer"
                }
            }
        },
        "response.UserGenderSwag": {
            "type": "string",
            "enum": [
                "Male",
                "Female",
                "Other"
            ],
            "x-enum-varnames": [
                "Male",
                "Female",
                "Other"
            ]
        },
        "response.UserSwag": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "AvatarID stores the identifier of the user's avatar image.",
                    "type": "string"
                },
                "birthday": {
                    "description": "Birthday stores the birthdate of the user.",
                    "type": "string"
                },
                "description": {
                    "description": "Description stores additional information about the user.",
                    "type": "string"
                },
                "firstname": {
                    "description": "FirstName stores the first name of the user.",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender stores the gender of the user.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/response.UserGenderSwag"
                        }
                    ]
                },
                "id": {
                    "description": "ID uniquely identifies the user.",
                    "type": "integer"
                },
                "login": {
                    "description": "Login is the username used for authentication.",
                    "type": "string"
                },
                "middlename": {
                    "description": "Patronymic stores the middle name of the user, if available.",
                    "type": "string"
                },
                "password": {
                    "description": "Password is the hashed password of the user.",
                    "type": "string"
                },
                "phonenumber": {
                    "description": "PhoneNumber stores the phone number of the user.",
                    "type": "string"
                },
                "surname": {
                    "description": "Surname stores the last name of the user.",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Mail",
	Description:      "API server for mail",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
