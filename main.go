package main

import (
	"encoding/json"
	"log"
)

type Json struct {
	Title  string   `json:"title"`
	Name   string   `json:name`
	Year   int      `json:"released"`
	Color  bool     `json:"color"`
	Actors []string `json:"actors"`
}

func main() {
	j := Json{
		Title: "Unfortunate Events",
		Name:  "movie name",
		Year:  2018,
		Color: true,
		Actors: []string{
			"Violet",
			"Klause",
			"Sunny",
		},
	}

	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Println(string(b))

	var s struct {
		Title string `json:"title"`
	}
	err = json.Unmarshal(b, &s)

	if err != nil {
		panic(err)
	}
	log.Println(s)
}
