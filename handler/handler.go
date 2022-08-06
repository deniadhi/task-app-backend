package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task-app-backend/model"
)

type taskHandler struct {
	repository model.Repository
}

func NewHandler (taskRepository model.Repository) *taskHandler {
	return &taskHandler{taskRepository}
}

func (handler *taskHandler) FetchAllHandler (c *gin.Context) {
	data, _ := handler.repository.FetchAll()

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (handler *taskHandler) PostHandler (c *gin.Context) {
	var input model.Input
	c.ShouldBindJSON(&input)
	handler.repository.Save(input)

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *taskHandler) FetchByIDHandler (c *gin.Context) {
	ID := c.Param("id")
	data, _ := handler.repository.FetchByID(ID)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (handler *taskHandler) UpdateHandler (c *gin.Context) {
	ID := c.Param("id")
	var input model.Input
	c.ShouldBindJSON(&input)

	handler.repository.Update(ID, input)

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *taskHandler) DeleteHandler (c *gin.Context) {
	ID := c.Param("id")
	handler.repository.Delete(ID)
	c.JSON(http.StatusOK, gin.H{})
}