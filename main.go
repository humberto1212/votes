package main

import (
	"Desktop/Golang/Votes/datafile"
	"fmt"
)

func main() {

	fmt.Println(datafile.GetStrings("./files/votes"))

}
