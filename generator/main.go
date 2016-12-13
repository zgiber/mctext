package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	mct "github.com/zgiber/mctext"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("usage:\ngenerate inputfile length")
		return
	}

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mc := mct.New()
	mc.Parse(f)

	length, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mc.Generate("", int(length)))
	return
}
