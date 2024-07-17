package entities

import "time"

type Entity struct {
	ID       int64      `db:"id"`
	CreateAt *time.Time `db:"create_at"`
}
