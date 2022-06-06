package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"data":    posts,
	})
}

// @Summary 查询
// @Description 查询
// @Tags posts
// @Accept json
// @Produce json
// @Router /posts/{id} [get]
// @Param id path int true "pid"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
func Show(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"data":    post,
		})
	}

}

// @Summary 添加post
// @Description 添加post
// @Tags posts
// @Accept json
// @produce json
// @Router /posts [post]
// @Param content body string true "json"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Failure 500 {object} Response
func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}
	post.Status = "Active"
	db.Create(&post)
	c.JSON(http.StatusOK, Response{
		Msg:  "",
		Data: post,
	})
}

func Delete(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		c.JSON(http.StatusOK, Response{
			Msg:  "not found",
			Data: nil,
		})
	}
	db.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, Response{
		Msg:  "",
		Data: post,
	})
}
func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	/*if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "post not found",
			"data":    "",
		})
	}*/
	return post
}
