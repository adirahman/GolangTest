package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	EktpID    string
	firstName string
	lastName  string
}

type Company struct {
	gorm.Model
	companyID   string
	companyName string
}

type Friend struct {
	gorm.Model
	ktpId1 string
	ktpId2 string
}

type WorkHistory struct {
	gorm.Model
	ktpId         string
	companyID     string
	title         string
	workingStatus string
}
