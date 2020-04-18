package entry

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

const createSQL = `
INSERT INTO entries (uuid, date, mood, sleep, stress, notes, updated_at, created_at)
	VALUES (:uuid, :date, :mood, :sleep, :stress, :notes, :updated_at, :created_at)
`

func (es *Service) CreateEntry(entry *Model) (*Model, error) {
	entry.UUID = uuid.New()
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = entry.CreatedAt

	_, err := es.DB.NamedExec(createSQL, entry)
	if err != nil {
		return nil, fmt.Errorf("unable to create entry: %w", err)
	}

	return es.getEntryByUUID(entry.UUID)
}

func (es *Service) ListEntries(limit int64) (*[]Model, error) {
	entries := &[]Model{}
	err := es.DB.Select(entries, `SELECT * FROM entries ORDER BY date desc LIMIT $1`, limit)
	if err != nil {
		return nil, fmt.Errorf("problem getting entries: %w", err)
	}
	return entries, nil
}

func (es *Service) GetEntry(id int64) (*Model, error) {
	entry := &Model{}
	err := es.DB.Get(entry, "SELECT * FROM entries WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("unable to find entry(%d): %w", id, err)
	}
	return entry, nil
}

func (es *Service) getEntryByUUID(uuid uuid.UUID) (*Model, error) {
	entry := &Model{}
	err := es.DB.Get(entry, "SELECT * FROM entries WHERE uuid=$1", uuid)
	if err != nil {
		return nil, fmt.Errorf("unable to find entry(%d): %w", uuid, err)
	}
	return entry, nil
}
