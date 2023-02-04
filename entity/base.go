package entity

import "time"

type Base struct {
	CreatedBy int64
	CreatedAt time.Time
	ID        int64
}

func (b Base) IsEmpty() bool {
	return b.CreatedAt.IsZero() && b.CreatedBy == 0 && b.ID == 0
}
