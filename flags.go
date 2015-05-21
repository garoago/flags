package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
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
}

func get_flag(cc string) {
	url := fmt.Sprintf("%[1]v/%[2]v/%[2]v.gif", BASE_URL, strings.ToLower(cc))
	fmt.Println(cc, url)
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body[:6]), resp.ContentLength)
	save_flag(body, fmt.Sprintf("%v.gif", cc))
}

func main() {
	pop20_cc := strings.Split(POP20_CC, " ")
	sort.Strings(pop20_cc)
	for _, cc := range pop20_cc {
		get_flag(cc)
	}
}
