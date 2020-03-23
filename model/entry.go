package model

// Entry ...
type Entry struct {
	ID     int    `json:"id" form:"id"`
	Date   string `json:"date" form:"date"`
	Mood   int    `json:"mood" form:"mood"`
	Sleep  int    `json:"sleep" form:"sleep"`
	Stress int    `json:"stress" form:"stress"`
}
