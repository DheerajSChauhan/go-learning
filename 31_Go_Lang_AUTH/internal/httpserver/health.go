package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"ok":	 true,
		"status": "healthy",
		"service": "go-auth",
		"time":time.Now().UTC(),
	})

}