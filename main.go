package main

import (
	"log"
	uploadreels "newsbbc/app/upload_reels"
	"time"
	"github.com/go-co-op/gocron"
	
	// download "newsbbc/app/get_video"
	// "os"
	// editvideo "newsbbc/app/edit_video"
	// download "newsbbc/app/get_video"
)

func main() {
	log.Println("standby")
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")
	schedules := gocron.NewScheduler(timeLoc)
	// schedules.Every(10).Day().At("01:30").Do(download.DownloadAllVideos)
	schedules.Every(1).Day().At("06:30").Do(uploadreels.RunUploadsReels)
	schedules.Every(1).Day().At("15:30").Do(uploadreels.RunUploadsReels)
	schedules.StartBlocking()
	
	// testing
	// download.GetContent()
	// geturl.GetDatasetFromAPIFY()
	// download.DownloadContent()
	// uploadreels.RunUploadsReels()
	// workDir, _ := os.Getwd()
	// editvideo.EditVideo(workDir+"/videos/video_upload5.mp4", workDir+"/videos/test.mp4")

}
