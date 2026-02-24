package api

import (
	"net/http"
	"not-vibe-coding/todo"
)

func main() {
	//http.HandleFunc 「どの」URLで「なに」をするか
	http.HandleFunc("/health", todo.Health)

	//サーバー立てて8080ポートで待つ　nilはデフォルトルーター使うって意味らしい
	http.ListenAndServe(":8080", nil)
}
