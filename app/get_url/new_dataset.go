qpackage geturl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"newsbbc/app"
)

type DataModels struct {
	DefaultDatasetId string `json:"defaultDatasetId"`
}

type ResModels struct {
	Data DataModels `json:"data"`
}

var Body = []byte(`{
    "addParentData": false,
    "directUrls": [
        "https://www.instagram.com/bbcindonesia/"
    ],
    "enhanceUserSearchWithFacebookPage": false,
    "isUserReelFeedURL": false,
    "isUserTaggedFeedURL": false,
    "onlyPostsNewerThan": "1 week",
    "resultsLimit": 50,
    "resultsType": "stories",
    "searchLimit": 1,
    "searchType": "hashtag"
}`)

func NewDatasetToAPIFY() (dataset_id string) {
	utils := app.Utils
	body := Body
	urls := fmt.Sprintf("%s?token=%s", utils.URLNewDatasetAPIFY, utils.TokenAPIFY)
	req, _ := http.NewRequest("POST", urls, bytes.NewBuffer(body))
	// fmt.Println(bytes.NewBuffer(body))
	var respModels ResModels

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, _ := client.Do(req)
	readRespone, _ := io.ReadAll(resp.Body)
	unmarshalErr := json.Unmarshal(readRespone, &respModels)
	if unmarshalErr != nil {
		log.Fatal("error")
	}
	resp.Body.Close()
	// fmt.Println(respModels.Data.DefaultDatasetId)
	return respModels.Data.DefaultDatasetId
}
