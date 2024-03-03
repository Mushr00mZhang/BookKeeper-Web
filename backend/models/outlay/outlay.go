package outlay

import (
	"bookkeeper-backend/models/outlaycat"
	"time"

	"github.com/google/uuid"
)

// 支出基础属性
type Base struct {
	Name     string    `gorm:"not null;size:64;"`     // 名称
	CatId    uuid.UUID `gorm:"not null;"`             // 类型Id
	Money    float32   `gorm:"not null;precision:2;"` // 金额
	Original *float32  `gorm:"precision:2;"`          // 原价
	Amount   *float32  `gorm:"precision:2;"`          // 数量
	Unit     *string   `gorm:""`                      // 单位
	Time     time.Time `gorm:"not null;"`             // 时间
	UserId   uuid.UUID `gorm:"not null;"`             // 用户Id
}

// 支出
type Outlay struct {
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
	Base
	Cat outlaycat.OutlayCat // 类型
}
