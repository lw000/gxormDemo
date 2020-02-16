package tables

import (
	"github.com/go-xorm/xorm"
	"log"
)

type User struct {
	Id      int32  `xorm:"pk autoincr"`
	Name    string `xorm:"name"`
	Age     int32  `xorm:"age"`
	Sex     int32  `xorm:"sex"`
	Address string `xorm:"address"`
}

func Insert(engine *xorm.Engine, name string, age, sex int32, address string) (int64, bool) {
	user := new(User)
	user.Name = name
	user.Age = age
	user.Sex = sex
	user.Address = address
	affected, err := engine.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}

func Del(engine *xorm.Engine, id int64) {
	user := new(User)
	_, _ = engine.Id(id).Delete(user)
}

func Query(engine *xorm.Engine, id int32) *User {
	user := &User{Id: id}
	is, _ := engine.Get(user)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	return user
}

func Query2(engine *xorm.Engine, id int32) *User {
	user := &User{Id: id}
	is, _ := engine.Where("id=?", id).Get(user)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	return user
}

func Find(engine *xorm.Engine) map[int32]User {
	users := make(map[int32]User)
	err := engine.Find(&users)
	if err != nil {
		log.Fatal("搜索结果不存在!")
	}
	return users
}

func Find2(engine *xorm.Engine) []User {
	users := make([]User, 0)
	err := engine.Find(&users)
	if err != nil {
		log.Fatal("搜索结果不存在!")
	}
	return users
}

func Update(engine *xorm.Engine, id int64, user *User) bool {
	affected, err := engine.ID(id).Update(user)
	if err != nil {
		log.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}

func Transactions(engine *xorm.Engine) {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	user := &User{Name: "1111", Age: 0, Sex: 0, Address: "1231312312"}
	var n int64
	n, err = session.Insert(user)
	if err != nil {
		log.Println(err)
		return
	}
	if n > 0 {

	}

	err = session.Rollback()
	if err != nil {
		log.Println(err)
		return
	}

	err = session.Commit()
	if err != nil {
		log.Println(err)
		return
	}
}
