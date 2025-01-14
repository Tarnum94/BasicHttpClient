package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Tarnum94/BasicHttpClient/gohttp"
)

func main() {

	client := gohttp.New()

	url := "http://github.com"
	var headers http.Header = http.Header{}
	headers.Add("Accept", "application/xml")
	headers.Add("Accept", "appliaction/json")

	/*:= Header{
	map[string][]string {
		"Accept":	[]string{"application/xml", "application/json"}
	}
	*/

	response, err := client.Get(url, headers)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
