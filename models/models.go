package models

import "gorm.io/gorm"

type Fact struct{
	//struct tags 
	gorm.Model 
	Question string `json:"question" gorm:"text;not null default:null`
	Answer string `json:"answer" gorm:"text;not null default:null` 
}