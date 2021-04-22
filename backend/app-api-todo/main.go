package main
// package main じゃないとRunできない

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"strconv"

	// _操作はこのパッケージをインポートするだけでパッケージの中の関数を直接使うわけではなく、このパッケージの中にあるinit関数をコールします。
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"

)

const SUCCESS_MSG = "{status : 200}"

func main() {

	setApiServe()
}

func AddTodo(w rest.ResponseWriter, r *rest.Request) () {
	db := dbConn()
	todo := Todo{}
	title := r.PathParam("title")
	text := r.PathParam("text")

	todo.ID = 0
	// テーブル定義でAUTO_INCREMENTを指定すれば０でInsertしたときに次の数字が入る

	todo.Title = title
	todo.Text = text
	todo.Favorite = true

	db.Create(&todo) // insert

	_ = w.WriteJson(SUCCESS_MSG)
}

func UpdateTodo(w rest.ResponseWriter, r *rest.Request) {
	db := dbConn()
	todoById := Todo{}

	id := r.PathParam("id")
	todoById.ID, _ = strconv.Atoi(id)
	todo := todoById

	// id でUserとってこれる
	db.First(&todo)
	todo.Favorite = !todoById.Favorite

	db.Save(&todo)

	_ = w.WriteJson(SUCCESS_MSG)
}

func UpdateTodoText(w rest.ResponseWriter, r *rest.Request) {
	db := dbConn()
	todo := Todo{}

	id := r.PathParam("id")
	text := r.PathParam("text")

	db.Model(todo).Where("id = ?", id).Update("text", text)

	_ = w.WriteJson(SUCCESS_MSG)
}

func GetAllTodos(w rest.ResponseWriter, r *rest.Request) {
	//var setting mysql.MysqlSetting
	//db := setting.Connect()
	fmt.Println("get all todos")
	db := dbConn()

	var todos []Todo
	db.Find(&todos)

	// json
	//json, _ := json.Marshal(&todos)

	fmt.Println(todos)

	defer db.Close()

	_ = w.WriteJson(&todos)
}

func DeleteTodo(w rest.ResponseWriter, r *rest.Request) {
	db := dbConn()
	todo := Todo{}

	id := r.PathParam("id")
	todo.ID, _ = strconv.Atoi(id)
	// まずは削除したいレコードの情報を埋める
	db.First(&todo)
	db.Delete(&todo)

	defer db.Close()

	_ = w.WriteJson(SUCCESS_MSG)
}

func setApiServe() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/get", GetAllTodos),
		rest.Post("/add/:title/:text", AddTodo),
		rest.Put("/update/:id", UpdateTodo),
		rest.Put("/modify/:id/:text", UpdateTodoText),
		rest.Delete("/delete/:id", DeleteTodo),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func dbConn() (database *gorm.DB) {
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

type Todo struct {
	ID       int
	Title    string
	Text  string
	Favorite bool
}