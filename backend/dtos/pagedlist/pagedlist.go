package pagedlist

import (
	"net/url"
	"strconv"
)

// 获取分页列表Dto
type GetDto struct {
	// 序号
	Index int
	// 大小
	Size int
}

// 分页列表Dto
type Dto[T any] struct {
	// 总数
	Total int
	// 列表
	Items []T
}

func ParseGetDto(values url.Values) GetDto {
	dto := GetDto{}
	size, err := strconv.ParseInt(values.Get("size"), 10, 32)
	if err == nil {
		dto.Size = int(size)
	}
	index, err := strconv.ParseInt(values.Get("index"), 10, 32)
	if err == nil {
		dto.Index = int(index)
	}
	return dto
}
func Empty[T any]() Dto[T] {
	return Dto[T]{
		Total: 0,
		Items: make([]T, 0),
	}
}
