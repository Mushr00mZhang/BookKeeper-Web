package outlay

import (
	"bookkeeper-backend/dtos/outlay"
	"bookkeeper-backend/dtos/pagedlist"
	outlay_model "bookkeeper-backend/models/outlay"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TABLE = "outlays"
const NAME = "支出"

func List(db *gorm.DB, dto outlay.ListDto) []outlay.Dto {
	var res []outlay.Dto
	tx := db.Model(&outlay_model.Outlay{}).Where("Parent = ?", dto.ParentId)
	if dto.CatId != nil {
		tx = tx.Where("CatId = ?", dto.CatId)
	}
	if dto.LowMoney != nil {
		tx.Where("Money >= ?", dto.LowMoney)
	}
	if dto.TopMoney != nil {
		tx.Where("Money <= ?", dto.TopMoney)
	}
	if dto.STime != nil {
		tx.Where("Date >= ?", dto.STime)
	}
	if dto.ETime != nil {
		tx.Where("Date <= ?", dto.ETime)
	}
	tx = tx.Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlay.Dto, 0)
	}
	return res
}
func PagedList(db *gorm.DB, dto outlay.PagedListDto) pagedlist.ResDto[outlay.Dto] {
	list := List(db, dto.ListDto)
	total := len(list)
	pagedList := make([]outlay.Dto, 0)
	sIndex := dto.Dto.Index * dto.Dto.Size
	eIndex := sIndex + dto.Dto.Size
	if sIndex < total && eIndex <= total {
		pagedList = list[sIndex:eIndex]
	} else if eIndex > total {
		pagedList = list[sIndex:total]
	}
	return pagedlist.ResDto[outlay.Dto]{
		Total: total,
		Items: pagedList,
	}
}
func Get(db *gorm.DB, id uuid.UUID) *outlay.Dto {
	var res outlay.Dto
	tx := db.Model(&outlay_model.Outlay{
		Id: id,
	}).First(&res)
	if tx.Error != nil {
		log.Printf("获取%s失败。%s\n", NAME, tx.Error)
		return nil
	}
	return &res
}
func Create(db *gorm.DB, dto outlay.CreateDto) *uuid.UUID {
	model := outlay_model.Outlay{
		Base: dto,
		Id:   uuid.New(),
	}
	tx := db.Create(&model)
	if tx.Error != nil {
		log.Printf("创建%s失败。%s\n", NAME, tx.Error)
		return nil
	}
	return &model.Id
}
