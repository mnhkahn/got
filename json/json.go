package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, data, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	out.WriteTo(os.Stdout)
}
