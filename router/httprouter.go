package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sudipto-003/sweet-gin/handlers"
	"github.com/sudipto-003/sweet-gin/models"
)

func GetHttpRouter(store *handlers.ParcelCollection) *gin.Engine {
	router := gin.Default()

	router.POST("/neworder", func(c *gin.Context) {
		var parcel models.Parcel
		if err := c.ShouldBindJSON(&parcel); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if err := store.CreateNewParcelOreder(context.Background(), &parcel); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, parcel)
	})

	return router
}
