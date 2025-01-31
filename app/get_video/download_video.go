package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// editvideo "newsbbc/app/edit_video"
	"os"
	"time"
)

func downloadVideo(url, filename string) error {
	// Buat permintaan HTTP ke URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Buka file untuk menulis hasil download
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Salin konten dari HTTP response ke file
	_, err = io.Copy(out, resp.Body)
	return err
}

func Downloader(url string) (filename string, x bool) {
	// url := "https://scontent-jnb2-1.cdninstagram.com/o1/v/t16/f2/m86/AQMKyZn9XhR7lyerN39COo1t3Lqplb-guVaqJaNm8eNaA_lHk4rDEPPc5WwlSqaT1hu6BXwIlEZsGFsHmX3TM9V1lS6eDOOOvbZUuWw.mp4?stp=dst-mp4\u0026efg=eyJxZV9ncm91cHMiOiJbXCJpZ193ZWJfZGVsaXZlcnlfdnRzX290ZlwiXSIsInZlbmNvZGVfdGFnIjoidnRzX3ZvZF91cmxnZW4uY2xpcHMuYzIuNzIwLmJhc2VsaW5lIn0\u0026_nc_cat=107\u0026vs=605662245398925_843808722\u0026_nc_vs=HBksFQIYUmlnX3hwdl9yZWVsc19wZXJtYW5lbnRfc3JfcHJvZC81ODRDRDZDRTY2NEQ1NDY4REE3RjUwN0ZBQUNEOTZBNl92aWRlb19kYXNoaW5pdC5tcDQVAALIAQAVAhg6cGFzc3Rocm91Z2hfZXZlcnN0b3JlL0dFVEFDaHkta1Y1M2FqQUNBQUhnbGt6bzc5SmhicV9FQUFBRhUCAsgBACgAGAAbABUAACb0ub%2BFsKOTQBUCKAJDMywXQFYj1wo9cKQYEmRhc2hfYmFzZWxpbmVfMV92MREAdf4HAA%3D%3D\u0026_nc_rid=fd691897d8\u0026ccb=9-4\u0026oh=00_AYCwBn0l38_6LKy2LxMoM1boS_93YU1Iafb99uEF4lFbaQ\u0026oe=6796D969\u0026_nc_sid=10d13b" // Ganti dengan URL video
	fmt.Println("running download")
	workDir, _ := os.Getwd()
	dir, _ := os.ReadDir(workDir+"/videos")
	filename = fmt.Sprintf("videos/video_upload%v.mp4", len(dir))
	err := downloadVideo(url, filename)
	if err != nil {
		log.Printf("Gagal mengunduh video: %v", filename)
		return filename, false
	}
	log.Println(">>> Video berhasil diunduh :", filename)
	time.Sleep(time.Duration(5) * time.Second)
	return filename, true
}
