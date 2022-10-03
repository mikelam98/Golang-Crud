package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

type ListItemStorage interface {
	ListToDoItem(ctx context.Context,
		condition map[string]interface{},
		paging *todomodel.DataPaging,
	) ([]todomodel.ToDoItem, error)
}
type listAllbiz struct {
	store ListItemStorage
}

func NewGetAllTodoItemStorage(store ListItemStorage) *listAllbiz {
	return &listAllbiz{store: store}
}
func (biz *listAllbiz) ListToDoItems(ctx context.Context, condition map[string]interface{}, paging *todomodel.DataPaging) ([]todomodel.ToDoItem, error) {
	result, err := biz.store.ListToDoItem(ctx, condition, paging)
	if err != nil {
		return nil, err
	}
	return result, err
}
