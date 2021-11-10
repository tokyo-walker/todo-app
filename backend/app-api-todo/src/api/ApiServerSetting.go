package api

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"

	"../data"
	"../mysql"
	"../repository"
)

type Setting struct {}
const SUCCESS_MSG = "{status : 200}"

func  (setting Setting) Connect() ()  {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/defer", DeferTest),

		rest.Get("/get", GetAllTodos), //ok
		rest.Post("/add/:title/:text", AddTodo), //ok
		rest.Put("/update/:id", UpdateTodo), //?
		rest.Put("/modify/:id/:text", UpdateTodoText), //ok
		rest.Delete("/delete/:id", DeleteTodo), //ok
		rest.Get("/api/v1/user", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson("{hogehoge: 0}")
		}),
		&rest.Route{"Get", "/hoge", func(writer rest.ResponseWriter, request *rest.Request) {

			writer.WriteJson("{hogehoge: 0}")
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func AddTodo(w rest.ResponseWriter, r *rest.Request) () {
	var setting mysql.Setting

	db := setting.Connect()
	todo := data.Todo{}
	title := r.PathParam("title")
	text := r.PathParam("text")
	repository.Add(title, text, db)

	todo.ID = 0
	// テーブル定義でAUTO_INCREMENTを指定すれば０でInsertしたときに次の数字が入る

	todo.Title = title
	todo.Text = text
	todo.Favorite = true

	db.Create(&todo) // insert

	_ = w.WriteJson(SUCCESS_MSG)
}

func UpdateTodo(w rest.ResponseWriter, r *rest.Request) {
	var setting mysql.Setting
	db := setting.Connect()
	todoById := data.Todo{}

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
	var setting mysql.Setting
	db := setting.Connect()
	todo := data.Todo{}

	id := r.PathParam("id")
	text := r.PathParam("text")

	db.Model(todo).Where("id = ?", id).Update("text", text)

	_ = w.WriteJson(SUCCESS_MSG)
}

func GetAllTodos(w rest.ResponseWriter, r *rest.Request) {
	var setting mysql.Setting
	db := setting.Connect()
	fmt.Println("get all todos")

	var todos []data.Todo
	db.Find(&todos)

	fmt.Println(todos)

	db.Close()

	_ = w.WriteJson(&todos)
}

func DeleteTodo(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")

	var setting mysql.Setting
	db := setting.Connect()
	todo := data.Todo{}

	todo.ID, _ = strconv.Atoi(id)
	// まずは削除したいレコードの情報を埋める
	db.First(&todo)
	db.Delete(&todo)
	db.Close()

	_ = w.WriteJson(SUCCESS_MSG)
}

func DeferTest(w rest.ResponseWriter, r *rest.Request) {
	fmt.Println("aaaaaa....?")

	defer fmt.Println("This msg printed after hoge")

	fmt.Println("hoge")

	_ = w.WriteJson(SUCCESS_MSG)
}