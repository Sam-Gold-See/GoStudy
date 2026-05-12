package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Name string
	Age  int
}

func main() {
	username := "root"
	pwd := "123456"
	addr := "localhost:3306"
	database := "local_test"
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, pwd, addr, database)
	// 连接数据库
	db, err := gorm.Open("mysql", args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	// 自动创建表
	db.AutoMigrate(&User{})
	// 插入一行记录
	user := User{Name: "xiaoming", Age: 18}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	// 查询记录
	var tempUser User
	err = db.Where("name = ?", "xiaoming").First(&tempUser).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tempUser)
}
