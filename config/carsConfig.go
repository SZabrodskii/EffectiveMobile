package config

import (
	"embed"
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

func GetCarsParamsFromEnv() CarsParams {
	return GetCarsParams{
		Page:         GetEnvInt("PAGE", 1),
		PageSize:     GetEnvInt("PAGE_SIZE", 10),
		RegNum:       GetEnv("REG_NUM", ""),
		Mark:         GetEnv("MARK", ""),
		Model:        GetEnv("MODEL", ""),
		Year:         GetEnvInt("YEAR", 0),
		OwnerName:    GetEnv("OWNER_NAME", ""),
		OwnerSurname: GetEnv("OWNER_SURNAME", ""),
	}
}
