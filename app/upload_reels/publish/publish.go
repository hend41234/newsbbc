package publish

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	tag := strings.Join(Tag[:], " ")
	desc := fmt.Sprintf("%v\n.\n.\n.\n.\n.\n.\n.\n%v",params.Description, tag)
	paramsAdd := url.Values{}
	paramsAdd.Add("access_token", params.AccessToken)
	paramsAdd.Add("video_id", params.VideoID)
	paramsAdd.Add("upload_phase", params.UploadPhase)
	paramsAdd.Add("video_state", params.VideoState)
	paramsAdd.Add("description", desc)
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

var Tag = []string{
	"#news",
	"#lenews",
	"bbc",
	"#breakingnews",
	"#worldnews",
	"#localnews",
	"#newsupdate",
	"#newsalert",
	"#headlines",
	"#latestnews",
	"#newsreport",
	"#newsworthy",
	"#newsfeed",
	"#newsflash",
	"#newscoverage",
	"#newstoday",
	"#newsnow",
	"#newsroom",
	"#newschannel",
	"#newslive",
	"#newsnetwork",
	"#newsblog",
	"#newsmedia",
	"#newsstory",
	"#newsanchor",
	"#newspaper",
	"#newsbulletin",
	"#newsdesk",
	"#newscast",
	"#newsportal",
	"#newsagency",
	"#newswire",
	"#newsjournal",
	"#newsbroadcast",
	"#newsservice",
	"#newsreporter",
	"#newsphotography",
	"#newsphoto",
	"#newsjournalism",
	"#newscoverage",
	"#newsupdate",
	"#newsticker",
	"#newsheadline",
	"#newsapp",
	"#newsplatform",
	"#newswebsite",
	"#newsapp",
	"#newsaggregator",
	"#newsdigest",
}
