package lib

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Spdr struct {
	urlsDownWaitHashset map[string]struct{}
	downedUrlsHashset   map[string]struct{}
	parserUrlQueue      string
}

func (s *Spdr) SpdrTest() {

	s.urlsDownWaitHashset = NewSet(fmt.Sprintf("spdr/urlsDownWait%s.json", time.Now().Format("02_150405")))
	s.downedUrlsHashset = NewSet("spdr/downedUrls.json")
	s.parserUrlQueue = "spdr/downHtmTaskQue"
	//	timestamp := time.Now().Format("20060102_150405_000")
	startURL := "https://www.khmertimeskh.com/"

	startURL = "https://laotiantimes.com/"
	startURL = "https://laotiantimes.com/category/economy/"
	startURL = "https://laotiantimes.com/category/business/"
	startURL = "https://www.bangkokpost.com/"
	startURL = "https://myanmar-now.org/en/"
	startURL = "https://e.vnexpress.net/"
	startURL = "https://www.bangkokpost.com/"

	s.urlsDownWaitHashset[startURL] = struct{}{}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(3 * time.Second)
			s.downloadTask(startURL)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(3 * time.Second)
			s.parseHtmlFileTask(startURL)
		}
	}()

	wg.Wait()
}

func (s *Spdr) downloadTask(startURL string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in downloadTask", r)
		}
	}()
	//	return
	for url := range s.urlsDownWaitHashset {
		_, found := s.downedUrlsHashset[url]
		if !found {
			url1 := startURL + url
			if strings.HasPrefix(url, "http") {
				url1 = url
			}

			go func(url string) {
				// Implement the Download function according to your need
				Download(url, s.parserUrlQueue)
			}(url1)

			s.downedUrlsHashset[url] = struct{}{}
			s.downedUrlsHashset[url1] = struct{}{}
		}
	}

	for url := range s.downedUrlsHashset {
		delete(s.urlsDownWaitHashset, url)
	}
}

func (s *Spdr) parseHtmlFileTask(startURL string) {
	files, err := ioutil.ReadDir(s.parserUrlQueue)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(s.parserUrlQueue, file.Name())
		html, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		urls := ExtractHrefAttributes(string(html))
		for url := range urls {
			ext := filepath.Ext(url)
			if EndsWith(ext, "js css jpg png gif ico jpeg mp3 mp4") {
				continue
			}
			url1 := startURL + url
			if strings.HasPrefix(url, "http") {
				url1 = url
			}
			s.urlsDownWaitHashset[url1] = struct{}{}
		}

		newPath := filepath.Join("spdr", "downHtmldirLog", file.Name())
		MoveFileToDirectory(filePath, newPath)
	}
}

func main22() {
	//spdr := &Spdr{
	//
	//}

	//spdr.SpdrTest()
}
