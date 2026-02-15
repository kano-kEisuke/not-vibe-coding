package todo

import "time"

type Todo struct {
	TodoId    int
	TodoTitle string
	TodoDone  bool
	CreatedAt time.Time
}
