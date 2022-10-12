package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

type DeleteToDoItemStorage interface {
	FindItem(ctx context.Context,
		condition map[string]interface{}) (*todomodel.ToDoItem, error)
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}
type deleteBiz struct {
	store DeleteToDoItemStorage
}

func NewDeleteToDoItemBiz(store DeleteToDoItemStorage) *deleteBiz {
	return &deleteBiz{store: store}
}
func (biz *deleteBiz) DeleteItems(ctx context.Context, condition map[string]interface{}) error {
	_, err := biz.store.FindItem(ctx, condition)
	if err != nil {
		return err
	}
	if err := biz.store.DeleteItem(ctx, condition); err != nil {
		return err
	}
	return nil
}
