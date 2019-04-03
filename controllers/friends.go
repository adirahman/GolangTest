package controllers

import (
	"net/http"
	"payfazztest/structs"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetFriends(c *gin.Context) {
	var (
		friends []structs.Friend
		result  gin.H
	)
	id := c.Param("ktpId1")
	idb.DB.Where("ktpId1 = ?", id).Find(&friends)
	if len(friends) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": friends,
			"count":  len(friends),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateFriends(c *gin.Context) {
	var (
		friends structs.Friend
		result  gin.H
	)
	ktpId1 := c.PostForm("ktpId1")
	ktpId2 := c.PostForm("ktpId2")
	friends.ktpId1 = ktpId1
	friends.ktpId2 = ktpId2
	idb.DB.Create(&friends)
	result = gin.H{
		"result": friends,
	}
	c.JSON(http.StatusOK, result)
}
