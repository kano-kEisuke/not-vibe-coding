package todo

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) { // 引数で「どこに」「何を」書き込むかを指定　メッセンジャー的な
	// w: クライアントへの返信先（メッセンジャーの配達先）
	// r: クライアントからのリクエスト情報（何を要求しているか）

	//HTTP/1.1 200 OK みたいなのがヘッダに書き込まれて返る
	//ブラウザやクライアントが成功したと認識できる
	w.WriteHeader(http.StatusOK)

	//レスポンスボディにokを書き込んでる
	// []byte()で文字列をバイナリに変換（HTTPはバイナリでやり取りするため必要）
	w.Write([]byte("ok"))
}

func GetAllTodos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := GetAll(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}
}
