// Harry Podciborski
// Useful GoLang conversion library

package golib
import (
   "math" 
   "strconv" 
   "fmt"
   "bytes"
   "encoding/binary"
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
		return 0, fmt.Errorf("Error converting string to float64")	
	}
	f = math.Float64frombits(uint64(i))
	return f,nil
}

//StrToInt : Convert string to integer
func StrToInt(s string) (int,error) { 			
	i,err := strconv.Atoi(s)
	if err != nil {		
		return 0, fmt.Errorf("Error converting string to int")	
	} 
	return i,nil
}


// DecodeInt attempts to convert a byte array of High/Low Byte values into
// a 16-bit integer, and returns the result, also with an error, which will
// be non-nil if the decoding failed.
func DecodeInt(data []byte) (int16, error) {
	var i int16
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}


// GetString Byte to string from hex buffer
func GetString (in []byte) string  { 
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {			// copied 
		var hdata  = fmt.Sprintf("%02x",in[i+1])		
		hexbuffer.WriteString(hdata)	//concat
	}
	return hexbuffer.String()
}



// GetInt Byte to uint64
func GetInt (in []byte) (uint64,error)  { 		
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {		
		// take the next two bytes, if available
		// get hex data out of bytes
		var hdata  = fmt.Sprintf("%02x",in[i:i+2])
		hexbuffer.WriteString(hdata)	//concat
	}
	f, err := strconv.ParseUint(hexbuffer.String(), 16, 32)	
	if err != nil {		
		return 0, fmt.Errorf("Error converting byte to uint64")	
	} 
	return f,nil
}


//GetFloat32 some refernence docs:
//https://play.golang.org/
//http://www.perlmonks.org/?node_id=327171
func GetFloat32 (in []byte) (float32,error) { 		// 32bit encoded IEEE754 
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {		
		// take the next two bytes, if available
		// get hex data out of bytes
		var hdata  = fmt.Sprintf("%02x",in[i:i+2])		
		hexbuffer.WriteString(hdata)	//concat
	}
	f, err := ParseMath32(hexbuffer.String())
	if err != nil {		
		return 0, fmt.Errorf("Error converting byte to float32")	
	} 
	return f,nil
}

// GetFloat64 64bit encoded IEEE754 
//https://www.h-schmidt.net/FloatConverter/IEEE754.html
func GetFloat64 (in []byte) (float64,error) { 		
	var hexbuffer bytes.Buffer // concatenate x2 pairs of hex vals	
	for i := 0; i < len(in); i+=2 {		
		// take the next two bytes, if available
		// get hex data out of bytes
		var hdata  = fmt.Sprintf("%02x",in[i:i+2])
		hexbuffer.WriteString(hdata)	//concat
	}
	f, err := ParseMath64(hexbuffer.String())
	if err != nil {		
		return 0, fmt.Errorf("Error converting byte to float64")	
	} 
	return f,nil
}


// Float64ToString convert float64 to string
func Float64ToString(inputNum float64) string {
	// to convert a float number to a string
	// "var float float32" up here somewhere
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}


// StringToFloat - convert string to float 32
func StringToFloat (input string) (float32,error) { 	
	value, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, fmt.Errorf("Error converting string to float32")	
	}
	return float32(value),nil
}


// StringToFloat64 convert String to Float64
func StringToFloat64(input string)  (float64, error) {     
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Error converting string to float64")		
	}
	return float64(value),nil
}





