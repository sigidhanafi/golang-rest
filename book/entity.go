package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
