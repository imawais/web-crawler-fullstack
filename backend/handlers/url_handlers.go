package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imawais/web-crawler-fullstack/backend/database"
  "github.com/imawais/web-crawler-fullstack/backend/crawler"
	"net/http"
)

type AddURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func AddURL(c *gin.Context) {
	var input AddURLRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	res, err := database.DB.Exec("INSERT INTO urls (url, status) VALUES (?, 'queued')", input.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	id, _ := res.LastInsertId()

	// Trigger background crawling
	go func(url string, id int) {
		_ = database.DB.Exec("UPDATE urls SET status = 'running' WHERE id = ?", id)
		err := crawler.CrawlAndStore(url, int(id))
		if err != nil {
			_ = database.DB.Exec("UPDATE urls SET status = 'error' WHERE id = ?", id)
		}
	}(input.URL, int(id))

	c.JSON(http.StatusCreated, gin.H{"url": input.URL, "status": "queued", "id": id})
}


func ListURLs(c *gin.Context) {
	var urls []map[string]interface{}
	err := database.DB.Select(&urls, "SELECT * FROM urls ORDER BY id DESC LIMIT 100")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB fetch failed"})
		return
	}
	c.JSON(http.StatusOK, urls)
}
