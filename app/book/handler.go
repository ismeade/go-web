//
// handler.go
// liyang <ismeade99@sina.com>
// 2021-1-7
//
package book

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("init")
	//dao.InitBucket()
}

func create(c *gin.Context) {
}

func get(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getAll(c *gin.Context) {
}

func update(c *gin.Context) {
}

func del(c *gin.Context) {
}
