package main

import (
	"task-app-backend/handler"
	"task-app-backend/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=ec2-34-235-31-124.compute-1.amazonaws.com user=zhabzzpoprpqzr password=d4a1c36878582ec16046f0c3f1d818598ee5d579f671cc5108b04b3a64893f3c dbname=dm3bpjds7j83v port=5432"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	taskRepository := model.NewRepository(db)
	taskHandler := handler.NewHandler(taskRepository)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", taskHandler.FetchAllHandler)
	v1.POST("/", taskHandler.PostHandler)
	v1.GET("/:id", taskHandler.FetchByIDHandler)
	v1.PATCH("/:id", taskHandler.UpdateHandler)
	v1.DELETE("/:id", taskHandler.DeleteHandler)

	router.Run()
}
