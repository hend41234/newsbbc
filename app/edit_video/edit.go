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
	// png := workDir + "/lenanews_wm.png"
	frame := workDir + "/lena_news_frame.png"
	audio := workDir + "/lena_news_sound.aac"
	fcParams := "[1:v]crop=iw:ih*8/10:(iw-iw/2)/2:(ih-ih*8/10)/2," +
		"scale=820:-1," +
		"pad=iw+20:ih+20:10:10:grey[main];" +
		"[0:v][main]overlay=(W-w)/2:(H-h)/2[vout];" +
		"[2:a]volume=0.3[aud];" +
		"[aud][1:a]amix=inputs=2:duration=shortest:dropout_transition=2[aout];"

	edit := exec.Command(
		"cpulimit", "--limit", "30", "-f", "--",
		"ffmpeg",
		"-i", frame,
		"-i", name_video,
		"-stream_loop", "-1", "-i", audio,
		// "-stream_loop", "-1",
		"-map", "[vout]", "-map", "[aout]",
		"-t", fmt.Sprintf("%v", GetDuration(name_video)-2),
		"-filter_complex", fcParams,
		"-c:v", "libx264", "-preset", "ultrafast", "-crf", "35", // 🔥 Kurangi kualitas sedikit agar lebih cepat
		"-c:a", "aac", "-b:a", "192k",
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
}
