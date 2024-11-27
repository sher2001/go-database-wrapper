package main

import (
	"fmt"
	"log"

	mybase "github.com/sher2001/myBase/myBase"
)

func main() {

	db, err := mybase.New()
	if err != nil {
		log.Fatal(err)
	}

	// temporary support for string, future we'll support for other data types
	user := map[string]string{
		"name": "Vishnu",
		"age":  "24",
	}

	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}

	// collection, err := db.CreateCollection("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(id.String())
}
