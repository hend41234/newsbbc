package main

import (
	"log"
	uploadreels "newsbbc/app/upload_reels"
	"time"

	"github.com/go-co-op/gocron"
	// download "newsbbc/app/get_video"
	// "os"
	// download "newsbbc/app/get_video"
)

func main() {
	log.Println("standby")
	timeLoc, _ := time.LoadLocation("Asia/Jakarta")
	schedules := gocron.NewScheduler(timeLoc)
	// uploadreels.RunUploadsReels()
	// schedules.Every(10).Day().At("03:00").Do(download.DownloadAllVideos)
	schedules.Every(1).Day().At("00:53").Do(uploadreels.RunUploadsReels)
	schedules.Every(1).Day().At("06:40").Do(uploadreels.RunUploadsReels)
	schedules.Every(1).Day().At("17:40").Do(uploadreels.RunUploadsReels)
	// schedules.Every(1).Day().At("19:30").Do(uploadreels.RunUploadsReels)
	schedules.StartBlocking()

	// testing
	// workDir, _ := os.Getwd()
	// editvideo.EditVideo(workDir+"/videos/upload_video.mp4", workDir+"/videos/test.mp4")
	// download.GetContent()
	// geturl.GetDatasetFromAPIFY()
	// download.DownloadContent()
	// uploadreels.RunUploadsReels()

}
