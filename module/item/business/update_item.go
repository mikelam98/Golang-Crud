package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

type UpdateToDoItemStorage interface {
	FindItem(ctx context.Context,
		condition map[string]interface{}) (*todomodel.ToDoItem, error)
	UpdateItem(ctx context.Context,
		condition map[string]interface{},
		dataUpdate *todomodel.ToDoItem) error
}
type updatebiz struct {
	store UpdateToDoItemStorage
}

func NewUpdateToDoItemStorage(store UpdateToDoItemStorage) *updatebiz {
	return &updatebiz{store: store}
}
func (biz *updatebiz) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *todomodel.ToDoItem) error {
	oldItem, err := biz.store.FindItem(ctx, condition)
	if err != nil {
		return err
	}
	// just a demo in case we dont allow update Finished item
	if oldItem.Status == "Finished" {
		return err
	}
	if err := biz.store.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}
	return nil
}
