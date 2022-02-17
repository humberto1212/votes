package count

import "fmt"

func CountLines(list []string) {
	var names []string
	var counts []int

	for _, line := range list {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if !matched {
			names = append(names, line)
			counts = append(counts, 1)

		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

}
