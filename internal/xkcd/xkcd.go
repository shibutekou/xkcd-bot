package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BaseURL string = "https://xkcd.com"
const EndURL string = "/info.0.json"

func GetComics(number string) Comics {
	response, err := http.Get(fmt.Sprintf("%s%s%s", BaseURL, number, EndURL))
	if err != nil {
		log.Fatal(err.Error())
	}

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var comics Comics
	err = json.Unmarshal(resBody, &comics)

	return comics
}

type Comics struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
