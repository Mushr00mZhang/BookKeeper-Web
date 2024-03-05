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

func FilterList(tx *gorm.DB, dto outlay.ListDto) *gorm.DB {
	// if dto.ParentId != nil && *dto.ParentId != uuid.Nil {
	// 	tx = tx.Where("Parent = ?", dto.ParentId)
	// }
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
	if dto.UserId != nil {
		tx = tx.Where("UserId = ?", dto.UserId)
	}
	return tx
}
func List(db *gorm.DB, dto outlay.ListDto) []outlay.Dto {
	var res []outlay.Dto
	tx := db.Model(&outlay_model.Outlay{})
	tx = FilterList(tx, dto).Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlay.Dto, 0)
	}
	return res
}
func Count(db *gorm.DB, dto outlay.ListDto) int {
	var res int64 = 0
	tx := db.Model(&outlay_model.Outlay{})
	tx = FilterList(tx, dto).Count(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表数量失败。%s\n", NAME, tx.Error)
	}
	return int(res)
}
func PagedList(db *gorm.DB, dto outlay.PagedListDto) pagedlist.Dto[outlay.Dto] {
	total := Count(db, dto.ListDto)
	items := make([]outlay.Dto, 0)
	if total == 0 {
		return pagedlist.Dto[outlay.Dto]{
			Total: total,
			Items: items,
		}
	}
	tx := db.Model(&outlay_model.Outlay{})
	tx = FilterList(tx, dto.ListDto).Offset(dto.Index * dto.Size).Limit(dto.Size).Find(&items)
	if tx.Error != nil {
		log.Printf("获取%s分页列表失败。%s\n", NAME, tx.Error)
	}
	return pagedlist.Dto[outlay.Dto]{
		Total: total,
		Items: items,
	}
}
func Get(db *gorm.DB, id uuid.UUID) *outlay.Dto {
	var res outlay.Dto
	tx := db.Model(&outlay_model.Outlay{Id: id}).First(&res)
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
func Update(db *gorm.DB, dto outlay.UpdateDto) bool {
	model := outlay_model.Outlay{
		Base: dto.Base,
		Id:   dto.Id,
	}
	tx := db.Omit("Id").Updates(&model)
	if tx.Error != nil {
		log.Printf("更新%s失败。%s\n", NAME, tx.Error)
		return false
	}
	return true
}
func Delete(db *gorm.DB, id uuid.UUID) bool {
	tx := db.Delete(&outlay_model.Outlay{Id: id})
	if tx.Error != nil {
		log.Printf("删除%s失败。%s\n", NAME, tx.Error)
		return false
	}
	return true
}
