package gettoken

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetAccessTokenModels struct {
	EndPoint    string
	AccessToken string
}

type DataModels struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	ID          string `json:"ID"`
}
type ResGetAccessTokenModels struct {
	Data []DataModels
}

func (models GetAccessTokenModels) GetAccessToken() (ResGetAccessTokenModels, error) {
	urls := fmt.Sprintf("%v?access_token=%s",
		models.EndPoint,    //Base URL
		models.AccessToken, //access_token 'USER ACCOUNT TOKEN'
	)
	req, _ := http.NewRequest("GET", urls, nil)
	client := http.Client{}
	resp, _ := client.Do(req)
	resModel := ResGetAccessTokenModels{}
	readRes, _ := io.ReadAll(resp.Body)
	decodeErr := json.Unmarshal(readRes, &resModel)
	resp.Body.Close()

	if decodeErr != nil {
		log.Fatal("error get token")
		return resModel, decodeErr
	}
	return resModel, nil
}
