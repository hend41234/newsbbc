package publish

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// var idTest = "512951895152794"

type PublishModelsParams struct {
	EndPoint    string
	AccessToken string
	VideoID     string
	UploadPhase string
	VideoState  string
	Description string
}

type ResPublishModels struct {
	Success bool `json:"success"`
}

func (params PublishModelsParams) Publish() {
	urls := params.EndPoint
	paramsAdd := url.Values{}
	paramsAdd.Add("access_token", params.AccessToken)
	paramsAdd.Add("video_id", params.VideoID)
	paramsAdd.Add("upload_phase", params.UploadPhase)
	paramsAdd.Add("video_state", params.VideoState)
	paramsAdd.Add("description", params.Description)
	finalURL := fmt.Sprintf("%s?%s", urls, paramsAdd.Encode())

	req, _ := http.NewRequest("POST", finalURL, nil)
	client := http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Fatal("published error")
	}
	resModels := ResPublishModels{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&resModels)
	if decodeErr != nil {
		log.Fatal("error decode")
	}
	fmt.Println(resModels.Success)
}
