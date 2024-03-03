package outlaycat

import (
	"bookkeeper-backend/dtos/outlaycat"
	"bookkeeper-backend/dtos/pagedlist"
	outlaycat_model "bookkeeper-backend/models/outlaycat"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TABLE = "outlaycats"
const NAME = "支出类型"

func List(db *gorm.DB, dto outlaycat.ListDto) []outlaycat.Dto {
	var res []outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{}).Where("Parent = ?", dto.ParentId).Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlaycat.Dto, 0)
	}
	return res
}
func PagedList(db *gorm.DB, dto outlaycat.PagedListDto) pagedlist.ResDto[outlaycat.Dto] {
	list := List(db, dto.ListDto)
	total := len(list)
	pagedList := make([]outlaycat.Dto, 0)
	sIndex := dto.Dto.Index * dto.Dto.Size
	eIndex := sIndex + dto.Dto.Size
	if sIndex < total && eIndex <= total {
		pagedList = list[sIndex:eIndex]
	} else if eIndex > total {
		pagedList = list[sIndex:total]
	}
	return pagedlist.ResDto[outlaycat.Dto]{
		Total: total,
		Items: pagedList,
	}
}
func Get(db *gorm.DB, id uuid.UUID) *outlaycat.Dto {
	var res outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{Id: id}).First(&res)
	if tx.Error != nil {
		log.Printf("获取%s失败。%s\n", NAME, tx.Error)
		return nil
	}
	return &res
}
func Create(db *gorm.DB, dto outlaycat.CreateDto) *uuid.UUID {
	var dumplicate *outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{
		Base: outlaycat_model.Base{
			Name: dto.Name,
		},
	}).First(&dumplicate)
	if dumplicate != nil && dumplicate.Id != uuid.Nil && dumplicate.Name == dto.Name {
		log.Printf("创建%s失败。存在重复名称。\n", NAME)
		return nil
	} else if tx.Error != nil {
		log.Printf("创建%s失败。查重失败。%s\n", NAME, tx.Error)
		return nil
	}
	model := outlaycat_model.OutlayCat{
		Base: dto,
		Id:   uuid.New(),
	}
	tx = db.Create(&model)
	if tx.Error != nil {
		log.Printf("创建%s失败。%s\n", NAME, tx.Error)
		return nil
	}
	return &model.Id
}
