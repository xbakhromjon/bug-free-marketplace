package common

type PageableResult[T any] struct {
	content    []T
	totalCount int
}

func CreatePageableResult[T any](content []T, totalCount int) *PageableResult[T] {

	return &PageableResult[T]{content: content, totalCount: totalCount}
}

type PageableRequest struct {
	page uint64
	size uint64
}

func (p *PageableRequest) GetPage() (uint64, bool) {
	if p.page == 0 {
		return 0, false
	}
	return p.page, true
}

func (p *PageableRequest) GetSize() (uint64, bool) {
	if p.size == 0 {
		return 0, false
	}
	return p.size, true
}

type SortRequest struct {
	sort      string
	direction string
}

func CreateDefaultPageableRequest() *PageableRequest {
	return CreatePageableRequest(1, 10)
}

func CreatePageableRequest(page uint64, size uint64) *PageableRequest {

	return &PageableRequest{
		page: page,
		size: size,
	}
}
