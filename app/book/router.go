//
// router.go
// liyang <ismeade99@sina.com>
// 2021-1-7
//
package book

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/book", create)
	e.GET("/book/:id", get)
}
