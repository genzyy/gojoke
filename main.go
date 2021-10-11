package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	fmt.Println("Joke in 3..2..1...")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	// fmt.Printf("API response as struct %+v\n", responseObject)
	fmt.Println(responseObject.Joke)
}
