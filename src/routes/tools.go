package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge/src/controllers"
)

func RegisterToolRoutes(g *gin.Engine) {
	t := g.Group("/api/v1/tools")
	{
		c := controllers.NewToolsController()
		t.GET("/", c.GetAllTools)
	}
}
