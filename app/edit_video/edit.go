package editvideo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func EditVideo(name_video string, new_name string) {
	xLabel := -150
	yLabel := -310
	workDir, _ := os.Getwd()
	png := workDir + "/lena_news.png"
	// fcParams := fmt.Sprintf("[1:v]scale=300:300[scaled];[0:v][scaled]overlay=x=%v:y=%v:enable='mod(t,1)'",
	fcParams := fmt.Sprintf("[1:v]scale=500:800[scaled];[0:v][scaled]overlay=x=%v:y=%v'",
		xLabel,
		yLabel)
	edit := exec.Command(
		"ffmpeg",
		"-i", name_video,
		"-i", png,
		"-filter_complex", fcParams,
		"-c:a", "copy",
		new_name, "-y",
	)
	// edit.Stdout = os.Stdout
	// edit.Stderr = os.Stderr
	cmdErr := edit.Run()
	if cmdErr != nil {
		log.Fatal("error edit video")
	}
	fmt.Println(">> Success Edit Video")
}
