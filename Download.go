package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

func download(url string, filename string, w *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get -> %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll -> %s", err.Error())
		return
	}
	defer res.Body.Close()
	if err = ioutil.WriteFile(LLConfig.DownloadDestFolder+string(filepath.Separator)+filename, data, 0777); err != nil {
		log.Println("Error Saving:", filename, err)
	} else {
		//fmt.Println(time.Now().Format("2006-01-02 15:04:05 "),"Saved:", filename,"url:",url)
		log.Println("New version found,Download in " + LLConfig.DownloadDestFolder + filename)
	}
	w.Done()
}
