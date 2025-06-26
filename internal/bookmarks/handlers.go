package bookmarks

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/bookmarks", getAll)
	r.GET("/bookmarks/:id", getByID)
	r.POST("/bookmarks", create)
	r.PUT("/bookmarks/:id", update)
	r.DELETE("/bookmarks/:id", delete)
}

func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, GetAll())
}

func getByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	b, err := Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

func create(c *gin.Context) {
	var b Bookmark
	if err := c.ShouldBind(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Create(b))
}

func update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookmark not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
