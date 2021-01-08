package serve

import (
	"fmt"
	"github.com/Jansora/pancake/backend/serve/routers"
	"github.com/Jansora/pancake/backend/tools"
	"github.com/gin-gonic/gin"
)

func App() {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routers.Init(r)

	port := fmt.Sprintf(":%d", tools.Port)
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
