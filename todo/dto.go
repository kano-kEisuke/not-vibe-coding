package todo

// Todo作成リクエストのリクエストボディ
type InsertTodoRequest struct {
	TodoTitle string `json:"todo_title"`
}

// Todo作成のレスポンスボディ
type InsertTodoResponse struct {
	TodoId int `json:"todo_id"`
}

// Todo更新リクエスト
type UpdateTodoRequest struct {
	TodoTitle string `json:"todo_title"`
}

// Todo更新レスポンス
type UpdateTodoResponse struct {
	RowsAffected int64 `json:"rows_affected"` //何行更新されたか
}
