package pagedlist

import "strconv"

// 分页Dto
type Dto struct {
	// 序号
	Index int
	// 大小
	Size int
}

// 分页结果Dto
type ResDto[T any] struct {
	Total int
	Items []T
}

func ParseDto(vars map[string]string) Dto {
	dto := Dto{}
	size, err := strconv.ParseInt(vars["Size"], 10, 32)
	if err == nil {
		dto.Size = int(size)
	}
	index, err := strconv.ParseInt(vars["Index"], 10, 32)
	if err == nil {
		dto.Index = int(index)
	}
	return dto
}
