package web

import "github.com/gin-gonic/gin"

var server *gin.Engine

func SetEngine(srv *gin.Engine) {
	server = srv
}
