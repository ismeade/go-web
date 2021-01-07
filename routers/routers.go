//
// routers.go
// liyang <ismeade99@sina.com>
// 2021-1-6
//
package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/html", "./static")
	for _, opt := range options {
		opt(r)
	}
	return r
}
