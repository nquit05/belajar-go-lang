package main

import "fmt"

func main() {

	person := map[string]string{
		"name":   "Michael",
		"alamat": "OKe",
	}

	person["title"] = "Ganteng"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	// var book map[string]string = make(map[string]string)
	book := make(map[string]string)

	book["title"] = "Senna Cantekkk"
	book["author"] = "Michael Gantengg"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "author")

	fmt.Println(book)

}
