package count

func CountLines(list []string) ([]string, []int) {
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

	return names, counts

}
