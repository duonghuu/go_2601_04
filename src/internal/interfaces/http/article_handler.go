package http

import (
	"net/http"
	"strconv"

	app "go_2601_04/internal/application/article"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	service *app.Service
}

func NewArticleHandler(service *app.Service) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) Register(r *gin.Engine) {
	articles := r.Group("/articles")
	{
		articles.POST("", h.create)
		articles.GET("", h.list)
		articles.GET("/:id", h.get)
		articles.PUT("/:id", h.update)
		articles.DELETE("/:id", h.delete)
	}
}

func (h *ArticleHandler) create(c *gin.Context) {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Create(req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) list(c *gin.Context) {
	articles, _ := h.service.List()
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Title  string `json:"title"`
		Content string `json:"content"`
	}
	c.ShouldBindJSON(&req)

	article, err := h.service.Update(uint(id), req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
