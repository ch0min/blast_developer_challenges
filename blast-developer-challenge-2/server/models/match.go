package models

type Match struct {
	Rounds      []Round            `json:"rounds"`
	Players     map[string]*Player `json:"players"`
	TotalRounds int                `json:"total_rounds"`
}