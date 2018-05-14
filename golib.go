// Harry Podciborski
// Useful GoLang conversion library

package golib
import (
   "math" 
   "strconv" 
)


func ParseMath32(s string) (f float32, err error) {
	i, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return
	}
	f = math.Float32frombits(uint32(i))
	return
}

func ParseMath64(s string) (f float64, err error) {
	i, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return
	}
	f = math.Float64frombits(uint64(i))
	return
}

func strToint(s string) int { 		// convert string to integer
	var i int
	i,_ = strconv.Atoi(s)
	return i
}