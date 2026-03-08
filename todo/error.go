package todo

import (
	"encoding/json"
	"net/http"
)

// エラーレスポンスの構造体
type ErrorResponse struct {
	Error string `json:"error"` //JSONに変換した時のキー名を"error"に指定している
}

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json") //ヘッダにContent-Typeをapplication/jsonに設定してる。これでクライアントはレスポンスがJSON形式だと認識できる
	w.WriteHeader(status)                              //HTTPステータスコードを設定してる(404とか)これでクライアントはレスポンスがエラーだと認識できる

	json.NewEncoder(w).Encode(ErrorResponse{
		Error: message, //引数で受け取ったmessageをErrorResponse構造体のErrorフィールドにセットしている
	}) //ErrorResponse構造体をエンコード関数でJSONエンコードしているので、キー名が"error"で、値がmessageに格納された値のJSONがレスポンスボディに書き込まれることになる
}
