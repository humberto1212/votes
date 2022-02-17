package main

import (
	"Desktop/Golang/Votes/count"
	"Desktop/Golang/Votes/datafile"
	"log"
)

func main() {

	lines, err := datafile.GetStrings("./files/votes")
	if err != nil {
		log.Fatal(err)
	}

	count.CountLines(lines)

}
