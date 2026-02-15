package todo

import (
	"database/sql"
	"errors"
)

// 指定されたIDに紐づくデータを取得する関数
func GetById(db *sql.DB, id int) (*Todo, error) {
	row := db.QueryRow("SELECT todo_id, todo_title, todo_done, created_at FROM todo WHERE todo_id = $1", id)
	var todo Todo
	if err := row.Scan(&todo.TodoId, &todo.TodoTitle, &todo.TodoDone, &todo.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) { //該当するデータがない場合はnil返す
			return nil, nil
		}
		return nil, err //それ以外のエラーはエラー内容をそのまま返す
	}
	return &todo, nil //該当するデータがある場合はTodo構造体のポインタを返す
}
