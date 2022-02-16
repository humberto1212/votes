package main

import (
	"Desktop/Golang/Votes/datafile"
	"fmt"
	"log"
)

func main() {

	lines, err := datafile.GetStrings("./files/votes")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

}
