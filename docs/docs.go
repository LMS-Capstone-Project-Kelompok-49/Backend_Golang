// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"github.com/swaggo/swag"
)


// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/app",
	Schemes:          []string{},
	Title:            "Docs",
	Description:      "Swagger docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

const docTemplate = `
{
  "swagger": "2.0",
  "info": {
    "version": "1.0",
    "title": "Capstone LMS",
    "contact": {}
  },
  "host": "3.95.181.246",
  "basePath": "/",
  "securityDefinitions": {},
  "schemes": [
    "http"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/auth/login": {
      "post": {
        "summary": "Login",
        "tags": [
          "Auth"
        ],
        "operationId": "Login",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "summary": "Register",
        "tags": [
          "Auth"
        ],
        "operationId": "Register",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/RegisterRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/user/all": {
      "get": {
        "summary": "Get All Users",
        "tags": [
          "User"
        ],
        "operationId": "GetAllUsers",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/user/2": {
      "get": {
        "summary": "Get One User",
        "tags": [
          "User"
        ],
        "operationId": "GetOneUser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      },
      "put": {
        "summary": "Update User",
        "tags": [
          "User"
        ],
        "operationId": "UpdateUser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/UpdateUserRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      },
      "delete": {
        "summary": "Delete User",
        "tags": [
          "User"
        ],
        "operationId": "DeleteUser",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course/all": {
      "get": {
        "summary": "Get All Course",
        "tags": [
          "Course"
        ],
        "operationId": "GetAllCourse",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course/2": {
      "get": {
        "summary": "Get All Course1",
        "tags": [
          "Course"
        ],
        "operationId": "GetAllCourse1",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course/create": {
      "post": {
        "summary": "Post Course",
        "tags": [
          "Course"
        ],
        "operationId": "PostCourse",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/PostCourseRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course/delete/2": {
      "delete": {
        "summary": "Delete Coourse",
        "tags": [
          "Course"
        ],
        "operationId": "DeleteCoourse",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/role/2": {
      "get": {
        "summary": "Get Roles",
        "tags": [
          "Roles"
        ],
        "operationId": "GetRoles",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/type_course": {
      "get": {
        "summary": "Get Type Course",
        "tags": [
          "Type Course"
        ],
        "operationId": "GetTypeCourse",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/type_course/2": {
      "get": {
        "summary": "Get One Type Course",
        "tags": [
          "Type Course"
        ],
        "operationId": "GetOneTypeCourse",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course_category": {
      "get": {
        "summary": "Get All Category",
        "tags": [
          "Course Category"
        ],
        "operationId": "GetAllCategory",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course_category/2": {
      "get": {
        "summary": "Get One Category",
        "tags": [
          "Course Category"
        ],
        "operationId": "GetOneCategory",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course_category/create": {
      "post": {
        "summary": "Post Category",
        "tags": [
          "Course Category"
        ],
        "operationId": "PostCategory",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/PostCategoryRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        },
        "security": []
      }
    },
    "/material/course/1": {
      "get": {
        "summary": "get material",
        "tags": [
          "Materials"
        ],
        "operationId": "getmaterial",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/material/create/1": {
      "post": {
        "summary": "post material",
        "tags": [
          "Materials"
        ],
        "operationId": "postmaterial",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
          {
            "name": "materialname",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "ppt",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "video",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/api/course/delete/1": {
      "delete": {
        "summary": "create_course",
        "tags": [
          "course"
        ],
        "operationId": "Deletecreate_course",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "x-api-key",
            "in": "header",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        },
        "security": []
      }
    },
    "/api/type_course/1": {
      "get": {
        "summary": "type_course",
        "tags": [
          "type_course"
        ],
        "operationId": "type_course",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/role/create": {
      "post": {
        "summary": "post_role Copy",
        "tags": [
          "lokal"
        ],
        "operationId": "post_roleCopy",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Body",
            "in": "body",
            "required": true,
            "description": "",
            "schema": {
              "$ref": "#/definitions/post_roleCopyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        },
        "security": []
      }
    }
  },
  "definitions": {
    "LoginRequest": {
      "title": "LoginRequest",
      "example": {
        "email": "user@mail.com",
        "password": "user3241"
      },
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "required": [
        "email",
        "password"
      ]
    },
    "RegisterRequest": {
      "title": "RegisterRequest",
      "example": {
        "name": "user",
        "email": "user@mail.com",
        "password": "user3241"
      },
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "required": [
        "name",
        "email",
        "password"
      ]
    },
    "UpdateUserRequest": {
      "title": "UpdateUserRequest",
      "example": {
        "name": "mamang",
        "email": "kesbor@mail.com",
        "password": "dik3sMar2"
      },
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "required": [
        "name",
        "email",
        "password"
      ]
    },
    "PostCourseRequest": {
      "title": "PostCourseRequest",
      "example": {
        "coursetype": "lifetime",
        "category": "website",
        "coursename": "Dasar Basic Wordpress",
        "description": "Dalam course ini, anda akan mempelajari.......",
        "courseprice": 1000000
      },
      "type": "object",
      "properties": {
        "coursetype": {
          "type": "string"
        },
        "category": {
          "type": "string"
        },
        "coursename": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "courseprice": {
          "type": "integer",
          "format": "int32"
        }
      },
      "required": [
        "coursetype",
        "category",
        "coursename",
        "description",
        "courseprice"
      ]
    },
    "PostCategoryRequest": {
      "title": "PostCategoryRequest",
      "example": {
        "category": "Desktop"
      },
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        }
      },
      "required": [
        "category"
      ]
    },
    "loginCopyRequest": {
      "title": "loginCopyRequest",
      "example": {
        "email": "Admin",
        "password": "Admin"
      },
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "required": [
        "email",
        "password"
      ]
    },
    "registerCopyRequest": {
      "title": "registerCopyRequest",
      "example": {
        "name": "Jauh",
        "email": "jauhin@mail.com",
        "password": "pajajauh321"
      },
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "required": [
        "name",
        "email",
        "password"
      ]
    },
    "post_category_request1": {
      "title": "post_category_request1",
      "example": {
        "category": "Desktop"
      },
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        }
      },
      "required": [
        "category"
      ]
    },
    "post_roleCopyRequest": {
      "title": "post_roleCopyRequest",
      "example": {
        "rolename": "admin",
        "description": "Admin Bagian"
      },
      "type": "object",
      "properties": {
        "rolename": {
          "type": "string"
        },
        "description": {
          "type": "string"
        }
      },
      "required": [
        "rolename",
        "description"
      ]
    }
  },
  "security": [],
  "tags": [
    {
      "name": "Auth"
    },
    {
      "name": "User"
    },
    {
      "name": "Course"
    },
    {
      "name": "Roles"
    },
    {
      "name": "Type Course"
    },
    {
      "name": "Course Category"
    },
    {
      "name": "Materials"
    },
    {
      "name": "auth"
    },
    {
      "name": "user"
    },
    {
      "name": "course"
    },
    {
      "name": "role"
    },
    {
      "name": "type_course"
    },
    {
      "name": "course_category"
    },
    {
      "name": "material"
    },
    {
      "name": "lokal"
    }
  ]
}
`