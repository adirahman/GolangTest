package controllers

import (
	"payfazztest/structs"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetWorkHistory(c *gin.Context) {
	var (
		workHistory []structs.WorkHistory
		result      gin.H
	)

	id := c.Param("ktpId")
	idb.DB.Where("ktpId = ?", id).Find(&workHistory)
	if len(workHistory) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": workHistory,
			"count":  len(workHistory),
		}
	}
	c.JSON(http.Status, result)
}

func (idb *InDB) CreateWorkHistory(c *gin.Context) {
	var (
		work   structs.WorkHistory
		result gin.H
	)
	ktpId := c.PostForm("ktpId")
	companyID := c.PostForm("companyID")
	title := c.PostForm("title")
	workingStatus := c.PostForm("workingStatus")

	work.ktpId = ktpId
	work.companyID = companyID
	work.title = title
	work.workingStatus = workingStatus
	idb.DB.Create(&work)
	result = gin.H{
		"result": work,
	}
	c.JSON(http.StatusOK, result)
}
