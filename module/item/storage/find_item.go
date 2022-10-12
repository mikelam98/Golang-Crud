package todostorage

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(ctx context.Context, condition map[string]interface{}) (data *todomodel.ToDoItem, err error) {
	var itemdata todomodel.ToDoItem

	if err := s.db.Where(condition).First(&itemdata).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &itemdata, err
}
