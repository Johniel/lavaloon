package lavaloon

type tokenType uint8

type Position struct {
	line   int
	column int
}

type Token struct {
	Type tokenType
	Val  string
	Pos  Position
}
