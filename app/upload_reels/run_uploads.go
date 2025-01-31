package uploadreels

import (
	"fmt"
	"log"
	"newsbbc/app"
	geturl "newsbbc/app/get_url"
	download "newsbbc/app/get_video"
	createsession "newsbbc/app/upload_reels/create_session"
	gettoken "newsbbc/app/upload_reels/get_token"
	"newsbbc/app/upload_reels/publish"
	uploadphase "newsbbc/app/upload_reels/upload_phase"
)

func RunUploadsReels() {
	log.Println("==== [ get content ] ===")
	filePath, caption, x := download.GetContent() // get video
	if !x {
		defer geturl.GetDatasetFromAPIFY()
		defer download.DownloadAllVideos()
	}
	
	log.Println("==== [ run upload ] ===")
	// Get List Of Token
	log.Println("---> get token")
	runGetToken := gettoken.GetAccessTokenModels{
		EndPoint:    fmt.Sprintf("%v/%v/accounts", app.Utils.BaseUrlGraphApi, app.Utils.UserID),
		AccessToken: app.Utils.SystemUserAccess,
	}
	listOfToken, err := runGetToken.GetAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range listOfToken.Data {
		if list.ID != "229092887266424" { // ID Lena News
			continue
		}

		log.Println("---> create session")
		// create a session from each page
		runCreateSession := createsession.CreateSessionModels{
			EndPoint: fmt.Sprintf("%v/%v/video_reels",
				app.Utils.BaseUrlGraphApi, // Base URL
				list.ID,                   // page id
			),
			PageAccessToken: list.AccessToken, // page access token
			UploadPhase:     "start",          // upload phase
			Header: createsession.Header{ // header
				ContentType: "application/json",
			},
		}
		getSession, getSessionErr := runCreateSession.CreateSession()
		if getSessionErr != nil {
			log.Fatal("get session error")
		}

		log.Println("---> upload video")
		// Upload Video
		runUpload := uploadphase.UploadPhaseModels{
			Rupload:     getSession,
			AccessToken: list.AccessToken,
		}
		uploader, _ := runUpload.UploadPhase(filePath) // upload process
		if !uploader.Success {
			log.Fatal("upload error")
		}

		log.Println("---> publish")
		// published
		publishParams := publish.PublishModelsParams{
			EndPoint:    fmt.Sprintf("%v/%v/video_reels", app.Utils.BaseUrlGraphApi, list.ID),
			AccessToken: list.AccessToken,
			VideoID:     getSession.VideoID,
			UploadPhase: "finish",
			VideoState:  "PUBLISHED",
			Description: caption,
		}
		publishParams.Publish()
		// defer os.Remove()
	}
}
