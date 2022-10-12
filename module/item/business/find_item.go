package todobiz

import (
	todomodel "Golang-Crud/module/item/model"
	"context"
)

type FindToDoItem interface {
	FindItem(ctx context.Context,
		condition map[string]interface{}) (*todomodel.ToDoItem, error)
}
type findBiz struct {
	store FindToDoItem
}

func NewFindItemBiz(store FindToDoItem) *findBiz {
	return &findBiz{store: store}
}
func (biz *findBiz) FindItems(ctx context.Context, condition map[string]interface{}) (*todomodel.ToDoItem, error) {
	itemData, err := biz.store.FindItem(ctx, condition)
	if err != nil {
		return nil, err
	}
	return itemData, nil
}
