package download
import (
	"log"
	"strings"
	"io"
	"os"
	"github.com/kkdai/youtube/v2"
)
const YOUTUBE_BASE_LINK = "https://www.youtube.com/watch?v="

func Mp4_download(youtube_url string) {
	youtube_video_id := extract_youtube_video_id(youtube_url)
	log.Print("----------")
  log.Print(youtube_video_id)
  log.Print("----------")
	client := youtube.Client{}
	video, error := client.GetVideo(youtube_video_id)
	log.Print("----------")
  log.Print(video)
  log.Print("----------")
	if error != nil {
		log.Fatal(error)
	}
	formats := video.Formats.WithAudioChannels()
	stream, _, error := client.GetStream(video, &formats[0])
	if error != nil {
		log.Fatal(error)
	}
	file, error := os.Create("tmp/video.mp4")
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	_, error = io.Copy(file, stream)
	if error != nil {
		log.Fatal(error)
	}
}

func extract_youtube_video_id(youtube_url string) string {
	youtube_video_id := strings.Replace(youtube_url, YOUTUBE_BASE_LINK, "", 1)
	return youtube_video_id
}
