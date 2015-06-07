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

func save_flag(img []byte, filename string) {
	path := fmt.Sprintf("%v%v", DEST_DIR, filename)
	fmt.Println(path)
	err := ioutil.WriteFile(path, img, 0666)
	check(err)
}

func get_flag(cc string) {
	url := fmt.Sprintf("%[1]v/%[2]v/%[2]v.gif", BASE_URL, strings.ToLower(cc))
	fmt.Println(cc, url)
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(body[:6]), resp.ContentLength)
	return body
}

func download_many(cc_list []string) {
	sort.Strings(cc_list)
	for _, cc := range cc_list {
		body := get_flag(cc)
		save_flag(body, fmt.Sprintf("%v.gif", cc))
	}
	return len(cc_list)
}

func main() {
	t0 := time.Now()
	pop20_cc := strings.Split(POP20_CC, " ")

	count := download_many(pop20_cc)
	elapsed := time.Since(t0)
	fmt.Println(count, "flags downloaded in", elapsed)
}
