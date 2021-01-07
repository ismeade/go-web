//
// main.go
// Copyright (C) 2021 liyang <ismeade99@sina.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"go-web/app/book"
	"go-web/routers"
)

func main() {
	routers.Include(book.Routers)

	r := routers.Init()

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("StartUp service failed, err:%v\n", err)
	}
}
