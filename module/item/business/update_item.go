package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

type UpdateToDoItemStorage interface {
	FindItem(ctx context.Context,
		condition map[string]interface{},
		data *todomodel.ToDoItem)
	UpdateItem(ctx context.Context,
		condition map[string]interface{},
		dataUpdate *todomodel.ToDoItem)
}
type updatebiz struct {
	store UpdateToDoItemStorage
}

func NewUpdateToDoItemStorage(store UpdateToDoItemStorage) *updatebiz {
	return &updatebiz{store: store}
}
func (biz *updatebiz) UpdateItems(ctx context.Context, condition map[string]interface{}, dataUpdate *todomodel.ToDoItem) error {
	oldItem, err := biz.store.FindItem(ctx, condition)
}
