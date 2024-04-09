package controllers

import (
	"EffectiveMobile/config"
	"EffectiveMobile/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Repository struct {
	DB *gorm.DB
}

func GetIntQueryParam(c *gin.Context, key string, defaultValue int) int {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func (r *Repository) SetupRoutes(router *gin.Engine) {
	router.GET("/cars", r.GetCarsHandler)
	router.DELETE("/cars/:id", r.DeleteCarHandler)
	router.PUT("/cars/:id", r.UpdateCarHandler)
	router.POST("/cars", r.AddCarHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *Repository) GetCarsHandler(c *gin.Context) {
	// Получение параметров фильтрации и пагинации из запроса
	params := config.CarsParams{
		Page:         GetIntQueryParam(c, "page", 1),
		PageSize:     GetIntQueryParam(c, "page_size", 10),
		RegNum:       c.Query("reg_num"),
		Mark:         c.Query("mark"),
		Model:        c.Query("model"),
		Year:         GetIntQueryParam(c, "year", 0),
		OwnerName:    c.Query("owner_name"),
		OwnerSurname: c.Query("owner_surname"),
	}

	cars := []models.Car{}
	if err := r.DB.Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Where(&models.Car{
		RegNum:   params.RegNum,
		Mark:     params.Mark,
		CarModel: params.Model,
		Year:     params.Year,
		Owner: models.People{
			Name:       params.OwnerName,
			Surname:    params.OwnerSurname,
			Patronymic: "",
		},
	}).Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (r *Repository) DeleteCarHandler(c *gin.Context) {
	id := c.Param("id")

}

func (r *Repository) UpdateCarHandler(c *gin.Context) {
	id := c.Param("id")

}

func (r *Repository) AddCarHandler(c *gin.Context) {

}
