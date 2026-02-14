package main

import (
	"database/sql"
	"log"

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
}
