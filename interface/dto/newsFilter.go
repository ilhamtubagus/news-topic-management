package dto

type NewsFilter struct {
	Status string `query:"s"`
	Topic  uint64 `query:"t"`
}
