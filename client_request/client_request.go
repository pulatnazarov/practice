package client_request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AgifyResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func ClientRequest() {
	var name string

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, "https://api.agify.io", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var response AgifyResponse

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}

	fmt.Printf("%s's age is %d\n", response.Name, response.Age)
}
