package lib

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// EndsWith checks if the string ext ends with any of the suffixes in extss.
func EndsWith(ext string, extss string) bool {
	a := strings.Split(extss, " ")
	for _, ex := range a {
		if strings.HasSuffix(ext, ex) {
			return true
		}
	}
	return false
}
func Download(url, dataDir string) {
	// 生成一个新的 UUID
	newUuid := uuid.New()
	html := GetHtmlContent(url)
	WriteAllText(fmt.Sprintf("%s/%s.htm", dataDir, newUuid), html)
}

func GetHtmlContent(url string) string {
	methodName := "GetHtmlContent"
	fmt.Printf("Calling method: %s with arguments: %s\n", methodName, url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Request creation error: %s\n", err.Error())
		return ""
	}

	// 设置用户代理
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request error: %s\n", err.Error())
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return ""
	}

	htmlContent := string(body)
	if len(htmlContent) > 300 {
		fmt.Printf("Returning content: %s...\n", htmlContent[:300])
	} else {
		fmt.Printf("Returning content: %s...\n", htmlContent)
	}

	return htmlContent
}

func WriteAllText(filename, data string) {

	dir := filepath.Dir(filename)
	os.MkdirAll(dir, os.ModePerm)

	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err.Error())
	}

}
