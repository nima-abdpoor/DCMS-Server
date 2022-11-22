package api

import "github.com/gin-gonic/gin"

type Server struct {
	//store *db.Store
	router *gin.Engine
}

func NewServer() *Server {

	server := &Server{}
	router := gin.Default()
	router.POST("/config", server.getConfig)
	server.router = router
	return server
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}