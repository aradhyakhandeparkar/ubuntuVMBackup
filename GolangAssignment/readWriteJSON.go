package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// type Book struct {
// 	Id         int
// 	BookName   string
// 	AuthorName string
// 	LoggedAt   time.Time
// }

func main() {
	//using structure
	// book := Book{}
	// book.Id = 1
	// book.BookName = "Title"
	// book.AuthorName = "FirstNameLastName"
	// book.LoggedAt = time.Now()

	//using map
	// book := map[string]string{
	// 	"id":     "1",
	// 	"name":   "Title",
	// 	"author": "FirstNameLastName",
	// }

	//using nested map
	var book = map[string]map[string]string{}
	book["bk1"] = map[string]string{}
	book["bk2"] = map[string]string{}

	book["bk1"]["id"] = "1"
	book["bk1"]["name"] = "Title1"
	book["bk1"]["author"] = "FNLN"

	book["bk2"]["id"] = "2"
	book["bk2"]["name"] = "Title2"
	book["bk2"]["author"] = "fNlN"
	//...................................
	//Writing struct type to a JSON file
	//...................................
	content, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("book.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	content, err = ioutil.ReadFile("book.json")
	if err != nil {
		log.Fatal(err)
	}
	//book2 := Book{}
	var book2 map[string]map[string]interface{}
	err = json.Unmarshal(content, &book2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Id:%v, BookName:%v, AuthorName:%v \n", book2["bk1"]["id"], book2["bk1"]["name"], book2["bk1"]["author"])
	log.Printf("Id:%v, BookName:%v, AuthorName:%v \n", book2["bk2"]["id"], book2["bk2"]["name"], book2["bk2"]["author"])

}
