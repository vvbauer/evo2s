package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	getTerminalIDs()
	getReciepts()
}

func getTerminalIDs() {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://api.evotor.ru/api/v1/inventories/devices/search", nil)

	request.Header.Add("X-Authorization", "token")
	request.Header.Add("User-Agent", "PostmanRuntime/7.20.1")
	request.Header.Add("Accept", "*/*")

	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func getReciepts() {

	client := http.Client{}

	request, err := http.NewRequest("GET", "https://api.evotor.ru/api/v1/inventories/stores/search", nil)

	request.Header.Add("X-Authorization", "token")
	request.Header.Add("User-Agent", "PostmanRuntime/7.20.1")
	request.Header.Add("Accept", "*/*")

	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", body)
}
