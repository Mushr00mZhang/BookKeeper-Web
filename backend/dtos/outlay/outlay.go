package outlay

import (
	"strconv"
	"time"

	"bookkeeper-backend/dtos/pagedlist"
	"bookkeeper-backend/models/outlay"
	"bookkeeper-backend/utils"

	"github.com/google/uuid"
)

// 支出Dto
type Dto = outlay.Outlay

// 创建支出Dto
type CreateDto = outlay.Base

// 更新支出Dto
type UpdateDto struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	outlay.Base
}

// 获取支出列表Dto
type ListDto struct {
	ParentId *uuid.UUID // 父级Id
	CatId    *uuid.UUID // 类型Id
	LowMoney *float32   // 最低金额
	TopMoney *float32   // 最高金额
	STime    *time.Time // 开始时间
	ETime    *time.Time // 结束时间
}

// 解析获取支出列表Dto
func ParseListDto(vars map[string]string) ListDto {
	dto := ListDto{}
	parentId, err := uuid.Parse(vars["ParentId"])
	if err == nil {
		dto.ParentId = &parentId
	}
	catId, err := uuid.Parse(vars["CatId"])
	if err == nil {
		dto.CatId = &catId
	}
	lowMoney, err := strconv.ParseFloat(vars["LowMoney"], 32)
	if err == nil {
		temp := float32(lowMoney)
		dto.LowMoney = &temp
	}
	topMoney, err := strconv.ParseFloat(vars["TopMoney"], 32)
	if err == nil {
		temp := float32(topMoney)
		dto.TopMoney = &temp
	}
	sTime, err := utils.ParseTime(vars["STime"])
	if err == nil {
		dto.STime = sTime
	}
	eTime, err := utils.ParseTime(vars["ETime"])
	if err == nil {
		dto.ETime = eTime
	}
	return dto
}

// 获取支出分页列表Dto
type PagedListDto struct {
	pagedlist.Dto
	ListDto
}

// 解析获取支出分页列表Dto
func ParsePagedListDto(vars map[string]string) PagedListDto {
	return PagedListDto{
		Dto:     pagedlist.ParseDto(vars),
		ListDto: ParseListDto(vars),
	}
}
