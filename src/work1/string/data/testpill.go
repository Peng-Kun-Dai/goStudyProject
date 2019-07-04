package main

//go:generate stringer -type=testPill
type testPill int

const (
	A testPill = iota
	B
	C
	D
	E = D
)
