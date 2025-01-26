package createsession

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Header struct {
	ContentType string
}
type CreateSessionModels struct {
	EndPoint        string
	PageAccessToken string
	UploadPhase     string
	Header          Header
}
type ResCreateSessionModels struct {
	VideoID   string `json:"video_id"`
	UploadUrl string `json:"upload_url"`
}

func (models CreateSessionModels) CreateSession() (ResCreateSessionModels, error) {
	urls := models.EndPoint
	// rebuild body Models
	bodys := map[string]string{
		"upload_phase": models.UploadPhase,
		"access_token": models.PageAccessToken,
	}
	bodyDecode, _ := json.Marshal(bodys)
	req, _ := http.NewRequest("POST", urls, bytes.NewBuffer(bodyDecode))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, _ := client.Do(req)
	readResp, _ := io.ReadAll(resp.Body)
	fmt.Println(string(readResp))
	resModels := ResCreateSessionModels{}
	decodeErr := json.Unmarshal(readResp, &resModels)
	if decodeErr != nil {
		log.Fatal("error create session upload")
		return resModels, decodeErr
	}
	resp.Body.Close()
	return resModels, nil
}
