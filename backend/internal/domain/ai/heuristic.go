package ai

func scoreLine(line []string, me, opp string) int {
	mine, theirs := 0, 0
	for _, c := range line {
		if c == me {
			mine++
		}
		if c == opp {
			theirs++
		}
	}
	if mine > 0 && theirs > 0 {
		return 0
	}
	if mine > 0 {
		return mine * mine
	}
	if theirs > 0 {
		return -(theirs * theirs)
	}
	return 0
}
