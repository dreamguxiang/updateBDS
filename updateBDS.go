package main

import (
	"C"
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var w sync.WaitGroup
var ww sync.WaitGroup
var build strings.Builder

//export onPostInit
func onPostInit() {
	inits()
	own, _ := FileChecker("plugins\\BDSUpdate\\BDSUpdate.json")
	if own == false {
		weitejson()
	}
	LLConfig = ReadConfig()
	first()
	//ww.Add(1)
	gopd()
	//ww.Wait()
}
func gopd() {
	go urls()
}
func first() {
	if LLConfig.Eula == false || LLConfig.PrivacyPolicy == false {
		log.Println("Minecraft End User License Agreement: https://account.mojang.com/terms\nPrivacy Policy: https://go.microsoft.com/fwlink/?LinkId=521839\nDo you agree to the terms above? \nIf you agree, please open /plugins/BDSUpdate/BDSUpdate.json and change \"eula\" and \"privacypolicy\" to true")
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}
}
func FileChecker(filename string) (bool, string) {
	file_path := filename
	_, err := os.Stat(file_path)
	if err == nil {
		return true, "FileChecker:::Found " + file_path
	} else {
		return false, "FileChecker:::NotFound "
	}
}

func urls() {
	log.Println("BDSUpdate, By DreamGuXiang")
	log.Println("Version Release 0.0.1 (GoLang)")
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			resp, err := http.Get("https://www.minecraft.net/en-us/download/server/bedrock/")
			if err != nil {
				//log.Println("http get error.")
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println("http read error.")
				}
				src := string(body)
				docs := strings.NewReader(src)
				doc, err := html.Parse(docs)
				if err != nil {
					fmt.Fprint(os.Stderr, "findlinks: %v\n", err)
					os.Exit(1)
				}
				for _, link := range visit(nil, doc) {
					rds := bufio.NewReader(strings.NewReader(link))
					for {
						lined, err := rds.ReadString('\n')
						res1 := strings.Contains(link, "https://minecraft.azureedge.net/")
						if err == nil || io.EOF == err {
							if res1 == true {
								bdswin := strings.Contains(lined, "bin-win")
								bdslinux := strings.Contains(lined, "bin-linux")
								w.Add(1)
								if bdswin == true {
									delline := strings.TrimPrefix(lined, "https://minecraft.azureedge.net/bin-win/")
									bts := "windows-"
									build.WriteString(bts)
									build.WriteString(delline)
									s3 := build.String()
									build.Reset()
									own, _ := FileChecker(LLConfig.DownloadDestFolder + s3)
									if own == true {
										break
									} else if own == false {
										go download(lined, s3, &w)
										w.Wait()
										break
									}
								} else if bdslinux == true {
									delline := strings.TrimPrefix(lined, "https://minecraft.azureedge.net/bin-linux/")
									bts := "linux-"
									build.WriteString(bts)
									build.WriteString(delline)
									s3 := build.String()
									build.Reset()
									own, _ := FileChecker(LLConfig.DownloadDestFolder + s3)
									if own == true {
										break
									} else if own == false {
										go download(lined, s3, &w)
										w.Wait()
										break
									}
									break
								}
							}
							break
						}
					}
				}
			}
			t1.Reset(time.Second * LLConfig.Times)
		}
	}
}

func inits() {
	_ = os.MkdirAll("plugins\\BDSUpdate", 0777)
}

func main() {}
