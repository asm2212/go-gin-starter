package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"` // hashed
	CreatedAt time.Time `db:"created_at"`
}
