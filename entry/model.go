package entry

import (
	"time"

	"github.com/google/uuid"
)

// Model holds the basic data structure for an entry
type Model struct {
	ID        int        `json:"id" db:"id"`
	UUID      uuid.UUID  `json:"uuid" db:"uuid"`
	Date      string     `json:"date" db:"date"`
	Mood      int        `json:"mood" db:"mood"`
	Sleep     int        `json:"sleep" db:"sleep"`
	Stress    int        `json:"stress" db:"stress"`
	Notes     *string    `json:"notes" db:"notes"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
