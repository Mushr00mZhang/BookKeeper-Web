package main

import (
	"time"

	"github.com/google/uuid"
)

// 支出基础属性
type OutlayBase struct {
	Name      string    `gorm:"not null;size:64;"`     // 名称
	CatId     uuid.UUID `gorm:"not null;"`             // 类型Id
	Price     float32   `gorm:"not null;precision:2;"` // 价格
	Original  *float32  `gorm:"precision:2;"`          // 原价
	UnitPrice *float32  `gorm:"precision:2;"`          // 单价
	Amount    *float32  `gorm:"precision:2;"`          // 数量
	Time      time.Time `gorm:"not null;"`             // 时间
}

// 支出
type Outlay struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	OutlayBase
	// Name      string    `gorm:"not null;size:64;"`              // 名称
	// CatId     uuid.UUID `gorm:"not null;"`                      // 类型Id
	// Price     float32   `gorm:"not null;precision:2;"`          // 价格
	// Original  *float32  `gorm:"precision:2;"`                   // 原价
	// UnitPrice *float32  `gorm:"precision:2;"`                   // 单价
	// Amount    *float32  `gorm:"precision:2;"`                   // 数量
	// Time      time.Time `gorm:"not null;"`                      // 时间
	Cat OutlayCat `gorm:"foreignKey:Id;references:CatId"` // 类型
}

// 创建支出Dto
type CreateOutlayDto = OutlayBase

// 更新支出Dto
type UpdateOutlayDto struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	OutlayBase
}

// 获取支出列表Dto
type ListOutlayDto struct {
}

// 获取支出列表分页Dto
type PagedListOutlayDto struct {
	ListOutlayDto
}

// 支出类型
type OutlayCat struct {
	Id       uuid.UUID   `gorm:"not null;primaryKey;"`              // Id
	ParentId *uuid.UUID  `gorm:""`                                  // 父级Id
	Name     string      `gorm:"not null;size:64;"`                 // 名称
	Unit     string      `gorm:"not null;size:16;"`                 // 单位
	Priority uint8       `gorm:"not null;"`                         // 优先级
	Stable   bool        `gorm:"not null;"`                         // 是否固定
	Remark   *string     `gorm:"not null;size:256;"`                // 备注
	Parent   *OutlayCat  `gorm:"foreignKey:Id;references:ParentId"` // 父级
	Children []OutlayCat `gorm:"foreignKey:ParentId;references:Id"` // 子级列表
	OutLays  []Outlay    `gorm:"foreignKey:CatId;references:Id"`    // 支出列表
}

// 收入
type Income struct {
	Id    uuid.UUID `gorm:"not null;primaryKey;"`  // Id
	Name  string    `gorm:"not null;size:64;"`     // 名称
	CatId uuid.UUID `gorm:"not null;"`             // 类型
	Price float32   `gorm:"not null;precision:2;"` // 价格
	Time  time.Time `gorm:"not null;"`             // 时间
	Cat   IncomeCat // 类型
}

// 收入类型
type IncomeCat struct {
	Id       uuid.UUID   `gorm:"not null;primaryKey;"`              // Id
	ParentId *uuid.UUID  `gorm:""`                                  // 父级Id
	Name     string      `gorm:"not null;size:64;"`                 // 名称
	Stable   bool        `gorm:"not null;"`                         // 是否固定
	Remark   string      `gorm:"not null;size:256;"`                // 备注
	Parent   *IncomeCat  `gorm:"foreignKey:Id;references:ParentId"` // 父级
	Children []IncomeCat `gorm:"foreignKey:ParentId;references:Id"` // 子级列表
	Incomes  []Income    `gorm:"foreignKey:CatId;references:Id"`    // 收入列表
}
