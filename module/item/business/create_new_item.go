package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
	"errors"
)


type CreateToDoItemStorage struct {
	CreateItem(ctx context.Context,data **todomodel.ToDoItem) error
}
type createBiz struct {
	store CreateToDoItemStorage
}
func NewCreateToDoItemBiz(ctx context.Context) *createBiz{
	return &createBiz{store: store}
}
func (biz *createBiz)CreateNewItem(ctx context.Context,data *todomodel.ToDoItem)error{
	if data.Title == ""{
		return errors.New("title cannot be empty")
	}
	data.Status = "Doing"
	if err := biz.store.CreateItem(ctx,data);err != nil {
		return err
	}
}
