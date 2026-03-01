package todo

import "errors"

// データ作成時のバリデーション
func ValidateInsertTodoRequest(req InsertTodoRequest) error {
	if req.TodoTitle == "" {
		return errors.New("todo_title is required")
	}

	if len(req.TodoTitle) > 64 {
		return errors.New("todo_title must be less than 255 characters")
	}

	return nil
}

// 更新時のバリデーション
func ValidateUpdateTodoRequest(req UpdateTodoRequest) error {
	if req.TodoTitle == "" {
		return errors.New("todo_title is required")
	}
	if len(req.TodoTitle) > 64 {
		return errors.New("todo_title must be 255 characters or less")
	}
	return nil
}
