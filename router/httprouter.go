package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sudipto-003/sweet-gin/handlers"
	"github.com/sudipto-003/sweet-gin/repository"
)

func GetHttpRouter(repo *repository.Repos) *gin.Engine {

	router := gin.Default()

	router.POST("/neworder", handlers.NewParcelHandler(repo))
	router.GET("/getorder", handlers.GetParcelByIDHandler(repo))
	router.GET("/allorder", handlers.GetAllParcelsHandler(repo))
	router.GET("/parcel/:pid", handlers.GetParcelByPID(repo))
	router.GET("parcel/date", handlers.GetParcelsByDate(repo))

	return router
}
