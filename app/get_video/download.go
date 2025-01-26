package download

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	geturl "newsbbc/app/get_url"
	"os"
)

type DataOfVideo struct {
	NameVideo   string
	Description string
}
type SaveDetailOfVideo []DataOfVideo

func OpenURL() (result geturl.DataSetModels) {
	workDir, _ := os.Getwd()
	file, _ := os.Open(workDir + "/urls.json")
	readFile, _ := io.ReadAll(file)
	decodeErr := json.Unmarshal(readFile, &result)
	if decodeErr != nil {
		log.Fatal(decodeErr)
		return
	}
	return
}
func OpenDetailVideo() (result SaveDetailOfVideo) {
	workDir, _ := os.Getwd()
	file, _ := os.Open(workDir + "/detail_video.json")
	readFile, _ := io.ReadAll(file)
	decodeErr := json.Unmarshal(readFile, &result)
	if decodeErr != nil {
		log.Fatal(decodeErr)
		return
	}
	return
}
func RemoveIndexZeroDetailVideo(data SaveDetailOfVideo) SaveDetailOfVideo {
	if len(data) == 0 {
		log.Fatal("data is empty")
		return data
	}
	return data[1:]
}
func WriteJSON(data SaveDetailOfVideo) {
	fmt.Println("====== write file json =======")
	dataSet, _ := json.Marshal(data)
	// fmt.Println(string(dataSet))
	workDir, _ := os.Getwd()
	writeFileErr := ioutil.WriteFile(workDir+"/detail_video.json", dataSet, 0644)
	if writeFileErr != nil {
		log.Fatal(writeFileErr)
	}

}
func DownloadContent() {
	geturl.GetDatasetFromAPIFY()
	// time.Sleep(time.Duration(80) * time.Second)
	saveDetailOfVideo := SaveDetailOfVideo{}
	for _, d := range OpenURL() {
		if d.VideoDuaration < 90 {
			nameVideo := Downloader(d.VideoUrl)
			saveDetailOfVideo = append(saveDetailOfVideo, DataOfVideo{
				NameVideo:   nameVideo,
				Description: d.Caption,
			})
		}
		continue
	}
	WriteJSON(saveDetailOfVideo)

	// save detail video
}
