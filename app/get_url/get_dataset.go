package geturl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"newsbbc/app"
	"os"
	"time"
)

// type DatasetRespModels struct{
//
// }

type SliceDatasetModels struct {
	VideoUrl       string  `json:"videoUrl"`
	Caption        string  `json:"caption"`
	VideoDuaration float64 `json:"videoDuration"`
}

type DataSetModels []SliceDatasetModels

func WriteJSON(data DataSetModels) {
	fmt.Println("====== write file json =======")
	dataSet, _ := json.Marshal(data)
	workDir, _ := os.Getwd()
	writeFileErr := ioutil.WriteFile(workDir+"/urls.json", dataSet, 0644)
	if writeFileErr != nil {
		log.Fatal(writeFileErr)
	}
	log.Println("done...")
}
func FilterDurationVideo(data DataSetModels) {
	validVideo := DataSetModels{}
	for _, list := range data {
		if list.VideoDuaration < 90 {
			validVideo = append(validVideo, list)
		}
	}
	WriteJSON(validVideo)
}
func GetDatasetFromAPIFY() {
	datasetID := NewDatasetToAPIFY()
	time.Sleep(time.Duration(900) * time.Second)
	// datasetID := "XCsIpcQlQRuBGV3KZ"
	fmt.Println("=========== Get Dataset ===========")
	urls := fmt.Sprintf("%s/datasets/%s/items",
		app.Utils.BaseURLAPIFY, //base urls
		datasetID,              // dataset ID
	)
	req, _ := http.NewRequest("GET", urls, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	readResp, _ := io.ReadAll(resp.Body)
	datasetModel := DataSetModels{}
	decodeErr := json.Unmarshal(readResp, &datasetModel)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	FilterDurationVideo(datasetModel)
	fmt.Println("Done..")

}
