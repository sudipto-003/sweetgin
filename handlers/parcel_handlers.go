package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudipto-003/sweet-gin/models"
	"github.com/sudipto-003/sweet-gin/repository"
)

func NewParcelHandler(store *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}
