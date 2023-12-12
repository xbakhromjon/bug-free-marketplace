package common

type PageableResult[T any] struct {
	Content    []T
	TotalCount int
}

type PageableRequest struct {
	Page int
	Size int
	SortRequest
}

type SortRequest struct {
	Sort      string
	Direction string
}
