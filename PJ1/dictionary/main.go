package main

import (
	"fmt"

	"github.com/minju-kim98/learngo/PJ1/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	def := "Greeting"
	err := dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}

	//Test Search
	definition, err2 := dictionary.Search(word)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}

	err3 := dictionary.Add(word, def)
	if err3 != nil {
		fmt.Println(err3)
	}

	// Test Update
	err4 := dictionary.Update("Hello", "Adele's song")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(word, "'s definition changed.")
	}

	// Test Delete
	_, err5 := dictionary.Search(word)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("Search Compelete!")
	}

	err6 := dictionary.Delete(word)
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Println("Delete Compelete!")
	}

	_, err5 = dictionary.Search(word)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("Search Compelete!")
	}

	err6 = dictionary.Delete(word)
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Println("Delete Compelete!")
	}

}
