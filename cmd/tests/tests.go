package main

import (
	"net/http"
	"net/url"
)

const domain = "http://localhost/"

func main() {
	client := http.Client{}

}

func POST(client http.Client, path string, request []byte) (int, []byte) {
	req := http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Path: domain + path,
		},
	}
}

// TODO: test with bad data
// TODO: test nonexsting endpoints
// TODO: test with wrong methods
// TODO: bad JSON
