package routes

import (
	"DCMS/api/controllers"
	"DCMS/db/postgresql/sqlc"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup, server *db.Store) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler(server))
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup, server *db.Store) {
	g.GET("/dashboard", controllers.DashboardGetHandler(server))
	g.GET("/logout", controllers.LogoutGetHandler())
}
