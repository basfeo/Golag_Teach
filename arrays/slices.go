package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := [][]uint8{}
	for i := 0; i < dy; i++ {
		v := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			v[j] = uint8(float64(i*i)/(float64(j+1)) + float64(j*j)/(float64(i+1)))
		}
		result = append(result, v)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
