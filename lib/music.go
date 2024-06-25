package lib

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// DownloadVideo 下载视频并返回文件名
func DownloadVideo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("无法下载视频: %v", err)
	}
	defer resp.Body.Close()

	file, err := os.Create("video.mp4")
	if err != nil {
		return "", fmt.Errorf("无法创建文件: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("无法保存视频: %v", err)
	}

	return "video.mp4", nil
}

// ConvertToMP3 将视频文件转换为 MP3
func ConvertToMP3(videoFile string) (string, error) {
	mp3File := "output.mp3"
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-q:a", "0", "-map", "a", mp3File)
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("无法转换为 MP3: %v", err)
	}

	return mp3File, nil
}

// SearchAndDownloadSong 搜索歌曲并下载 MP3
func SearchAndDownloadSong(apiKey, query string) error {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("无法创建 YouTube 服务: %v", err)
	}

	call := service.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return fmt.Errorf("搜索失败: %v", err)
	}

	if len(response.Items) == 0 {
		return fmt.Errorf("未找到相关视频")
	}

	videoID := response.Items[0].Id.VideoId
	videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

	videoFile, err := DownloadVideo(videoURL)
	if err != nil {
		return fmt.Errorf("下载视频失败: %v", err)
	}

	_, err = ConvertToMP3(videoFile)
	if err != nil {
		return fmt.Errorf("转换为 MP3 失败: %v", err)
	}

	fmt.Println("下载并转换为 MP3 成功")
	return nil
}

func main() {
	apiKey := "YOUR_YOUTUBE_API_KEY"
	songName := "Shape of You Ed Sheeran"
	err := SearchAndDownloadSong(apiKey, songName)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Println("下载成功")
	}
}
