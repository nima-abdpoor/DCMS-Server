package api

import (
	db "DCMS/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", homePage)
	router.POST("/upload/single", server.uploadSingleFile)
	router.StaticFS("/images", http.Dir("public"))
	router.GET("/config/:id", server.getConfig)
	router.GET("/config", server.getDefaultConfig)
	router.POST("/config", server.postConfig)
	router.POST("/sendLogFile", server.postLogFile)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
