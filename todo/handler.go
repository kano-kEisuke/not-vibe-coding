package todo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// 全件取得ハンドラ
func GetAllTodos(db *sql.DB) http.HandlerFunc { //返り値で関数を返している（クロージャっていうらしい）
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := GetAll(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error")) //httpはバイトで通信するからバイト型に変換
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos) //インスタンス化してモジュール呼んでる。EncodeはJSONに変換して、wって宛先に書き込んでる。
	} //上記の関数追記メモ：ストリーム処理はデータ全部をメモリに溜めずに少しずつ処理すること。
}

// todo作成ハンドラ
func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストで送られてきたJSONデータをGOの構造体にデコード
		var req InsertTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request body"))
			return
		}

		// バリデーション
		if err := ValidateInsertTodoRequest(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// データベースに挿入
		id, err := InsertData(db, req.TodoTitle)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
			return
		}

		// レスポンス返却
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(InsertTodoResponse{TodoId: id})
	}
}

// 一件取得ハンドラ
func GetTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID parameter")) //httpはバイトで通信するからバイト型に変換
			return
		}
		todo, err := GetById(db, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		if todo == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Todo Not Found"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			log.Printf("JSON encode error: %v", err)
		}
	}
}

// todo更新ハンドラ
func UpdateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// パスからID取得
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID parameter"})
			return
		}
		// リクエストボディのデコード
		var req UpdateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
			return
		}
		// 入力チェック
		if err := ValidateUpdateTodoRequest(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		rowsAffected, err := UpdateData(db, id, req.TodoTitle)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(UpdateTodoResponse{RowsAffected: rowsAffected})
	}
}

// 一件削除ハンドラ
func DeleteTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID parameter"})
			return
		}
		err = DeleteData(db, id)
		if err != nil {
			// 該当IDが存在しない場合
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
				return
			}
			// その他のエラー
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
