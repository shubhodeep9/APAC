package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fd, _ := os.Open("A-large-practice.in")
	out, _ := os.Create("A-large-practice.out")
	var n, c, items int
	fmt.Fscanf(fd, "%d", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(fd, "%d", &c)
		fmt.Fscanf(fd, "%d", &items)
		p := make([]int, items)
		for j := 0; j < items; j++ {
			fmt.Fscanf(fd, "%d", &p[j])
		}
		for j := 0; j < items-1; j++ {
			for k := j + 1; k < items; k++ {
				if p[j]+p[k] == c {
					fmt.Fprintln(out, "Case #"+strconv.Itoa(i+1)+":", j+1, k+1)
				}
			}
		}
	}
}
