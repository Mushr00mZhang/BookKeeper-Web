package book

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 支出基础属性
type OutlayBase struct {
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
	OutlayBase
	Id  uuid.UUID `gorm:"not null;primaryKey;"`           // Id
	Cat OutlayCat `gorm:"foreignKey:Id;references:CatId"` // 类型
}

// 创建支出Dto
type CreateOutlayDto = OutlayBase

// 更新支出Dto
type UpdateOutlayDto struct {
	OutlayBase
	Id uuid.UUID `gorm:"not null;primaryKey;"` // Id
}

// 获取支出列表Dto
type ListOutlayDto struct {
	ParentId *uuid.UUID // 父级Id
	CatId    *uuid.UUID // 类型Id
	LowMoney *float32   // 最低金额
	TopMoney *float32   // 最高金额
	STime    *time.Time // 开始时间
	ETime    *time.Time // 结束时间
}

// 获取支出列表分页Dto
type PagedListOutlayDto struct {
	ListOutlayDto
}

// 账单类型
type BookCat struct {
	Id       uuid.UUID  `gorm:"not null;primaryKey;"`              // Id
	ParentId *uuid.UUID `gorm:""`                                  // 父级Id
	Name     string     `gorm:"not null;size:64;"`                 // 名称
	Unit     string     `gorm:"not null;size:16;"`                 // 单位
	Sort     uint8      `gorm:"not null;"`                         // 排序
	Stable   bool       `gorm:"not null;"`                         // 是否固定
	Remark   *string    `gorm:"size:256;"`                         // 备注
	Parent   *BookCat   `gorm:"foreignKey:Id;references:ParentId"` // 父级
	Children []BookCat  `gorm:"foreignKey:ParentId;references:Id"` // 子级列表
}

// 支出类型
type OutlayCat struct {
	BookCat
	OutLays []Outlay `gorm:"foreignKey:CatId;references:Id"` // 支出列表
}

func ListOutlayCat(db *gorm.DB) []OutlayCat {
	var res []OutlayCat
	result := db.Find(&res)
	fmt.Println(result.Error)
	return res
}
func GetOutlayCat(db *gorm.DB, id uuid.UUID) *OutlayCat {
	var res *OutlayCat = nil
	db.Model(&OutlayCat{
		BookCat: BookCat{Id: id},
	}).First(&res)
	return res
}
func CreateOutlayCat(db *gorm.DB) *OutlayCat {
	var dumplicate *OutlayCat
	db.Model(&OutlayCat{
		BookCat: BookCat{Name: "其他"},
	}).First(&dumplicate)
	if dumplicate != nil && dumplicate.Id != uuid.Nil && dumplicate.Name == "其他" {
		return nil
	}
	res := OutlayCat{
		BookCat: BookCat{
			Id:       uuid.New(),
			ParentId: nil,
			Name:     "其他",
			Unit:     "个",
			Sort:     99,
			Stable:   false,
			Remark:   nil,
		},
	}
	result := db.Create(&res)
	fmt.Println(result.Error)
	return &res
}

// 收入
type Income struct {
	Id     uuid.UUID `gorm:"not null;primaryKey;"`           // Id
	Name   string    `gorm:"not null;size:64;"`              // 名称
	CatId  uuid.UUID `gorm:"not null;"`                      // 类型
	Money  float32   `gorm:"not null;precision:2;"`          // 金额
	Time   time.Time `gorm:"not null;"`                      // 时间
	UserId uuid.UUID `gorm:"not null;"`                      // 用户Id
	Cat    IncomeCat `gorm:"foreignKey:Id;references:CatId"` // 类型
}

// 收入类型
type IncomeCat struct {
	BookCat
	Incomes []Income `gorm:"foreignKey:CatId;references:Id"` // 收入列表
}
