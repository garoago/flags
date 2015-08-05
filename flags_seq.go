package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

// country codes of the 20 most populous countries in 2014
const POP20_CC = "CN IN US ID BR PK NG BD RU JP MX PH VN ET EG DE IR TR CD FR"

const BASE_URL = "http://flupy.org/data/flags"

const DEST_DIR = "downloads/"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveFlag(img []byte, filename string) {
	path := fmt.Sprintf("%v%v", DEST_DIR, filename)
	err := ioutil.WriteFile(path, img, 0660)
	check(err)
}

func getFlag(cc string) []byte {
	url := fmt.Sprintf("%[1]v/%[2]v/%[2]v.gif", BASE_URL, strings.ToLower(cc))
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body
}

func downloadOne(cc string) {
	image := getFlag(cc)
	saveFlag(image, fmt.Sprintf("%v.gif", strings.ToLower(cc)))
}

func DownloadMany(cc_list []string) int {
	sort.Strings(cc_list)
	for _, cc := range cc_list {
		downloadOne(cc)
		fmt.Print(cc, " ")
	}
	return len(cc_list)
}

func main() {
	t0 := time.Now()
	count := DownloadMany(strings.Split(POP20_CC, " "))
	elapsed := time.Since(t0)
	msg := "\n%d flags downloaded in %.2fs\n"
	fmt.Printf(msg, count, elapsed.Seconds())
}
