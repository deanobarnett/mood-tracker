package entry

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        int        `json:"id" db:"id"`
	UUID      uuid.UUID  `json:"uuid" db:"uuid"`
	Date      string     `json:"date" db:"date"`
	Mood      int        `json:"mood" db:"mood"`
	Sleep     int        `json:"sleep" db:"sleep"`
	Stress    int        `json:"stress" db:"stress"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
