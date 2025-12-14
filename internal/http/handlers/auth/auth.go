package auth

import (
	"net/http"
	"task-manager/internal/http/dto"
	"task-manager/internal/models"
	"task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *services.AuthService
}

func NewHandler(service *services.AuthService) *Handler {
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

func (h *Handler) Register(c *gin.Context) {
	var req dto.CreateAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := h.Service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.CreateAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := h.Service.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
