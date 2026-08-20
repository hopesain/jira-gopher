package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hash Maps")

	m := make(map[string]int)
	m["hope"] = 24
	m["trevor"] = 25
	m["zione"] = 31
	m["chikondi"] = 25
	m["thoko"] = 29
	m["blessings"] = 27
	m["patrick"] = 33

	age, exists := m["trevor"]
	fmt.Println(age, exists) 

	age, exists = m["notname"]
	fmt.Println(age, exists)

	delete(m, "thoko")

	for name, age := range m {
		fmt.Println(name, age)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Jira")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}
