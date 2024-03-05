package outlaycat

import (
	"bookkeeper-backend/dtos/pagedlist"
	"bookkeeper-backend/models/outlay"
	"bookkeeper-backend/models/outlaycat"

	"github.com/google/uuid"
)

// 支出类型
type Dto struct {
	outlaycat.OutlayCat
	Outlays []outlay.Outlay // 支出列表
	// Parent *outlaycat.OutlayCat `gorm:"foreignKey:Id;references:ParentId"` // 父级
}

// 创建支出Dto
type CreateDto = outlaycat.Base

// 更新支出Dto
type UpdateDto struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	outlaycat.Base
}

// 获取支出列表Dto
type ListDto struct {
	ParentId *uuid.UUID // 父级Id
}

func ParseListDto(vars map[string]string) ListDto {
	dto := ListDto{}
	parentId, err := uuid.Parse(vars["ParentId"])
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

func ParsePagedListDto(vars map[string]string) PagedListDto {
	return PagedListDto{
		GetDto:  pagedlist.ParseGetDto(vars),
		ListDto: ParseListDto(vars),
	}
}
