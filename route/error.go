package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/sworker/system"
)

func method404(c *gin.Context) {
	c.JSON(404, gin.H{"code": system.Method404Code, "message": system.Method404Msg})
}
