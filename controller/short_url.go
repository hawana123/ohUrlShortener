package controller

import (
	"net/http"
	"ohurlshortener/service"
	"ohurlshortener/utils"

	"github.com/gin-gonic/gin"
)

func ShortUrlDetail(c *gin.Context) {
	url := c.Param("url")
	if utils.EemptyString(url) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title":   "404 - ohUrlShortener",
			"code":    http.StatusNotFound,
			"message": "您访问的页面已失效",
			"label":   "Status Not Found",
		})
		return
	}

	destUrl, err := service.Search4ShortUrl(url)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"title":   "内部错误 - ohUrlShortener",
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"label":   "Error",
		})
		return
	}

	if utils.EemptyString(destUrl) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title":   "404 - ohUrlShortener",
			"code":    http.StatusNotFound,
			"message": "您访问的页面已失效",
			"label":   "Status Not Found",
		})
		return
	}

	go service.NewAccessLog(url, c.ClientIP(), c.Request.UserAgent(), c.Request.Referer()) //TODO: add more params to access logs

	c.Redirect(http.StatusFound, destUrl)
}
