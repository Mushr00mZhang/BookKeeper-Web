package outlaycat

import (
	"github.com/google/uuid"
)

type Base struct {
	ParentId *uuid.UUID `gorm:""`                  // 父级Id
	Name     string     `gorm:"not null;size:64;"` // 名称
	Unit     string     `gorm:"not null;size:16;"` // 单位
	Sort     uint8      `gorm:"not null;"`         // 排序
	Stable   bool       `gorm:"not null;"`         // 是否固定
	Remark   *string    `gorm:"size:256;"`         // 备注
}

// 支出类型
type OutlayCat struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	Base
	Children []OutlayCat `gorm:"foreignKey:Id;references:ParentId"` // 子级列表
	// Outlays []outlay.Outlay `gorm:"foreignKey:Id;references:CatId"` // 支出列表
}
