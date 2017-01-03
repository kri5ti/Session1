package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s $URL", os.Args[0])
	}
	result := make(map[string]string)
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := result["origin"]; ok {
		fmt.Println(result["origin"])
	} else {
		fmt.Println("Result not standart")
	}
}
