//
// user.go
// liyang <ismeade99@sina.com>
// 2021-1-7
//
package models

import (
	"encoding/json"
	"go-web/dao"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const bucketUser = "user"

func init() {
	dao.InitBucket([]byte(bucketUser))
}

func CreateUser(user *User) {
	b, err := json.Marshal(user)
	if err == nil {
		dao.Insert([]byte(bucketUser), string(user.Id), string(b))
	}
}
