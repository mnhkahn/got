package main

// tr -d '\040\011\012\015'
import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err)
		return
	}

	out := strings.Replace(string(data), " ", "", -1)
	out = strings.Replace(out, "\n", "", -1)
	log.Println(out)
}
