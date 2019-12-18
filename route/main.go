package route

import (
	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()
	r.GET("/", profile)
	r.Run(":8999")
}
