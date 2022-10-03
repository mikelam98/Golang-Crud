package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
	"errors"
)

type CreateToDoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
}
type createBiz struct {
	store CreateToDoItemStorage
}

func NewCreateToDoItemBiz(store CreateToDoItemStorage) *createBiz {
	return &createBiz{store: store}
}
func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if data.Title == "" {
		return errors.New("title cannot be empty")
	}
	if data.Status == "" {
		return errors.New("status cannot be empty")
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
