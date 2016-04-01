package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func outputResponse(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func main() {
	ret, err := outputResponse("http://uuid.jp/")
	if err != nil {
		return
	}
	fmt.Println(string(ret))
}
