package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	T9 := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
		0: {" "},
	}
	in, _ := os.Open("C-large-practice.in")
	out, _ := os.Create("C-large-practice.out")
	reader := bufio.NewReader(in)
	var n int
	h, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(h[:len(h)-1])
	for m := 0; m < n; m++ {
		var str string
		str, _ = reader.ReadString('\n')
		str = str[:len(str)-1]
		fmt.Fprint(out, "Case #"+strconv.Itoa(m+1)+": ")
		for i := range str {

			for j := range T9 {
				for k := range T9[j] {
					if T9[j][k] == string(str[i]) {
						for n := range T9[j] {
							if i > 0 && T9[j][n] == string(str[i-1]) {
								fmt.Fprint(out, " ")
								break
							}
						}
						for l := 0; l <= k; l++ {
							fmt.Fprint(out, j)
						}
					}
				}
			}
		}
		fmt.Fprintln(out)
	}
}
