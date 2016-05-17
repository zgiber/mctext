package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	mc "github.com/zgiber/mchaintext"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage:\ngenerate inputfile length")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	mc := &mc.{
		map[string]*mc.Node{},
	}
	mc.Parse(f)

	length, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mc.Generate("", int(length)))
	return
}
