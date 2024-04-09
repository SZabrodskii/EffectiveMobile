package config

import (
	"embed"
	"github.com/SZabrodskii/microservices/libs/utils"
)

//go:embed files
var configFile embed.FS

type CarsParams struct {
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
	RegNum       string `json:"reg_num"`
	Mark         string `json:"mark"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	OwnerName    string `json:"owner_name"`
	OwnerSurname string `json:"owner_surname"`
}

func GetCarsParams() CarsParams {
	return CarsParams{
		Page:         utils.GetEnv("PAGE", 1),
		PageSize:     utils.GetEnv("PAGE_SIZE", 10),
		RegNum:       utils.GetEnv("REG_NUM", ""),
		Mark:         utils.GetEnv("MARK", ""),
		Model:        utils.GetEnv("MODEL", ""),
		Year:         utils.GetEnv("YEAR", 0),
		OwnerName:    utils.GetEnv("OWNER_NAME", ""),
		OwnerSurname: utils.GetEnv("OWNER_SURNAME", ""),
	}
}
