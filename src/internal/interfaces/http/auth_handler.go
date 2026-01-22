package http

import (
	"fmt"
	app "go_2601_04/internal/application/auth"
	"go_2601_04/pkg/response"
	"go_2601_04/pkg/validator"

	"github.com/davecgh/go-spew/spew"
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
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}
	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		c.JSON(422, response.ApiResponse[any]{
			Success: false,
			Errors:  errs,
			Message: "Dữ liệu đầu vào không hợp lệ",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.ApiResponse[any]{Success: false, Message: "Sai định dạng JSON", Errors: err.Error()})
		return
	}
	fmt.Println("==========req")
	accessToken, err := h.service.Login(req.Email, req.Password)
	spew.Dump(err)
	if err != nil {
		response.Error(
			c,
			401,
			response.CodeUnauthorized,
			"Username or password not valid",
			nil,
		)
		return
	} else {
		data := make(map[string]any)
		data["accessToken"] = accessToken
		response.OK(c, data)
	}

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
