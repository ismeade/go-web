//
// bolt.go
// liyang <ismeade99@sina.com>
// 2021-1-7
//
package dao

import (
	"fmt"
	"go-web/util"
	"os"

	"github.com/boltdb/bolt"
)

var (
	Db *bolt.DB
)

func init() {
	//创建bolt数据库本地文件
	home, _ := util.Home()
	os.MkdirAll(home+"/.go-web", os.ModePerm)
	dbc, err := bolt.Open(home+"/.go-web/data.db", 0600, nil)
	if err != nil {
		fmt.Println("open err:", err)
		return
	} else {
		Db = dbc
	}
}

func InitBucket(bucket []byte) {
	Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			_, err := tx.CreateBucket(bucket)
			return err
		}
		return nil
	})
}

func Insert(bucket []byte, key, value string) error {

	k := []byte(key)
	v := []byte(value)
	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		err := b.Put(k, v)
		return err
	})
	return err
}

func Rm(bucket []byte, key string) {
	k := []byte(key)
	Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete(k)
		return err
	})
}

func Read(bucket []byte, key string) string {
	k := []byte(key)
	var val []byte
	Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		val = b.Get(k)
		return nil
	})
	return string(val)
}

func FetchAll(bucket []byte) {
	Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			fmt.Printf("key is %s,value is %s\n", k, v)
		}
		return nil
	})
}
