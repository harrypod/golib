// Harry Podciborski
// Useful GoLang conversion library

package golib
import (
   "math" 
   "strconv" 
   "fmt"
   "bytes"
)

//ParseMath32 : Convert string to 32bit Float
func ParseMath32(s string) (f float32, err error) {
	i, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return
	}
	f = math.Float32frombits(uint32(i))
	return
}


//ParseMath64 : Convert string to 64bit Float
func ParseMath64(s string) (f float64, err error) {
	i, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return
	}
	f = math.Float64frombits(uint64(i))
	return
}

//StrToInt : Convert string to integer
func StrToInt(s string) int { 		
	var i int
	i,_ = strconv.Atoi(s)
	return i
}



//GetFloat32 some refernence docs:
//https://play.golang.org/
//http://www.perlmonks.org/?node_id=327171
func GetFloat32 (in []byte) float32 { 		// 32bit encoded IEEE754 
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {		
		// take the next two bytes, if available
		// get hex data out of bytes
		var hdata  = fmt.Sprintf("%02x",in[i:i+2])		
		hexbuffer.WriteString(hdata)	//concat
	}
	f, err := ParseMath32(hexbuffer.String())
	if err != nil {
		return (0)
	} 
	return f
}

// GetFloat64 64bit encoded IEEE754 
//https://www.h-schmidt.net/FloatConverter/IEEE754.html
func GetFloat64 (in []byte) float64 { 		
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {		
		// take the next two bytes, if available
		// get hex data out of bytes
		var hdata  = fmt.Sprintf("%02x",in[i:i+2])
		hexbuffer.WriteString(hdata)	//concat
	}
	f, err := ParseMath64(hexbuffer.String())
	if err != nil {
		return (0)
	} 
	return f
}