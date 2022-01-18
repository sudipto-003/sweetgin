package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudipto-003/sweet-gin/models"
	"github.com/sudipto-003/sweet-gin/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetParcelByIDHandler(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parcel models.Parcel
		var id primitive.ObjectID

		id, err := primitive.ObjectIDFromHex(c.Query("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := repo.GetParcelInfoById(context.Background(), id, &parcel); err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusOK, gin.H{"message": "No Such Document Exist"})
				return
			} else {
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
				return
			}
		}

		c.JSON(http.StatusOK, parcel)
	}
}

func GetAllParcelsHandler(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parcels []models.Parcel
		var err error
		if parcels, err = repo.GetAllParcels(context.Background()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, parcels)
	}
}
