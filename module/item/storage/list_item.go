package todostorage

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

func (s *mysqlStorage) ListToDoItem(ctx context.Context,
	condition map[string]interface{},
	paging *todomodel.DataPaging,
) ([]todomodel.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []todomodel.ToDoItem
	if err := s.db.Table(todomodel.ToDoItem{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
