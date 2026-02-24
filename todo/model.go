package todo

import "time"

type Todo struct {
	TodoId    int       `json:"todo_id"`
	TodoTitle string    `json:"todo_title"`
	TodoDone  bool      `json:"todo_done"`
	CreatedAt time.Time `json:"created_at"`
}
