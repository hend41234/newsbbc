package uploadphase

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	createsession "newsbbc/app/upload_reels/create_session"
	"os"
	"strconv"
)

type Header struct {
	Authorization string
	Offset        int
	FileSize      int
	ContentType   string
}

type UploadPhaseModels struct {
	Rupload     createsession.ResCreateSessionModels
	AccessToken string
}

type ResUploadPhaseModels struct {
	Success bool
}

func (models UploadPhaseModels) UploadPhase(videos string) (ResUploadPhaseModels, error) {
	urls := models.Rupload.UploadUrl
	var Headers = Header{
		Authorization: fmt.Sprintf("OAuth %v", models.AccessToken),
		Offset:        0,
		FileSize:      GetFileSize(videos),
		// ContentType:   "application/octet-stream",
	}
	video, _ := os.Open(videos)
	req, _ := http.NewRequest("POST", urls, video)
	req.Header.Set("Authorization", Headers.Authorization)
	req.Header.Set("offset", strconv.Itoa(Headers.Offset))
	req.Header.Set("file_size", strconv.Itoa(Headers.FileSize))
	// req.Header.Set("Content-Type", Headers.ContentType)
	req.ContentLength = int64(Headers.FileSize)

	client := http.Client{}
	resp, _ := client.Do(req)
	resModels := ResUploadPhaseModels{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&resModels)
	if decodeErr != nil {
		log.Fatal("decode error")
		return resModels, decodeErr
	}
	// fmt.Println(string(readResp))
	return resModels, nil

}
