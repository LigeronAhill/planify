package types

type Response[T any] struct {
	Ok     bool `json:"ok"`
	Result T    `json:"result"`
}
