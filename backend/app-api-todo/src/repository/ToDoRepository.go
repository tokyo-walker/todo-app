package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"../data"
)

func Add(title string, text string, db *gorm.DB) () {
	todo := data.Todo{}

	todo.ID = 0
	// テーブル定義でAUTO_INCREMENTを指定すれば０でInsertしたときに次の数字が入る

	todo.Title = title
	todo.Text = text
	todo.Favorite = true

	db.Create(&todo) // insert
}

func Update(id int,  db *gorm.DB) {
	todoById := data.Todo{}

	todoById.ID = id
	todo := todoById

	// id でUserとってこれる
	db.First(&todo)
	todo.Favorite = !todoById.Favorite

	db.Save(&todo)
}

func UpdateText(id int, text string, db *gorm.DB) {
	todo := data.Todo{}
	db.Model(todo).Where("id = ?", id).Update("text", text)
}

func GetAll(db *gorm.DB) {
	fmt.Println("get all todos")

	var todos []data.Todo
	db.Find(&todos)

	fmt.Println(todos)
}

func DeleteById(id int, db *gorm.DB) {
	todo := data.Todo{}
	todo.ID = id
	// まずは削除したいレコードの情報を埋める
	db.First(&todo)
	db.Delete(&todo)
}
