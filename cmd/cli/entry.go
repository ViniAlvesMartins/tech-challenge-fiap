package cli

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

// "fmt"

// "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/adapter/handler"

func Execute() {

	// 	fmt.Println("entry")

	// 	messageHandler := h.NewMessageHandler()

	// 	messageHandler.PostMessage()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})

	router.Run(":8080")
}
