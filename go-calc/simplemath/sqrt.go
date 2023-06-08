package simplemath

import "math"

// Sqrt 计算平方根
func Sqrt(a int) int {
	v := math.Sqrt(float64(a))
	return int(v)
}
