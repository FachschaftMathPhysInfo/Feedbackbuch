package model

import "time"

// Comment is the comment type for our API
type Comment struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Content    string    `json:"content"`
	Upvotes    int       `json:"upvotes"`
	Timestamp  time.Time `json:"timestamp"`
	References *int      `json:"references"`
}
