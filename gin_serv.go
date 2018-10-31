package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	req := gin.Default()
	v1 := req.Group("/mailing/v1")

	templ := v1.Group("/template")

	templ.POST("/:template_id", func(c *gin.Context) { postTemplate(c) })
	templ.GET("/:template_id_name", func(c *gin.Context) { getTemplate(c) })

	send := v1.Group("/send")

	send.POST("/specific/:template_id", func(c *gin.Context) { sendSpecific(c) })
	send.POST("/broadcast/:template_id", func(c *gin.Context) { sendBroadcast(c) })

	req.Run(":22000")

}

func postTemplate(c *gin.Context) {
	bindeo := &Template{}
	err := c.Bind(&bindeo)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%s", string(bindeo.template_name))
	c.JSON(http.StatusAccepted, gin.H{"mensaje": string(bindeo.template_id)})
	return
}
func getTemplate(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"mensaje": (c.Param("template_id_name"))})
}
func sendSpecific(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"mensaje": (c.Param("template_id"))})
}
func sendBroadcast(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"mensaje": (c.Param("template_id"))})
}
