package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine) {
	c.POST("/users/signup")
	c.POST("/users/login")
	c.POST("/admin/product")
}
