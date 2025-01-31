package download

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	editvideo "newsbbc/app/edit_video"
	geturl "newsbbc/app/get_url"
	"time"
	"os"
)

type NewVideoModels struct {
	Path     string
	Caption  string
	Duration float64
}
type NewModels []NewVideoModels

func (data NewModels) WriteJSON() {
	fmt.Println("---> write data content json")
	dataSet, _ := json.Marshal(data)
	workDir, _ := os.Getwd()
	writeFileErr := ioutil.WriteFile(workDir+"/content.json", dataSet, 0644)
	if writeFileErr != nil {
		log.Fatal(writeFileErr)
	}
	log.Println("---> done...")
}

func (data NewModels) RemoveIndex() NewModels {
	if len(data) == 0 {
		log.Println("data is empty")
		return NewModels{}
	}
	return data[1:]
}

func DownloadAllVideos() {
	models := NewModels{}
	oldUrl := OpenURL()
	for _, list := range oldUrl {
		fileName, x := Downloader(list.VideoUrl)
		if !x{
			continue
		}
		tempData := NewVideoModels{
			Path:     fileName,
			Caption:  list.Caption,
			Duration: list.VideoDuaration,
		}
		models = append(models, tempData)
	}
	models.WriteJSON()
}

func OpenContent() (result NewModels) {
	workDir, _ := os.Getwd()
	file, _ := os.Open(workDir + "/content.json")
	defer file.Close()
	readFile, _ := ioutil.ReadAll(file)
	decodeErr := json.Unmarshal(readFile, &result)
	if decodeErr != nil {
		log.Fatal(decodeErr)
		return
	}
	return
}

func GetContent() (file_path string, caption string, x bool) {
	content := OpenContent()
	// file_path = content[0].Path
	workDir, _ := os.Getwd()
	file_path = workDir + "/videos/upload_video.mp4"
	if len(content) == 0{
		log.Println("content is empty, we must download content first, wait for it.")
		geturl.GetDatasetFromAPIFY()
		DownloadAllVideos()
		content = OpenContent()
	}
	videoToEdit := fmt.Sprintf("%v/%v",workDir,content[0].Path)
	caption = content[0].Caption
	editvideo.EditVideo(videoToEdit, file_path)
	time.Sleep(time.Duration(60) * time.Second)
	defer os.Remove(videoToEdit)
	remove := content.RemoveIndex()
	remove.WriteJSON()
	if len(remove) == 0 {
		x = false
	} else{
		x=true
	}
	return
}
