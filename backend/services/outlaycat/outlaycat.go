package outlaycat

import (
	"bookkeeper-backend/dtos/outlaycat"
	"bookkeeper-backend/dtos/pagedlist"
	outlaycat_model "bookkeeper-backend/models/outlaycat"
	"errors"
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
func List(db *gorm.DB, dto outlaycat.ListDto) ([]outlaycat.Dto, int8, error) {
	var res []outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{})
	tx = FilterList(tx, dto).Find(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表失败。%s\n", NAME, tx.Error)
		return make([]outlaycat.Dto, 0), 1, tx.Error
	}
	return res, 0, nil
}
func Count(db *gorm.DB, dto outlaycat.ListDto) (int, int8, error) {
	var res int64 = 0
	tx := db.Model(&outlaycat_model.OutlayCat{})
	tx = FilterList(tx, dto).Count(&res)
	if tx.Error != nil {
		log.Printf("获取%s列表数量失败。%s\n", NAME, tx.Error)
		return 0, 1, tx.Error
	}
	return int(res), 0, nil
}
func PagedList(db *gorm.DB, dto outlaycat.PagedListDto) (pagedlist.Dto[outlaycat.Dto], int8, error) {
	res := pagedlist.Empty[outlaycat.Dto]()
	total, code, err := Count(db, dto.ListDto)
	if code != 0 {
		return res, 1, err
	}
	if total == 0 {
		return res, 0, nil
	}
	tx := db.Model(&outlaycat_model.OutlayCat{})
	tx = FilterList(tx, dto.ListDto).Offset(dto.Index * dto.Size).Limit(dto.Size).Find(&res.Items)
	if tx.Error != nil {
		log.Printf("获取%s分页列表失败。%s\n", NAME, tx.Error)
		return res, 2, tx.Error
	}
	res.Total = total
	return res, 0, nil
}
func Get(db *gorm.DB, id uuid.UUID) (*outlaycat.Dto, int8, error) {
	var res outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{Id: id}).First(&res)
	if tx.Error != nil {
		log.Printf("获取%s失败。%s\n", NAME, tx.Error)
		return nil, 1, tx.Error
	}
	return &res, 0, nil
}
func Create(db *gorm.DB, dto outlaycat.CreateDto) (*uuid.UUID, int8, error) {
	var dumplicate *outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{
		Base: outlaycat_model.Base{
			Name: dto.Name,
		},
	}).First(&dumplicate)
	if dumplicate != nil && dumplicate.Id != uuid.Nil && dumplicate.Name == dto.Name {
		log.Printf("创建%s失败。存在重复名称。\n", NAME)
		return nil, 1, errors.New("存在重复名称")
	} else if tx.Error != nil {
		log.Printf("创建%s失败。查重失败。%s\n", NAME, tx.Error)
		return nil, 2, tx.Error
	}
	model := outlaycat_model.OutlayCat{
		Base: dto,
		Id:   uuid.New(),
	}
	tx = db.Create(&model)
	if tx.Error != nil {
		log.Printf("创建%s失败。%s\n", NAME, tx.Error)
		return nil, 3, tx.Error
	}
	return &model.Id, 0, nil
}
func Update(db *gorm.DB, dto outlaycat.UpdateDto) (bool, int8, error) {
	var dumplicate *outlaycat.Dto
	tx := db.Model(&outlaycat_model.OutlayCat{
		Base: outlaycat_model.Base{
			Name: dto.Name,
		},
	}).Not("Id = ?", dto.Id).First(&dumplicate)
	if dumplicate != nil && dumplicate.Id != uuid.Nil && dumplicate.Name == dto.Name {
		log.Printf("更新%s失败。存在重复名称。\n", NAME)
		return false, 1, errors.New("存在重复名称")
	} else if tx.Error != nil {
		log.Printf("更新%s失败。查重失败。%s\n", NAME, tx.Error)
		return false, 2, tx.Error
	}
	model := outlaycat_model.OutlayCat{
		Base: dto.Base,
		Id:   dto.Id,
	}
	tx = db.Omit("Id").Updates(&model)
	if tx.Error != nil {
		log.Printf("更新%s失败。%s\n", NAME, tx.Error)
		return false, 3, tx.Error
	}
	return true, 0, nil
}
func Delete(db *gorm.DB, id uuid.UUID) (bool, int8, error) {
	tx := db.Delete(&outlaycat_model.OutlayCat{Id: id})
	if tx.Error != nil {
		log.Printf("删除%s失败。%s\n", NAME, tx.Error)
		return false, 1, tx.Error
	}
	return true, 0, nil
}
