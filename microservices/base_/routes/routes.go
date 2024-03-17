package routes

import (
	"github.com/gin-gonic/gin"
)

// AppRouting slice contains all the routes of the application
var AppRouting = []Group{}

// Group .
type Group struct {
	Prefix      string
	Middlewares []gin.HandlerFunc
	Routes      []Route
}

// Route .
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
