package editvideo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func EditVideo(name_video string, new_name string) {
	log.Println("---> start editing")
	log.Println(name_video)
	log.Println(new_name)
	workDir, _ := os.Getwd()
	png := workDir + "/lena_news.png"
	fcParams := "[0:v][1:v]overlay=x=(W-w)/2:y=0"

	edit := exec.Command(
		"cpulimit", "--limit", "60", "--",
		"ffmpeg",
		"-i", name_video,
		"-i", png,
		"-filter_complex", fcParams,
		"-c:v", "libx264", "-preset", "ultrafast", "-crf", "35", // 🔥 Kurangi kualitas sedikit agar lebih cepat
		"-c:a", "copy",
		"-threads", "1",
		new_name, "-y",
	)
	// edit.Stdout = os.Stdout
	// edit.Stderr = os.Stderr
	cmdErr := edit.Run()
	
	if cmdErr != nil {
		log.Fatal(cmdErr)
		log.Fatal("error edit video")
	}
	fmt.Println("done...")
}
