package main

import (
	"Desktop/Golang/Votes/count"
	"Desktop/Golang/Votes/datafile"
	"fmt"
	"log"
)

func main() {

	lines, err := datafile.GetStrings("./files/votes")
	if err != nil {
		log.Fatal(err)
	}

	counts := count.CountLines(lines)

	for key, vote := range counts {
		fmt.Println(key, vote)
	}

	//------ slice version ----------------
	// names, counts := count.CountLines(lines)

	// for i, name := range names {
	// 	fmt.Printf("%s: %d\n", name, counts[i])
	// }

}
