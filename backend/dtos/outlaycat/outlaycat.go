package outlaycat

import (
	"bookkeeper-backend/dtos/pagedlist"
	"bookkeeper-backend/models/outlay"
	"bookkeeper-backend/models/outlaycat"
	"net/url"

	"github.com/google/uuid"
)

// 支出类型Dto
type Dto struct {
	outlaycat.OutlayCat
	Outlays *[]outlay.Outlay `gorm:"foreignKey:CatId;references:Id" json:"outlays"` // 支出列表
	// Parent *outlaycat.OutlayCat `gorm:"foreignKey:Id;references:ParentId"` // 父级
}

// 创建支出Dto
type CreateDto = outlaycat.Base

// 更新支出Dto
type UpdateDto struct {
	Id uuid.UUID `gorm:"not null;primaryKey;" json:"id"` // Id
	outlaycat.Base
}

// 获取支出列表Dto
type ListDto struct {
	ParentId *uuid.UUID `json:"parentId"` // 父级Id
}

// 解析获取支出列表Dto
func ParseListDto(values url.Values) ListDto {
	dto := ListDto{}
	parentId, err := uuid.Parse(values.Get("parentId"))
	if err == nil {
		dto.ParentId = &parentId
	}
	return dto
}

// 获取支出分页列表Dto
type PagedListDto struct {
	pagedlist.GetDto
	ListDto
}

// 解析获取支出分页列表Dto
func ParsePagedListDto(values url.Values) PagedListDto {
	return PagedListDto{
		GetDto:  pagedlist.ParseGetDto(values),
		ListDto: ParseListDto(values),
	}
}
