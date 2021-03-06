package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/encoding"
)

func main() {

	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error", resp.StatusCode)
		return 
	}

	e := determineEncoding(resp.Body)
	utf8Reader:= transform.NewReader(resp.Body,
		e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ :=charset.DetermineEncoding(bytes, "")
	return e
}
