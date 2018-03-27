package _06NumberofLinesToWriteString


func numberOfLines(widths []int, S string) []int {
	line, nextLine := 1, 0

	for _, c := range S {
		if nextLine + widths[int(c-97)] > 100 {
			nextLine = widths[int(c-97)]
			line++
		} else {
			nextLine += widths[int(c-97)]
		}
	}
	return []int{line, nextLine}
}
