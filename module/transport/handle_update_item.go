package transport

import (
	todobiz "Golang-Crud/module/item/business"
	todomodel "Golang-Crud/module/item/model"
	todostorage "Golang-Crud/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func HandleUpdateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var dataUpdate todomodel.ToDoItem
		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewUpdateToDoItemStorage(storage)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
