package task

import (
	"net/http"
	"task-manager/internal/http/dto"
	"task-manager/internal/models"
	"task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.TaskService
}

func NewHandler(service *services.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetAll(c *gin.Context) {
	tasks, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := models.Task{
		Title:  req.Title,
		UserID: req.UserID,
	}

	if err := h.Service.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}
