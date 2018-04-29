package lavaloon

type Position struct {
	line   int
	column int
}

type Token struct {
	Str string
	Pos Position
}
