// Download flags of top 20 countries by population
//
// Sequential version with no error checking.
//
// Sample run:
//
//	$ go run flags_seq.go
//	BD BR CD CN DE EG ET FR ID IN IR JP MX NG PH PK RU TR US VN
//	20 flags downloaded in 3.86s

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

func saveFlag(img []byte, filename string) {
	path := fmt.Sprintf("%v%v", DEST_DIR, filename)
	ioutil.WriteFile(path, img, 0660)
}

func getFlag(cc string) []byte {
	url := fmt.Sprintf("%[1]v/%[2]v/%[2]v.gif", BASE_URL, strings.ToLower(cc))
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
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
