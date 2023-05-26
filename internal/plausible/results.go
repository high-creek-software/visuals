package plausible

type Results[T any] struct {
    Results T `json:"results"`
}