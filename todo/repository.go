package todo

import (
	"database/sql"
	"errors"
)

// 全てのデータを取得する関数
func GetAll(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT todo_id, todo_title, todo_done, created_at FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close() //複数行データを取得しているので、db.Query関数はどこで処理を終了していいかわからないからこれで明示して閉じる
	var todos []Todo
	var todo Todo
	for rows.Next() { //変数内に入っているデータを一行ずつ読み込み、データがあればそのデータで後続処理を開始して、次のループではポインタを一個進める
		if err := rows.Scan(&todo.TodoId, &todo.TodoTitle, &todo.TodoDone, &todo.CreatedAt); err != nil { //現在の行を読み込んでいる
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

// 指定されたIDに紐づくデータを取得する関数
func GetById(db *sql.DB, id int) (*Todo, error) {
	row := db.QueryRow("SELECT todo_id, todo_title, todo_done, created_at FROM todo WHERE todo_id = $1", id)
	var todo Todo
	if err := row.Scan(&todo.TodoId, &todo.TodoTitle, &todo.TodoDone, &todo.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) { //errors.Isは該当するエラーが入ってないかを確認して、その場合何を返すかを決めてる
			return nil, nil //この場合０件取得の場合はエラーじゃなくてnilしか返さんようにしてる
		}
		return nil, err //それ以外のエラーはエラー内容をそのまま返す
	}
	return &todo, nil //該当するデータがある場合はTodo構造体のポインタを返す
}

// データを挿入して、挿入したデータのIDを返す関数
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
	result, err := db.Exec("UPDATE todo SET todo_title = $1 WHERE todo_id = $2", title, id)
	if err != nil {
		return err
	}
	rowsaffected, err := result.RowsAffected() //何行更新されたか
	if err != nil {
		return err
	}
	if rowsaffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// todoを完了未完了に更新する関数
func UpdateDone(db *sql.DB, id int) error {
	result, err := db.Exec(
		"UPDATE todo SET todo_done = NOT todo_done WHERE todo_id = $1", id) //todo_doneの値を反転させることで、完了状態と未完了状態を切り替える
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected() //何行更新されたか確認
	if err != nil {
		return err
	}

	if rowsAffected == 0 { //更新された行がない場合は、該当IDが存在しないことを示すエラーを返す
		return sql.ErrNoRows
	}

	return nil
}

// データを削除する関数
func DeleteData(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM todo WHERE todo_id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected() //何行削除されたか確認
	if err != nil {
		return err
	}

	// 削除された行がない = 該当IDが存在しない
	if rowsAffected == 0 {
		return sql.ErrNoRows //該当IDが存在しない場合はそのことを示すエラーを返す
	}

	return nil
}
