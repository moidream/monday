package template

import "fmt"

func ioTem() {
	var m, n int
	fmt.Scan(&m, &n)
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := range a {
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
}
