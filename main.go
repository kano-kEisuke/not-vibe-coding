package main

import (
	"database/sql"
	"fmt"
	"log"
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

	//CRUD動作チェック
	log.Println("CRUD動作チェック")

	//挿入
	id1, err := todo.InsertData(db, "インサートできた")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("挿入されたデータのID", id1)

	//1件取得
	todotest, err := todo.GetById(db, id1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("取得したデータ: ID=%d, Title=%s\n", todotest.TodoId, todotest.TodoTitle)

	//更新
	err = todo.UpdateData(db, id1, "更新できた")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("更新完了")

	//2件目挿入
	var id2 int
	id2, err = todo.InsertData(db, "インサート2件目")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("挿入されたデータのID:", id2)

	//全件取得
	todos, err := todo.GetAll(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("全件取得結果")
	for _, t := range todos {
		fmt.Printf("  ID=%d, Title=%s, Done=%v\n", t.TodoId, t.TodoTitle, t.TodoDone)
	}

	//削除
	err = todo.DeleteData(db, id1)
	if err != nil {
		log.Fatal(err)
	}
	err = todo.DeleteData(db, id2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("削除完了")

	//削除後にデータの残留確認
	todosAfter, err := todo.GetAll(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✓ 削除後のTodoの件数: %d\n", len(todosAfter))
}
