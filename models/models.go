package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	RegNum   string `json:"regNum"`
	Mark     string `json:"mark"`
	CarModel string `json:"model" gorm:"column:model"`
	Year     int    `json:"year"`
	Owner    People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
