package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
	}
}
