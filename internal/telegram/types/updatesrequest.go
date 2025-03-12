package types

type UpdatesRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
