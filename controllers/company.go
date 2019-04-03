package controllers

import (
	"net/http"
	"payfazztest/structs"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetCompany(c *gin.Context) {
	var (
		company structs.Company
		result  gin.H
	)
	id := c.Param("companyID")
	err := idb.DB.Where("companyID = ?", id).First(&company).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": company,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetCompanies(c *gin.Context) {
	var (
		companies []structs.Company
		result    gin.H
	)

	idb.DB.Find(&companies)
	if len(companies) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": companies,
			"count":  len(companies),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateCompany(c *gin.Context) {
	var (
		company structs.Company
		result  gin.H
	)
	companyID := c.PostForm("companyID")
	companyName := c.PostForm("companyName")
	company.companyID = companyID
	company.companyName = companyName
	idb.DB.Create(&company)
	result = gin.H{
		"result": company,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateCompany(c *gin.Context) {
	id := c.Query("companyID")
	companyName := c.PostForm("companyName")
	var (
		company    structs.Company
		newCompany structs.Company
		result     gin.H
	)

	err := idb.DB.First(&company, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newCompany.companyName = companyName
	err = idb.DB.Model(&company).Updates(newCompany).Error
	if err != nil {
		result = gin.H{
			"result": "updates failed",
		}
	} else {
		result = gin.H{
			"result": "successfully data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteCompany(c *gin.Context) {
	var (
		company structs.Company
		result  gin.H
	)
	id := c.Param("companyID")
	err := idb.DB.First(&company, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&company).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
