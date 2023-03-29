package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	lg "github.com/bahodurnazarov/Dogs_API/utils"
)

type scheme struct {
	Message string `json:"message"`
}

type DogLink struct {
	Image string
}

func GetDogLink() string {

	url := "https://dog.ceo/api/breeds/image/random"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		lg.Errl.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		lg.Errl.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		lg.Errl.Println(err)

	}
	scheme1 := scheme{}
	jsonErr := json.Unmarshal(body, &scheme1)
	if jsonErr != nil {
		lg.Errl.Fatal(jsonErr)
	}

	fmt.Println(scheme1.Message)

	return scheme1.Message
}
