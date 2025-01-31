package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Utilization struct {
	BaseURLAPIFY       string
	URLNewDatasetAPIFY string
	TokenAPIFY         string
	BaseUrlGraphApi    string
	UserID             string
	SecretAPP          string
	AppID              string
	SystemUserAccess   string
}

func GetEnv() (result Utilization) {
	workDir, _ := os.Getwd()
	loadEnv := godotenv.Load(fmt.Sprintf("%s/.env", workDir))
	if loadEnv == nil {
		//APIFY
		BaseUrlApify := os.Getenv("BASE_URL_APIFY")
		URLNewDatasetApify := os.Getenv("URL_NEW_DATASET_APIFY")
		TokenApify := os.Getenv("TOKEN_APIFY")
		// GRAPH API
		BaseUrlGraphApi := os.Getenv("BASE_URL_GRAPH_API")
		UserId := os.Getenv("USER_ID")
		SecretApp := os.Getenv("SECRET_APP")
		AppId := os.Getenv("APP_ID")
		SystemUserAccess := os.Getenv("SYSTEM_USER_ACCESS")

		return Utilization{
			BaseURLAPIFY:       BaseUrlApify,
			URLNewDatasetAPIFY: URLNewDatasetApify,
			TokenAPIFY:         TokenApify,
			BaseUrlGraphApi:    BaseUrlGraphApi,
			UserID:             UserId,
			SecretAPP:          SecretApp,
			AppID:              AppId,
			SystemUserAccess:   SystemUserAccess,
		}
	}
	return
}

var Utils = GetEnv()
