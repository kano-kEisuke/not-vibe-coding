package todo

import "errors"

// ValidateInsertTodoRequest はInsertTodoRequestのバリデーション
func ValidateInsertTodoRequest(req InsertTodoRequest) error {
	if req.TodoTitle == "" {
		return errors.New("todo_title is required")
	}

	if len(req.TodoTitle) > 255 {
		return errors.New("todo_title must be less than 255 characters")
	}

	return nil
}
