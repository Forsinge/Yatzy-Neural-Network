package input

import (
	"fmt"
	"io"
)

func ScanInt() int {
	i := 0
	fmt.Scanf("%d\r", &i)
	return i
}

func ScanInts() []int {
	ints := []int{}
	i := 0
	for {
		n, err := fmt.Scanf("%d", &i)
		if err == io.EOF || n == 0 || err != nil {
			break
		}
		ints = append(ints, i)
	}
	return ints
}
