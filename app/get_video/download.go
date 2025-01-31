package download

import (
	"encoding/json"
	"io"
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
	defer file.Close()
	readFile, _ := io.ReadAll(file)
	decodeErr := json.Unmarshal(readFile, &result)
	if decodeErr != nil {
		log.Fatal(decodeErr)
		return
	}
	return
}

func RemoveIndexZeroVideo(data geturl.DataSetModels) geturl.DataSetModels {
	if len(data) == 0 {
		log.Fatal("data is empty")
		return data
	}
	return data[1:]
}

func DownloadContent() (file_path_to_upload string, caption string) {
	urlsVideo := OpenURL()
	cpt := urlsVideo[0].Caption                       // get caption
	url := urlsVideo[0].VideoUrl                      // url of video
	fileName, _ := Downloader(url)                       // run download
	newDataUrlJson := RemoveIndexZeroVideo(urlsVideo) // remove index 0
	if len(newDataUrlJson) > 0 {
		geturl.WriteJSON(newDataUrlJson) // write after remove index 0
		return fileName, cpt
	}
	go geturl.GetDatasetFromAPIFY()
	return fileName, cpt
}
