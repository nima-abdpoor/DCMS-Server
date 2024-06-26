package api

import (
	"DCMS/api/middleware"
	"DCMS/api/routes"
	"DCMS/db/postgresql/sqlc"
	"DCMS/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Store  *db.Store
	Router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{Store: store}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(sessions.Sessions("session", cookie.NewStore(util.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public, server.Store)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private, server.Store)

	//Router.GET("/", homePage)
	router.POST("/dashboard/upload/single", server.uploadSingleFile)
	router.StaticFS("/images", http.Dir("public"))
	router.GET("/config/:id", server.getConfig)
	router.GET("/config", server.getDefaultConfig)
	router.POST("/config", server.postConfig)
	router.POST("/sendLogFile/:id", server.postLogFile)
	router.POST("/sendLog/:id", server.postLog)
	router.POST("/customer", server.postCustomer)
	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
