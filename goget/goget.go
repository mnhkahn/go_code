package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"
)

type GoGet struct {
	Url           string
	Cnt           int
	Latch         int
	Header        http.Header
	MediaType     string
	MediaParams   map[string]string
	FilePath      string // 包括路径和文件名
	GetClient     *http.Client
	ContentLength int
	DownloadRange [][]int
	File          *os.File
	TempFiles     []*os.File
	WG            sync.WaitGroup
}

func NewGoGet() *GoGet {
	get := new(GoGet)
	get.FilePath = "./"
	get.GetClient = new(http.Client)

	flag.Parse()
	get.Url = *urlFlag
	get.Cnt = *cntFlag

	return get
}

var urlFlag = flag.String("u", "http://7b1h1l.com1.z0.glb.clouddn.com/bryce.jpg", "Fetch file url")
var cntFlag = flag.Int("c", 1, "Fetch concurrently counts")

func main() {
	get := NewGoGet()

	download_start := time.Now()

	req, err := http.NewRequest("HEAD", get.Url, nil)
	resp, err := get.GetClient.Do(req)
	get.Header = resp.Header
	if err != nil {
		log.Panicf("Get %s error %v.\n", get.Url, err)
	}
	get.MediaType, get.MediaParams, _ = mime.ParseMediaType(get.Header.Get("Content-Disposition"))
	get.ContentLength = int(resp.ContentLength)
	if strings.HasSuffix(get.FilePath, "/") {
		get.FilePath += get.MediaParams["filename"]
	}
	get.File, err = os.Create(get.FilePath)
	if err != nil {
		log.Panicf("Create file %s error %v.\n", get.FilePath, err)
	}
	log.Printf("Get %s MediaType:%s, Filename:%s, Length %d.\n", get.Url, get.MediaType, get.MediaParams["filename"], get.ContentLength)
	if get.Header.Get("Accept-Ranges") != "" {
		log.Printf("Server %s support Range by %s.\n", get.Header.Get("Server"), get.Header.Get("Accept-Ranges"))
	} else {
		log.Printf("Server %s doesn't support Range.\n", get.Header.Get("Server"))
	}

	log.Printf("Start to download %s with %d thread.\n", get.MediaParams["filename"], get.Cnt)
	range_start, range_interval := 0, int(math.Ceil(float64(get.ContentLength/get.Cnt)))
	for i := 0; i < get.Cnt; i++ {
		if i != get.Cnt-1 {
			get.DownloadRange = append(get.DownloadRange, []int{range_start, range_start + range_interval - 1})
		} else {
			// 最后一块
			get.DownloadRange = append(get.DownloadRange, []int{range_start, get.ContentLength - 1})
		}
		range_start += range_interval
	}
	// Check if the download has paused.
	for i := 0; i < len(get.DownloadRange); i++ {
		range_i := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
		temp_file, err := os.OpenFile(get.FilePath+"."+range_i, os.O_RDONLY|os.O_APPEND, 0)
		if err != nil {
			temp_file, _ = os.Create(get.FilePath + "." + range_i)
		} else {
			fi, err := temp_file.Stat()
			if err == nil {
				get.DownloadRange[i][0] += int(fi.Size())
			}
		}
		get.TempFiles = append(get.TempFiles, temp_file)
	}

	go get.Watch()
	get.Latch = get.Cnt
	for i, _ := range get.DownloadRange {
		get.WG.Add(1)
		go get.Download(i)
	}

	get.WG.Wait()

	for i := 0; i < len(get.TempFiles); i++ {
		temp_file, _ := os.Open(get.TempFiles[i].Name())
		cnt, err := io.Copy(get.File, temp_file)
		if cnt <= 0 || err != nil {
			log.Printf("Download #%d error %v.\n", i, err)
		}
		temp_file.Close()
	}
	get.File.Close()
	log.Printf("Download complete and store file %s with %v.\n", get.FilePath, time.Now().Sub(download_start))
	defer func() {
		for i := 0; i < len(get.TempFiles); i++ {
			err := os.Remove(get.TempFiles[i].Name())
			if err != nil {
				log.Printf("Remove temp file %s error %v.\n", get.TempFiles[i].Name(), err)
			} else {
				log.Printf("Remove temp file %s.\n", get.TempFiles[i].Name())
			}
		}
	}()
}

func (get *GoGet) Download(i int) {
	defer get.WG.Done()
	if get.DownloadRange[i][0] > get.DownloadRange[i][1] {
		return
	}
	range_i := fmt.Sprintf("%d-%d", get.DownloadRange[i][0], get.DownloadRange[i][1])
	log.Printf("Download #%d bytes %s.\n", i, range_i)

	defer get.TempFiles[i].Close()

	req, err := http.NewRequest("GET", get.Url, nil)
	req.Header.Set("Range", "bytes="+range_i)
	resp, err := get.GetClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Download #%d error %v.\n", i, err)
	} else {
		cnt, err := io.Copy(get.TempFiles[i], resp.Body)
		if cnt == int64(get.DownloadRange[i][1]-get.DownloadRange[i][0]+1) {
			log.Printf("Download #%d complete.\n", i)
		} else {
			req_dump, _ := httputil.DumpRequest(req, false)
			resp_dump, _ := httputil.DumpResponse(resp, true)
			log.Panicf("Download error %d %v, expect %d-%d, but got %d.\nRequest: %s\nResponse: %s\n", resp.StatusCode, err, get.DownloadRange[i][0], get.DownloadRange[i][1], cnt, string(req_dump), string(resp_dump))
		}
	}
}

// http://stackoverflow.com/questions/15714126/how-to-update-command-line-output
func (get *GoGet) Watch() {
	fmt.Printf("[=================>]\n")
}
