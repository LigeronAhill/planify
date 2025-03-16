package types

type Story struct {
	ID   int   `json:"id"`
	Chat *Chat `json:"chat"`
}
