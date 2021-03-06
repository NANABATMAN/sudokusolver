package main

func main() {
	playField := PlayField{}
	playField.SetPlayField(&[SIZE][SIZE]int{
		{2, 0, 0, 0, 0, 5, 8, 9, 0},
		{0, 6, 0, 0, 9, 0, 7, 1, 0},
		{9, 0, 0, 7, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 6, 0, 0, 5, 8},
		{0, 0, 6, 0, 0, 0, 4, 0, 0},
		{5, 9, 0, 0, 3, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 8, 0, 0, 7},
		{0, 4, 9, 0, 2, 0, 0, 8, 0},
		{0, 8, 5, 1, 0, 0, 0, 0, 9},
	})
	playField.Solve()
	playField.PrintPlayField()
}
