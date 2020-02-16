package main

import (
	"demo/gxormDemo/tables"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	engine *xorm.Engine
)

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:lwstar@tcp(127.0.0.1:3306)/lw?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = engine.Close()
	}()

	if err = engine.Sync(new(tables.User)); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		tables.Insert(engine, "liwei", 30, 30, "陕西西安")
	}

	u := tables.Query(engine, 19)
	if u != nil {
		log.Println(u)
	}

	u1 := tables.Query2(engine, 19)
	if u1 != nil {
		log.Println(u1)
	}

	us := tables.Find(engine)
	log.Println(us)

	us1 := tables.Find2(engine)
	log.Println(us1)

	log.Println("ok")
}
