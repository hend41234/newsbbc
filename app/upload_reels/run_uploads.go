package uploadreels

import (
	"fmt"
	"log"
	"newsbbc/app"
	download "newsbbc/app/get_video"
	createsession "newsbbc/app/upload_reels/create_session"
	gettoken "newsbbc/app/upload_reels/get_token"
	"newsbbc/app/upload_reels/publish"
	uploadphase "newsbbc/app/upload_reels/upload_phase"
	"os"
)

func RunUploadsReels() {
	// Get List Of Token
	runGetToken := gettoken.GetAccessTokenModels{
		EndPoint:    fmt.Sprintf("%v/%v/accounts", app.Utils.BaseUrlGraphApi, app.Utils.UserID),
		AccessToken: app.Utils.SystemUserAccess,
	}
	listOfToken, err := runGetToken.GetAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range listOfToken.Data {

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
		// Upload Video
		// open detail_video.json
		videoList := download.OpenDetailVideo()
		detailVideo := videoList[0]
		runUpload := uploadphase.UploadPhaseModels{
			Rupload:     getSession,
			AccessToken: list.AccessToken,
		}
		filePath := detailVideo.NameVideo              // final filepath of video
		uploader, _ := runUpload.UploadPhase(filePath) // upload process
		if !uploader.Success {
			log.Fatal("upload error")
		}
		// published
		publishParams := publish.PublishModelsParams{
			EndPoint:    fmt.Sprintf("%v/%v/video_reels", app.Utils.BaseUrlGraphApi, list.ID),
			AccessToken: list.AccessToken,
			VideoID:     getSession.VideoID,
			UploadPhase: "finish",
			VideoState:  "PUBLISHED",
			Description: detailVideo.Description,
		}
		publishParams.Publish()
		rm := download.RemoveIndexZeroDetailVideo(videoList)
		download.WriteJSON(rm)
		os.Remove(filePath)
	}
}
