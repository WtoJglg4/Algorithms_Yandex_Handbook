package main

func HanoiTowers(n, fromPeg, toPeg int, s *int) {
	*s++
	if n == 1 {
		//fmt.Printf("Move disk from %v to peg %v\n", fromPeg, toPeg)
		return
	}
	unusedPeg := 6 - fromPeg - toPeg
	HanoiTowers(n-1, fromPeg, unusedPeg, s)
	//fmt.Printf("Move disk from %v to peg %v\n", fromPeg, toPeg)
	HanoiTowers(n-1, unusedPeg, toPeg, s)
}

// or 2^n - 1 => HanoiTowers - exponencial algorithm
func FastHanoiTowers(n int) int {
	if n == 1 {
		return 1
	}
	return 2*FastHanoiTowers(n-1) + 1
}
