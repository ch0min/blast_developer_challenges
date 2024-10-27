package models

type Round struct {
	RoundNumber int     `json:"round_number"`
	StartTime   int64   `json:"start_time"`
	EndTime     int64   `json:"end_time"`
	Duration    float64 `json:"duration"`
}