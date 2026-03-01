package todo

// Todo作成リクエストのリクエストボディ
type InsertTodoRequest struct {
	TodoTitle string `json:"todo_title"`
}

// Todo作成のレスポンスボディ
type InsertTodoResponse struct {
	TodoId int `json:"todo_id"`
}
