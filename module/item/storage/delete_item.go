package todostorage

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

func (s *mysqlStorage) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	if err := s.db.Table(todomodel.ToDoItem{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

//.Table(todomodel.ToDoItem{}.TableName())
