package models

import "time"

type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Priority  int64     `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}

// NewNote is
func NewNote(title string, content string, priority int64) *Note {
	return &Note{Title: title, Content: content, Priority: priority, CreatedAt: time.Now()}
}
