package transport

import (
	todobiz "Golang-Crud/module/item/business"
	todomodel "Golang-Crud/module/item/model"
	todostorage "Golang-Crud/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func HandleListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging todomodel.DataPaging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewGetAllTodoItemStorage(storage)

		result, err := biz.ListToDoItems(c.Request.Context(), nil, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
