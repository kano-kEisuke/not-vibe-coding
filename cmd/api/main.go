package main

import (
	"database/sql"
	"log"
	"net/http"
	"not-vibe-coding/todo"

	_ "github.com/lib/pq"
)

func main() {

	dsn := "user=keisukekano dbname=postgres sslmode=disable" //DB設定
	db, err := sql.Open("postgres", dsn)                      //DB接続
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil { //接続確認
		log.Panic(err)
	}
	log.Println("DB接続成功")
	defer db.Close()

	var dbName string
	row := db.QueryRow("SELECT current_database()") //現在のデータベースを取得する
	if err := row.Scan(&dbName); err != nil {       //クエリの結果を変数に格納
		log.Panic(err)
	}
	log.Println("現在のデータベース:", dbName)

	// http.HandleFunc 「どの」URLで「なに」をするか
	http.HandleFunc("/health", todo.Health)

	// GET/todos/all 全てのTodoリストを取得する
	http.HandleFunc("/todos/all", todo.GetAllTodos(db))

	// GET/todo?id={id} 指定されたIDのTodoを取得する
	http.HandleFunc("/todo", todo.GetTodo(db))

	// POST/todos/create Todoリストを新規作成する
	http.HandleFunc("/todos/create", todo.CreateTodo(db))

	// サーバー立てて8080ポートで待つ　nilはデフォルトルーター使うって意味らしい
	http.ListenAndServe(":8080", nil)

}
