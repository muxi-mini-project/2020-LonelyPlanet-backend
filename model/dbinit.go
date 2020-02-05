package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/howeyc/gopass"
	"github.com/jinzhu/gorm"
	"log"
)

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase(user, password string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", user+":"+password+"@tcp(localhost:3306)/mini_project")
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
	}
	db.SingularTable(true)
	return db, err
}

func (db *Database) Init() {
	var user string
	fmt.Println("输入数据库用户名:")
	fmt.Scanf("%s", &user)
	fmt.Println("输入密码:")
	password, _ := gopass.GetPasswdMasked()
	newDb, err := getDatabase(user, string(password))
	if err != nil {

	}
	Db = &Database{Self: newDb}
}

func (db *Database) Close() error {
	if err := Db.Self.Close(); err != nil {
		return err
	}
	return nil
}
