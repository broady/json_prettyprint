package main

import (
	"encoding/json"
	"os"
)

func main() {
	var x interface{}
	err := json.NewDecoder(os.Stdin).Decode(&x)
	if err != nil {
		panic(err)
	}
	o, err := json.MarshalIndent(x, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(o)
}
