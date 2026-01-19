package http

import (
	"net/http"
	"strconv"

	app "go_2601_04/internal/application/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *app.Service
}

func NewUserHandler(service *app.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", h.create)
		users.GET("", h.list)
		users.GET("/:id", h.get)
		users.PUT("/:id", h.update)
		users.DELETE("/:id", h.delete)
	}
}

func (h *UserHandler) create(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) list(c *gin.Context) {
	users, _ := h.service.List()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	c.ShouldBindJSON(&req)

	user, err := h.service.Update(uint(id), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
