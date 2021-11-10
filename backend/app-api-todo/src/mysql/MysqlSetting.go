package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// _操作はこのパッケージをインポートするだけでパッケージの中の関数を直接使うわけではなく、このパッケージの中にあるinit関数をコールします。
	"github.com/jinzhu/gorm"
	"time"
)

type Setting struct {}

// 引数に自分を入れないと他のクラスで使えない。
// function名の二個目目が戻り値
func (setting Setting) Connect() (database *gorm.DB)  {
	fmt.Println("hello hello")
	DBMS := "mysql"
	USER := "mysql"
	PASS := "mysql"
	PROTOCOL := "tcp(127.0.0.1:3307)"
	DBNAME := "todo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
