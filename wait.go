package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Prizes struct {
	Prizes []years `json:"prizes"`
}

type years struct {
	Year      string `json:"year"`
	Category  string `json:"category"`
	Laureates string `json:"laureates"`
}
type ids struct {
	ID         string `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Motivation string `json:"motivation"`
	Share      string `json:"share"`
}

func main() {
	jsonFile, err := os.Open("prize.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Json file")
	defer jsonFile.Close()

	file, _ := ioutil.ReadAll(jsonFile)

	var prizess Prizes

	json.Unmarshal(file, &prizess)

	// template, _ := template.ParseFiles("hello.html")
	// template.Execute(w, file)

	for i := 0; i < len(prizess.Prizes); i++ {
		fmt.Println("Year: ", prizess.Prizes[i].Year)
		fmt.Println("Category: ", prizess.Prizes[i].Category)
		fmt.Println("Lauretes:", prizess.Prizes[i].Laureates)
	}
}
