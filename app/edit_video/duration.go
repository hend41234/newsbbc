package editvideo

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetDuration(name_video string) (secon float64) {
	cmd := exec.Command("ffmpeg", "-i", name_video)
	output, _ := cmd.CombinedOutput()
	// extract suration
	dur := ""
	for _, line := range strings.Split(string(output), "\n") {
		fmt.Println(line)
		if strings.Contains(line, "Duration") {
			dur = strings.TrimSpace(strings.Split(line, ",")[0][10:])
			break
		}
	}
	if dur == "" {
		log.Println("could'nt get duration")
		return 0
	}
	var hours, minutes, secons float64
	_, err := fmt.Sscanf(dur, ": %02f:%02f:%05f", &hours, &minutes, &secons)
	if err != nil {
		log.Println("error formatting", err)
	}
	totalSecon := hours*3600 + minutes*60 + secons
	// return fmt.Sprintf("%v", totalSecon)
	return totalSecon
}
