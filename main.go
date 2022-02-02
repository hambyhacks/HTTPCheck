package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/hambyhacks/HTTPCheck/utils"
)

func main() {
	filename := flag.String("f", "", ".txt file which contains URLs")
	flag.Parse()

	regex_txt, _ := regexp.MatchString(".txt", *filename)

	if os.Args[1] == "" || os.Args[1] == "-h" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(os.Args) < 1 || !regex_txt {
		flag.PrintDefaults()
		log.Fatal("Error!")
		os.Exit(1)
	}
	utils.CheckHTTPResponse(*filename)
}
