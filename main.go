package main

import "fmt"

func Divisioner(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Zero division error")
	}
	return float64(a) / float64(b), nil
}

func main() {
	res, _ := Divisioner(12, 5)
	fmt.Println("result:", res)
}
