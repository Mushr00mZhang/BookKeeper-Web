package outlaycat

import (
	"bookkeeper-backend/dtos/outlaycat"
	"bookkeeper-backend/dtos/pagedlist"
	outlaycat_model "bookkeeper-backend/models/outlaycat"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TABLE = "outlay_cats"
const NAME = "支出类型"

func FilterList(tx *gorm.DB, dto outlaycat.ListDto) *gorm.DB {
	if dto.ParentId != nil && *dto.ParentId != uuid.Nil {
		tx = tx.Where("Parent = ?", dto.ParentId)
	}
	return tx
}
func List(db *gorm.DB, dto outlaycat.ListDto) []outlaycat.Dto {
	var res []outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{})
	tx = FilterList(tx, dto).Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlaycat.Dto, 0)
	}
	return res
}
func Count(db *gorm.DB, dto outlaycat.ListDto) int {
	var res int64 = 0
	tx := db.Model(&outlaycat_model.OutlayCat{})
	tx = FilterList(tx, dto).Count(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表数量失败。%s\n", NAME, tx.Error)
	}
	return int(res)
}
func PagedList(db *gorm.DB, dto outlaycat.PagedListDto) pagedlist.Dto[outlaycat.Dto] {
	total := Count(db, dto.ListDto)
	items := make([]outlaycat.Dto, 0)
	if total == 0 {
		return pagedlist.Dto[outlaycat.Dto]{
			Total: total,
			Items: items,
		}
	}
	tx := db.Model(&outlaycat_model.OutlayCat{})
	if dto.ParentId == nil || *dto.ParentId == uuid.Nil {
		tx = tx.Where("Parent = ?", dto.ParentId)
	}
	tx = tx.Offset(dto.Index * dto.Size).Limit(dto.Size).Find(&items)
	if tx.Error != nil {
		log.Printf("获取%s分页列表失败。%s\n", NAME, tx.Error)
	}
	return pagedlist.Dto[outlaycat.Dto]{
		Total: total,
		Items: items,
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
func Update(db *gorm.DB, dto outlaycat.UpdateDto) bool {
	var dumplicate *outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{
		Base: outlaycat_model.Base{
			Name: dto.Name,
		},
	}).Not("Id = ?", dto.Id).First(&dumplicate)
	if dumplicate != nil && dumplicate.Id != uuid.Nil && dumplicate.Name == dto.Name {
		log.Printf("更新%s失败。存在重复名称。\n", NAME)
		return false
	} else if tx.Error != nil {
		log.Printf("更新%s失败。查重失败。%s\n", NAME, tx.Error)
		return false
	}
	model := outlaycat_model.OutlayCat{
		Base: dto.Base,
		Id:   dto.Id,
	}
	tx = db.Omit("Id").Updates(&model)
	if tx.Error != nil {
		log.Printf("更新%s失败。%s\n", NAME, tx.Error)
		return false
	}
	return true
}
func Delete(db *gorm.DB, id uuid.UUID) bool {
	tx := db.Delete(&outlaycat_model.OutlayCat{Id: id})
	if tx.Error != nil {
		log.Printf("删除%s失败。%s\n", NAME, tx.Error)
		return false
	}
	return true
}
