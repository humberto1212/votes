package count

func CountLines(list []string) map[string]int {
	//------------ Map version --------------
	counts := make(map[string]int)

	for _, line := range list {
		counts[line]++
	}

	return counts

	//------------ slice version --------------

	// var names []string
	// var counts []int

	// for _, line := range list {
	// 	matched := false
	// 	for i, name := range names {
	// 		if name == line {
	// 			counts[i]++
	// 			matched = true
	// 		}
	// 	}
	// 	if !matched {
	// 		names = append(names, line)
	// 		counts = append(counts, 1)

	// 	}
	// }

	// return names, counts

}
