package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Tarnum94/BasicHttpClient/gohttp"
)

var (
	client = getGitHubClient()
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getGitHubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Add("Accept", "application/xml")

	client.SetHeaders(commonHeaders)

	return client
}

func getURL(url string) {

	headers := http.Header{"Accept": []string{"application/json"}}
	/*headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	headers.Add("Accept", "application/xml")
	headers.Add("Accept", )*/

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

func createUser(url string, user User) {
	headers := http.Header{"Accept": []string{"application/json"}}
	/*headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	headers.Add("Accept", "application/xml")
	headers.Add("Accept", )*/

	response, err := client.Post(url, headers, user)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}

func main() {
	url := "https://www.kinopolis-web.com"
	//user := User{FirstName: "Meik", LastName: "Okon"}
	getURL(url)
	//createUser(url, user)
}
