package transport

import (
	todobiz "Golang-Crud/module/item/business"
	todostorage "Golang-Crud/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func HandleDeleteItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewDeleteToDoItemBiz(storage)
		if err := biz.DeleteItems(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
