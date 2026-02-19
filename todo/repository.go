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

// データを挿入して挿入したデータのIDを返す関数
func InsertData(db *sql.DB, title string) (int, error) {
	var id int
	row := db.QueryRow("INSERT INTO todo (todo_title,todo_done) VALUES ($1, false) RETURNING todo_id", title)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil //小さいデータはポインタじゃなくてコピー渡すで大丈夫
}

// データを更新する関数
func UpdateData(db *sql.DB, id int, title string) error {
	_, err := db.Exec("UPDATE todo SET todo_title = $1 WHERE todo_id = $2", title, id) //使わない値を受け取る時に_を使う
	return err
}

// データを削除する関数
func DeleteData(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todo WHERE todo_id = $1", id)
	return err //ifでnilかエラーが入ったかで条件分けなくても戻り値error型やったらこの書き方でいいらしい
} //エラーあったらエラー返すしなかったらnil返してくれる
