// Harry Podciborski
// Useful GoLang conversion library

package podgolib
import (
   "math" 
   "strconv" 
)



func parseMath32(s string) (f float32, err error) {
	i, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return
	}
	f = math.Float32frombits(uint32(i))
	return
}

func parseMath64(s string) (f float64, err error) {
	i, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return
	}
	f = math.Float64frombits(uint64(i))
	return
}
