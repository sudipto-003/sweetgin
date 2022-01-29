package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
		if err = repo.GetAllParcels(context.Background(), &parcels); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, parcels)
	}
}

func GetParcelByPID(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parcel models.Parcel
		pid := c.Param("pid")

		if err := repo.GetParcelByPID(context.Background(), pid, &parcel); err != nil {
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

func GetParcelsByDate(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		const timeformat = "02-01-2006"
		defaultdate := time.Now().Format(timeformat)
		datestr := c.DefaultQuery("date", defaultdate)
		cdate, err := time.Parse(timeformat, datestr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date string can not be parsed"})
			return
		}
		var parcels []models.Parcel

		if err = repo.GetParcelByDate(context.Background(), cdate, &parcels); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, parcels)
	}
}

func GenerateOTPForParcelUpdate(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		var parcel models.Parcel
		pid := c.Param("pid")

		if err := repo.GetParcelByPID(context.Background(), pid, &parcel); err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusOK, gin.H{"message": "No Such Document Exist"})
				return
			} else {
				c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
				return
			}
		}
		otp := NewOTP()
		err := repo.SetOTP(context.Background(), parcel.ParcelId, otp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "OTP created successfully"})
	}
}

type UserOTP struct {
	OTP string `json:"otp" binding:"required"`
}

func VerifyAndUpdateParcel(repo *repository.Repos) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user_otp UserOTP
		// var parcel models.Parcel
		pid := c.Param("pid")
		if err := c.ShouldBindJSON(&user_otp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		otp, err := repo.GetOTP(context.Background(), pid)
		if err == redis.Nil {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Invalid pid or OTP expired"})
			return
		}

		if otp != user_otp.OTP {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP does not match"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OTP Verfied Succesfully"})
	}
}
