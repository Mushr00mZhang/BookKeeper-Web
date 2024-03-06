package outlay

import (
	"net/url"
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
	CatId    *uuid.UUID // 类型Id
	LowMoney *float32   // 最低金额
	TopMoney *float32   // 最高金额
	STime    *time.Time // 开始时间
	ETime    *time.Time // 结束时间
	UserId   *uuid.UUID // 用户Id
}

// 解析获取支出列表Dto
func ParseListDto(values url.Values) ListDto {
	dto := ListDto{}
	catId, err := uuid.Parse(values.Get("CatId"))
	if err == nil {
		dto.CatId = &catId
	}
	lowMoney, err := strconv.ParseFloat(values.Get("LowMoney"), 32)
	if err == nil {
		temp := float32(lowMoney)
		dto.LowMoney = &temp
	}
	topMoney, err := strconv.ParseFloat(values.Get("TopMoney"), 32)
	if err == nil {
		temp := float32(topMoney)
		dto.TopMoney = &temp
	}
	sTime, err := utils.ParseTime(values.Get("STime"))
	if err == nil {
		dto.STime = sTime
	}
	eTime, err := utils.ParseTime(values.Get("ETime"))
	if err == nil {
		dto.ETime = eTime
	}
	userId, err := uuid.Parse(values.Get("UserId"))
	if err == nil {
		dto.UserId = &userId
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
