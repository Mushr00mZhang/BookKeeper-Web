package outlaycat

import (
	"github.com/google/uuid"
)

type Base struct {
	ParentId *uuid.UUID `gorm:"" json:"parentId"`              // 父级Id
	Name     string     `gorm:"not null;size:64;" json:"name"` // 名称
	Unit     string     `gorm:"not null;size:16;" json:"unit"` // 单位
	Sort     uint8      `gorm:"not null;" json:"sort"`         // 排序
	Stable   bool       `gorm:"not null;" json:"stable"`       // 是否固定
	Remark   *string    `gorm:"size:256;" json:"remark"`       // 备注
}

// 支出类型
type OutlayCat struct {
	Id uuid.UUID `gorm:"not null;primaryKey;" json:"id"` // Id
	Base
	Children []OutlayCat `gorm:"foreignKey:Id;references:ParentId" json:"children"` // 子级列表
	// Outlays []outlay.Outlay `gorm:"foreignKey:Id;references:CatId"` // 支出列表
}
