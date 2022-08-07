package main

import (
	"task-app-backend/handler"
	"task-app-backend/model"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=ec2-34-235-31-124.compute-1.amazonaws.com user=zhabzzpoprpqzr password=d4a1c36878582ec16046f0c3f1d818598ee5d579f671cc5108b04b3a64893f3c dbname=dm3bpjds7j83v port=5432 sslmode=require TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	taskRepository := model.NewRepository(db)
	taskHandler := handler.NewHandler(taskRepository)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", taskHandler.FetchAllHandler)
	router.POST("/", taskHandler.PostHandler)
	router.GET("/:id", taskHandler.FetchByIDHandler)
	router.POST("/update/:id", taskHandler.UpdateHandler)
	router.GET("/delete/:id", taskHandler.DeleteHandler)

	router.Run()
}
