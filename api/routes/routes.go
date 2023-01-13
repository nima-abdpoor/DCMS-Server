package routes

import (
	"DCMS/api/controllers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
}
