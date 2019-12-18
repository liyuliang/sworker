package route

import "github.com/gin-gonic/gin"

func method404(c *gin.Context) {
	c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Not found"})
}
