package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var x interface{}
	err := json.NewDecoder(os.Stdin).Decode(&x)
	if err != nil {
		log.Fatal(err)
	}
	o, err := json.MarshalIndent(x, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(o)
}
