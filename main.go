package main

import (
	"Desktop/Golang/Votes/count"
	"Desktop/Golang/Votes/datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	var names []string
	lines, err := datafile.GetStrings("./files/votes")
	if err != nil {
		log.Fatal(err)
	}

	counts := count.CountLines(lines)

	for presidentName := range counts {
		names = append(names, presidentName)
	}

	sort.Strings(names)

	//fmt.Println(names)

	for _, name := range names {
		fmt.Printf("%s has %v votes \n", name, counts[name])
	}

	//------ slice version ----------------
	// names, counts := count.CountLines(lines)

	// for i, name := range names {
	// 	fmt.Printf("%s: %d\n", name, counts[i])
	// }

}
