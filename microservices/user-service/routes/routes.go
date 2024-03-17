package routes

import (
	"github.com/gin-gonic/gin"
)

// AppRouting slice que contiene todas las rutas de la aplicacion
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
