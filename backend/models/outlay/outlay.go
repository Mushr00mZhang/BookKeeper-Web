package outlay

import (
	"bookkeeper-backend/models/outlaycat"
	"time"

	"github.com/google/uuid"
)

// 支出基础属性
type Base struct {
	Name     string    `gorm:"not null;size:64;" json:"name"`      // 名称
	CatId    uuid.UUID `gorm:"not null;" json:"catId"`             // 类型Id
	Money    float32   `gorm:"not null;precision:2;" json:"money"` // 金额
	Original *float32  `gorm:"precision:2;" json:"original"`       // 原价
	Amount   *float32  `gorm:"precision:2;" json:"amount"`         // 数量
	Unit     *string   `gorm:"" json:"unit"`                       // 单位
	Time     time.Time `gorm:"not null;" json:"time"`              // 时间
	UserId   uuid.UUID `gorm:"not null;" json:"userId"`            // 用户Id
}

// 支出
type Outlay struct {
	Id uuid.UUID `gorm:"not null;primaryKey;" json:"id"` // Id
	Base
	Cat outlaycat.OutlayCat `json:"cat"` // 类型
}
