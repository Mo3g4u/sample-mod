package data

var ModCount uint

func CountUp() uint {
	ModCount++
	return ModCount
}
