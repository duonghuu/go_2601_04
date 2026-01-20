package http

import (
	app "go_2601_04/internal/application/auth"
	"go_2601_04/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *app.Service
}

func NewAuthHandler(service *app.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/logout", h.logout)
	}
}

func (h *AuthHandler) login(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, hashPassword)
	
	if !isValidPassword {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "fail"})
}

func (h *AuthHandler) logout(c *gin.Context) {}

// func (h *UserHandler) create(c *gin.Context) {
// 	var req struct {
// 		Name  string `json:"name"`
// 		Email string `json:"email"`
// 	}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := h.service.Create(req.Name, req.Email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, user)
// }

// func (h *UserHandler) list(c *gin.Context) {
// 	users, _ := h.service.List()
// 	c.JSON(http.StatusOK, users)
// }

// func (h *UserHandler) get(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user, err := h.service.Get(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (h *UserHandler) update(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	var req struct {
// 		Name  string `json:"name"`
// 		Email string `json:"email"`
// 	}
// 	c.ShouldBindJSON(&req)

// 	user, err := h.service.Update(uint(id), req.Name, req.Email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (h *UserHandler) delete(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	if err := h.service.Delete(uint(id)); err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }
