package apiclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL   string
	UserAgent string
}

func (c *Client) get(resp, error) {
	resp, err := http.Get("https://httpbin.org/anything/test")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
	return
}
