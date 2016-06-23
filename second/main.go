package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(numbers []string) []string {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	in, _ := os.Open("B-small-practice.in")
	out, _ := os.Create("B-small-practice.out")
	reader := bufio.NewReader(in)
	var n int
	h, _ := reader.ReadString('\n')

	n, _ = strconv.Atoi(h[:len(h)-1])
	for i := 0; i < n; i++ {
		var str string
		str, _ = reader.ReadString('\n')
		p := strings.Split(str[:len(str)-1], " ")
		p = reverse(p)
		str = strings.Join(p, " ")
		fmt.Fprintln(out, "Case #"+strconv.Itoa(i+1)+":", str)
	}
}
