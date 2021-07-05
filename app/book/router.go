//
// router.go
// liyang <ismeade99@sina.com>
// 2021-1-7
//
package book

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/book", getAll)
	e.GET("/book/:id", get)
	e.POST("/book", create)
	e.PUT("/book/:id", update)
	e.DELETE("/book/:id", del)
}
